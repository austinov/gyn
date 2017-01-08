package route

import (
	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/handler"
	"github.com/labstack/echo"
)

func initAPI(e *echo.Echo, h handler.Handler) {
	e.POST("/api/login", h.Login)
	e.GET("/api/profile", h.GetProfile, tokenMiddleware(config.Get()))
	e.GET("/api/dictionaries", h.GetDictionaries, tokenMiddleware(config.Get()))
	e.POST("/api/appointments", h.SearchAppointments, tokenMiddleware(config.Get()))
	e.PUT("/api/appointment", h.SaveAppointment, tokenMiddleware(config.Get()))
	e.GET("/api/appointment/:id", h.GetAppointment, tokenMiddleware(config.Get()))
	e.GET("/api/appointment/:id/docx", h.GetAppointmentDocx, tokenMiddleware(config.Get()))
}
