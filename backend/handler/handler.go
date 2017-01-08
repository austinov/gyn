package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/store"
	"github.com/austinov/gyn/backend/util"
	"github.com/nguyenthenguyen/docx"
	"github.com/pkg/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Handler interface {
	Login(c echo.Context) error
	GetProfile(c echo.Context) error
	GetDictionaries(c echo.Context) error
	SearchAppointments(c echo.Context) error
	GetAppointment(c echo.Context) error
	GetAppointmentDocx(c echo.Context) error
	SaveAppointment(c echo.Context) error
}

func New(cfg config.Config, dao store.Dao, ec ErrorCustomizer) Handler {
	return handler{
		cfg: cfg,
		dao: dao,
		ec:  ec,
	}
}

type handler struct {
	cfg config.Config
	dao store.Dao
	ec  ErrorCustomizer
}

type (
	loginForm struct {
		Login    string `form:"login" valid:"required~login-required"`
		Password string `form:"password" valid:"required~password-required,password~password-format"`
	}
	loginReply struct {
		RedirectURL string `json:"redirUrl"`
	}
	searchRequest struct {
		PatientName string `json:"patientName"`
	}
	saveReply struct {
		Id int64 `json:"id"`
	}
)

func (h handler) Login(c echo.Context) error {
	var lf loginForm
	if err := c.Bind(&lf); err != nil {
		return errors.WithStack(err)
	}

	if _, err := govalidator.ValidateStruct(lf); err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusBadRequest, h.ec.InvalidRequestParameterError(err))
	}

	if err := h.dao.Authenticate(lf.Login, lf.Password); err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusUnauthorized, h.ec.UserAuthenticationError(err))
	}

	token, err := util.NewToken(h.cfg.JWT.Issuer, lf.Login, "", h.cfg.JWT.Expiration, h.cfg.JWT.SignKey)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusBadRequest, h.ec.ServerError(err))
	}

	// just return access token to client
	cookie := createCookie(h.cfg.AuthCookieName, token)
	c.SetCookie(cookie)

	reply := loginReply{
		RedirectURL: "",
	}
	return c.JSON(http.StatusOK, reply)
}

func (h handler) GetProfile(c echo.Context) error {
	claims := c.Get(h.cfg.Ctx.Key).(*jwt.StandardClaims)

	profile, err := h.dao.GetProfile(claims.Audience)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusUnauthorized, h.ec.ServerError(err))
	}
	return c.JSON(http.StatusOK, profile)
}

func (h handler) GetDictionaries(c echo.Context) error {
	dicts, err := h.dao.GetDictionaries()
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	return c.JSON(http.StatusOK, dicts)
}

func (h handler) SearchAppointments(c echo.Context) error {
	var r searchRequest
	err := c.Bind(&r)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusBadRequest, h.ec.InvalidRequestParameterError(err))
	}
	aps, err := h.dao.SearchAppointments(r.PatientName)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	return c.JSON(http.StatusOK, aps)
}

func (h handler) GetAppointment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusBadRequest, h.ec.InvalidRequestParameterError(err))
	}
	ap, err := h.dao.GetAppointment(id)
	if err != nil {
		if err == store.ErrDataNotFound {
			return c.JSON(http.StatusBadRequest, h.ec.NoDataError(err))
		}
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	return c.JSON(http.StatusOK, ap)
}

func (h handler) GetAppointmentDocx(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusBadRequest, h.ec.InvalidRequestParameterError(err))
	}
	ap, err := h.dao.GetAppointment(id)
	if err != nil {
		if err == store.ErrDataNotFound {
			return c.JSON(http.StatusBadRequest, h.ec.NoDataError(err))
		}
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	// TODO
	file, err := util.FillDocx(ap, h.cfg.DocxDir+"template.docx", func(doc *docx.Docx) error {
		doc.Replace("[expByMenstruation]", ap.ExpByMenstruation, -1)
		return nil
	})
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	defer os.Remove(file.Name())

	c.Response().Header().Set(echo.HeaderContentType, "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	return c.Attachment(file, fmt.Sprintf("ap_%d_%d.docx", ap.Id, ap.PatientId))
}

func (h handler) SaveAppointment(c echo.Context) error {
	var ap store.Appointment
	err := c.Bind(&ap)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusBadRequest, h.ec.InvalidRequestParameterError(err))
	}
	claims := c.Get(h.cfg.Ctx.Key).(*jwt.StandardClaims)
	profile, err := h.dao.GetProfile(claims.Audience)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusUnauthorized, h.ec.ServerError(err))
	}
	ap.DoctorId = profile.Id
	c.Logger().Debugf("%#v", ap)

	if err := h.dao.SaveAppointment(&ap); err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	reply := saveReply{
		Id: ap.Id,
	}
	return c.JSON(http.StatusOK, reply)
}

func createCookie(authCookieName, accessToken string) *echo.Cookie {
	cookie := new(echo.Cookie)
	cookie.SetName(authCookieName)
	cookie.SetValue(accessToken)
	cookie.SetPath("/")
	//cookie.SetSecure(true)
	return cookie
}
