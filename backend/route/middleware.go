package route

import (
	"net/http"
	"strings"

	"github.com/austinov/gyn/backend/config"
	"github.com/austinov/gyn/backend/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	InvalidAuthHeaderError = echo.NewHTTPError(http.StatusForbidden, "invalid header format")
)

func initMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
}

func tokenMiddleware(cfg config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get access token from header.
			auth := c.Request().Header().Get("Authorization")
			split := strings.SplitN(auth, " ", 2)
			if len(split) != 2 || !strings.EqualFold(split[0], "bearer") {
				return InvalidAuthHeaderError
			}
			token := strings.TrimSpace(split[1])
			if token == "" {
				return InvalidAuthHeaderError
			}
			claims, err := util.ParseToken(cfg.JWT.Issuer, token, cfg.JWT.SignKey)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			// Store claims to context.
			c.Set(cfg.Ctx.Key, claims)
			return next(c)
		}

	}
}
