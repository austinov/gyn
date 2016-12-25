package route

import (
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	initMiddleware(e)
	initStatic(e)
	initAPI(e)
}
