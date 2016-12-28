package route

import (
	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/handler"
	"github.com/labstack/echo"
)

func initAPI(e *echo.Echo, h handler.Handler) {
	e.POST("/api/login", h.Login)
	e.GET("/api/profile", h.Profile, tokenMiddleware(config.Get()))
}
