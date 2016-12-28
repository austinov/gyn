package route

import (
	"github.com/austinov/gyn/backend/handler"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo, h handler.Handler) {
	initMiddleware(e)
	initStatic(e)
	initAPI(e, h)
}
