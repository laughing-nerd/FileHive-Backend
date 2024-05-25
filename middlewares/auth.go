package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("token")
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		return next(c)
	}
}
