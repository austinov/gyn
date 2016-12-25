package route

import (
	"net/http"

	"github.com/labstack/echo"
)

func initAPI(e *echo.Echo) {
	e.GET("/api/login", login)
}

func login(c echo.Context) error {
	// TODO
	return c.String(http.StatusOK, "login...")
}
