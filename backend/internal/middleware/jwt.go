package custommiddleware

import (
	"net/http"

	"github.com/IainMcl/HereWeGo/internal/util"
	"github.com/labstack/echo/v4"
)

func JWTWithInvalidationCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")[7:]
		// Check if the token is blacklisted
		if util.IsBlacklisted(token) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}
		return next(c)
	}
}
