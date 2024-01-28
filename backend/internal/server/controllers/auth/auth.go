package auth

import "github.com/labstack/echo/v4"

func Setup(e *echo.Group) *echo.Group {
	// Add routes under /auth
	auth := e.Group("/auth")
	auth.POST("/login", Login)
	auth.POST("/register", Register)
	return e
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept  json
// @Produce  json
func Login(c echo.Context) error {
	return c.JSON(200, "login")
}

func Register(c echo.Context) error {
	return c.JSON(200, "register")
}
