package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/store"
	"github.com/austinov/gyn/backend/util"
	"github.com/nguyenthenguyen/docx"
	"github.com/pkg/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const (
	docxHeaderContentType = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
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
	c.Logger().Debugf("GetAppointment: %#v\n", ap)
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
	c.Logger().Debugf("GetAppointmentDocx: %#v\n", ap)
	file, err := util.FillDocx(ap, h.cfg.DocxDir+"template.docx", fillDocxCallback)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	defer os.Remove(file.Name())

	c.Response().Header().Set(echo.HeaderContentType, docxHeaderContentType)
	dateReceipt := time.Unix(ap.DateReceipt, 0).Format("02-01-2006_15:04")
	return c.Attachment(file, fmt.Sprintf("%s_%s.docx", util.Translit(ap.PatientName), dateReceipt))
}

func (h handler) SaveAppointment(c echo.Context) error {
	var ap store.AppointmentView
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
	c.Logger().Debugf("SaveAppointment: %#v\n", ap)

	id, err := h.dao.SaveAppointment(ap)
	if err != nil {
		c.Logger().Debugf("%+v", errors.WithStack(err))
		return c.JSON(http.StatusInternalServerError, h.ec.ServerError(err))
	}
	reply := saveReply{
		Id: id,
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

func fillDocxCallback(appointment interface{}, doc *docx.Docx) error {
	ap := appointment.(store.AppointmentView)
	doc.Replace("dateReceipt", time.Unix(ap.DateReceipt, 0).Format("02-01-2006 15:04"), -1)

	receiptDiagnosis := addTextWithComma(ap.ReceiptDiagnosis != "", ap.ReceiptKindName, "с диагнозом "+ap.ReceiptDiagnosis, ", ")
	doc.Replace("receiptDiagnosis", receiptDiagnosis, -1)

	paritet := addTextWithComma(ap.ParitetB != "", "", "Б - "+ap.ParitetB, ", ")
	paritet = addTextWithComma(ap.ParitetP != "", paritet, "Р - "+ap.ParitetP, ", ")
	paritet = addTextWithComma(ap.ParitetA != "", paritet, "А - "+ap.ParitetA, ", ")
	paritet = addTextWithComma(ap.ParitetSV != "", paritet, "С/в - "+ap.ParitetSV, ", ")
	paritet = addTextWithComma(ap.ParitetNB != "", paritet, "Нераз-бер. - "+ap.ParitetNB, ", ")
	paritet = addTextWithComma(ap.ParitetEB != "", paritet, "Экт-бер. - "+ap.ParitetEB, ", ")
	paritet = addTextWithComma(ap.Paritet != "", paritet, ap.Paritet, ", ")
	doc.Replace("paritet", paritet, -1)

	pregnancy := "на инфекционные маркеры " + ap.InfectionMarkersStateName
	pregnancy = addTextWithComma(ap.InfectionMarkersDesc != "", pregnancy, ap.InfectionMarkersDesc, " ")
	pregnancy = addTextWithComma(true, pregnancy, "на наследственную тромбофлебию "+ap.TromboflebiaStateName, ", ")
	pregnancy = addTextWithComma(ap.TromboflebiaDesc != "", pregnancy, ap.TromboflebiaDesc, " ")
	doc.Replace("pregnancy", pregnancy, -1)

	oprv := addTextWithComma(ap.Oprv != "", "", "ОПВ "+ap.Oprv+", "+ap.OprvStateName, "")
	oprv = addTextWithComma(ap.UltraInReception, oprv, "УЗИ в приемном отделении", ", ")
	oprv = addTextWithComma(ap.DopplerInReception, oprv, "Допплерометрия в приемном отделении", ", ")
	if oprv != "" {
		oprv += "\n"
	}
	doc.Replace("oprv", oprv, -1)

	expByUltra := addTextWithComma(ap.ExpByUltraFirst != "", "", ap.ExpByUltraFirst, "\n")
	expByUltra = addTextWithComma(ap.ExpByUltraSecond != "", expByUltra, ap.ExpByUltraSecond, "\n")
	expByUltra = addTextWithComma(ap.ExpByUltraThird != "", expByUltra, ap.ExpByUltraThird, "\n")
	doc.Replace("expByUltra", expByUltra, -1)

	tongue := addTextWithComma(ap.TongueClean, "", "чистый", ", ")
	tongue = addTextWithComma(ap.TongueWet, tongue, "влажный", ", ")
	tongue = addTextWithComma(ap.TongueDry, tongue, "сухой", ", ")
	tongue = addTextWithComma(ap.TongueCoated, tongue, "обложен", ", ")
	tongue = addTextWithComma(ap.TongueUncoated, tongue, "не обложен", ", ")
	doc.Replace("tongue", tongue, -1)

	bellyState := addTextWithComma(ap.EpigastriumStateUse, ap.BellyStateName, "Область эпигастрия "+ap.EpigastriumStateName, "\n")
	bellyState = addTextWithComma(ap.ScarStateUse, bellyState, "Область послеоперационных рубцов "+ap.ScarStateName, "\n")
	doc.Replace("bellyState", bellyState, -1)

	dysuric := "нет"
	if ap.Dysuric {
		dysuric = "есть"
	}
	doc.Replace("dysuric", dysuric, -1)

	bowel := "не регулярный"
	if ap.Bowel {
		bowel = "регулярный"
	}
	doc.Replace("bowel", bowel, -1)

	bishop := addTextWithComma(ap.CervixBack, "", "отклонена кзади (0 баллов)", "")
	bishop = addTextWithComma(ap.CervixFront, bishop, "кпереди (1 балл)", ", ")
	bishop = addTextWithComma(ap.CervixCenter, bishop, "центрирована (2 балла)", ", ")
	bishop = addTextWithComma(ap.CervixTight, bishop, "плотная (0 баллов)", ", ")
	bishop = addTextWithComma(ap.CervixMiddleSoft, bishop, "умеренно размягчена (1 балл)", ", ")
	bishop = addTextWithComma(ap.CervixSoft, bishop, "мягкая (2 балла)", ", ")
	bishop = addTextWithComma(ap.CervixLength != "", bishop, "длиной "+ap.CervixLength+" см (>2 см 0 баллов, 1-2 см 1 балл, < 1 см 2 балла)", ", ")
	bishop = addTextWithComma(ap.UseExternalThroat, bishop, "Наружный зев "+ap.ExternalThroatStateName, "\n")
	bishop = addTextWithComma(!ap.UseExternalThroat, bishop, "Цервикальный канал проходим для "+ap.CervixChannel+" (1п.п. до внутреннего зева 1 балл, 1п.п. свободно и > 2 балла)", "\n")
	bishop = addTextWithComma(!ap.UseExternalThroat, bishop, "Плодный пузырь "+ap.FetalBladderStateName, "\n")
	bishop = "Шейка матки (оценка по Бишопу " + ap.Bishop + " сумма баллов):" + bishop

	doc.Replace("bishop", bishop, -1)

	birthPlanName, birthPlanValue := "", ""
	if ap.BirthPlanUse {
		birthPlanName = "План родов:"
		birthPlanValue = ap.BirthPlan
	}
	doc.Replace("birthPlanN", birthPlanName, -1)
	doc.Replace("birthPlanV", birthPlanValue, -1)
	return nil
}

func addTextWithComma(cond bool, target, text, sep string) string {
	if cond {
		if target != "" {
			target += sep + text
		} else {
			target = text
		}
	}
	return target
}
