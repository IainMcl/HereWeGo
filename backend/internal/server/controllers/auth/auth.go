package auth

import (
	"net/http"

	_ "github.com/IainMcl/HereWeGo/docs"
	data "github.com/IainMcl/HereWeGo/internal/data/user"
	models "github.com/IainMcl/HereWeGo/internal/models/user"
	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Group) *echo.Group {
	// Add routes under /auth
	auth := e.Group("/auth")
	auth.POST("/login", Login)
	auth.POST("/register", Register)
	return e
}

type loginRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token          string `json:"token"`
	TokenExpiresAt string `json:"token_expires_at"`
}

type registerRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	UserName string `json:"username"`
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
	return c.JSON(http.StatusOK, "login")
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
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	err := data.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}
	return c.JSON(http.StatusOK, registerResponse{UserName: user.UserName})
}

// Logout godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Success 200
// @Router /auth/logout [post]
func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, "logout")
}
