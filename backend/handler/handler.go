package handler

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/store"
	"github.com/austinov/gyn/backend/util"
	"github.com/pkg/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Handler interface {
	Login(c echo.Context) error
	Profile(c echo.Context) error
	Dictionaries(c echo.Context) error
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

func (h handler) Profile(c echo.Context) error {
	claims := c.Get(h.cfg.Ctx.Key).(*jwt.StandardClaims)

	profile, err := h.dao.GetProfile(claims.Audience)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusUnauthorized, h.ec.ServerError(err))
	}
	return c.JSON(http.StatusOK, profile)
}

func (h handler) Dictionaries(c echo.Context) error {
	dicts, err := h.dao.GetDictionaries()
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusUnauthorized, h.ec.ServerError(err))
	}
	return c.JSON(http.StatusOK, dicts)
}

func createCookie(authCookieName, accessToken string) *echo.Cookie {
	cookie := new(echo.Cookie)
	cookie.SetName(authCookieName)
	cookie.SetValue(accessToken)
	cookie.SetPath("/")
	//cookie.SetSecure(true)
	return cookie
}
