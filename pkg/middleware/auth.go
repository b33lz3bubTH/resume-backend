package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func RootKeyAuth(rootKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
			}

			if parts[1] != rootKey {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid root key")
			}

			return next(c)
		}
	}
}

