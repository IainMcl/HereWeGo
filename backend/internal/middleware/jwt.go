package custommiddleware

import (
	"net/http"
	"strings"

	"github.com/IainMcl/HereWeGo/internal/util"
	"github.com/labstack/echo/v4"
)

func JWTWithInvalidationCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") || len(authHeader) < 8 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		token := strings.Split(authHeader, "Bearer ")[1]
		// Check if the token is blacklisted
		if util.IsBlacklisted(token) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		return next(c)
	}
}
