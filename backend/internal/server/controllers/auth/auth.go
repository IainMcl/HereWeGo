package auth

import (
	_ "github.com/IainMcl/HereWeGo/docs"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Group) *echo.Group {
	// Add routes under /auth
	auth := e.Group("/auth")
	auth.POST("/login", Login)
	auth.POST("/register", Register)
	return e
}

type LoginRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token          string `json:"token"`
	TokenExpiresAt string `json:"token_expires_at"`
}

type RegisterRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Param body body LoginRequest true "Login Request"
// @Produce json
// @Success 200 {object} LoginResponse
// @Router /auth/login [post]
func Login(c echo.Context) error {
	return c.JSON(200, "login")
}

// Register godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Param body body RegisterRequest true "Register Request"
// @Success 200
// @Router /auth/register [post]
func Register(c echo.Context) error {
	return c.JSON(200, "register")
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Success 200
// @Router /auth/logout [post]
func Logout(c echo.Context) error {
	return c.JSON(200, "logout")
}
