package auth

import (
	"net/http"

	_ "github.com/IainMcl/HereWeGo/docs"
	models "github.com/IainMcl/HereWeGo/internal/models/user"
	"github.com/IainMcl/HereWeGo/internal/util"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var db *sqlx.DB

func Setup(d *sqlx.DB, a, r *echo.Group) (*echo.Group, *echo.Group) {
	db = d
	// Add restricted routes under /auth
	authRes := r.Group("/auth")
	// Add anonymous routes under /auth
	authAnon := a.Group("/auth")
	authAnon.POST("/login", Login)
	authAnon.POST("/register", Register)
	authRes.POST("/logout", Logout)
	return authAnon, authRes
}

type LoginRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterResponse struct {
	UserName string `json:"username"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			Auth
//	@Accept			json
//	@Param			body	body	LoginRequest	true	"Login Request"
//	@Produce		json
//	@Success		200	{object}	LoginResponse
//	@Router			/auth/login [post]
func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email and password are required"})
	}

	u, err := models.AuthenticateUser(db, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
	}

	token, err := util.GenerateToken(u.Username, u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

// Register godoc
//
//	@Summary		Register
//	@Description	Register
//	@Tags			Auth
//	@Accept			json
//	@Param			body	body	RegisterRequest	true	"Register Request"
//	@Success		200
//	@Router			/auth/register [post]
func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	err := user.CreateUser(db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
	}
	return c.JSON(http.StatusOK, RegisterResponse{UserName: user.Username})
}

// Logout godoc
//
//	@Summary		Logout
//	@Description	Logout
//	@Tags			Auth
//	@Success		200
//	@Security		ApiKeyAuth
//	@Router			/auth/logout [post]
func Logout(c echo.Context) error {
	// Add token to invalid list
	token := c.Request().Header.Get("Authorization")[7:]
	util.AddBlacklist(token)
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out"})
}

// ResetPassword godoc
//
//	@Summary		Reset Password
//	@Description	Reset Password
//	@Tags			Auth
//	@Accept			json
//	@Param			body	body	ResetPasswordRequest	true	"Reset Password Request"
//	@Success		200
//	@Router			/auth/resetpassword [post]
func ResetPassword(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Reset password"})
}

// SendResetPasswordEmail godoc
//
//	@Summary		Send Reset Password Email
//	@Description	Send Reset Password Email to the provided email if it exists
//	@Tags			Auth
//	@Accept			json
//	@Param			email	query	string	true	"Email to send reset password email to"
//	@Success		200
//	@Router			/auth/sendresetpasswordemail [get]
func SendResetPasswordEmail(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Send reset password email"})
}
