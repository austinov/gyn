package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func initMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}
