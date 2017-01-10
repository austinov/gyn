package route

import "github.com/labstack/echo"

func initStatic(e *echo.Echo) {
	e.File("/", "../ui-web/html/index.html")
	e.File("/favicon.ico", "../ui-web/html/favicon.ico")
	e.Static("/dist", "../ui-web/dist")
}
