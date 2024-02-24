package auth

import (
	"fmt"
	"net/http"
	"strings"

	_ "github.com/IainMcl/HereWeGo/docs"
	"github.com/IainMcl/HereWeGo/internal/logging"
	models "github.com/IainMcl/HereWeGo/internal/models/user"
	"github.com/IainMcl/HereWeGo/internal/services"
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
	authRes.GET("/username", Username)
	authAnon.POST("/register", Register)
	authRes.POST("/logout", Logout)
	authAnon.GET("/sendresetpasswordemail", SendResetPasswordEmail)
	authAnon.POST("/resetpassword", ResetPassword)
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
	Email    string `json:"email"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
	Token       string `json:"token"`
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

	token, err := util.GenerateToken(u.Username, u.Email, u.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

// Username godoc
//
//	@Summary		Get Username
//	@Description	Get Username as an authentication check for the logged in user
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Security		ApiKeyAuth
//	@Router			/auth/username [get]
func Username(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	user, err := models.GetUserFromToken(db, token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}
	return c.JSON(http.StatusOK, map[string]string{"username": user.Username})
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

	fmt.Println(err)
	// var pgErr *pgconn.PgError
	// if errors.As(err, &pgErr) {
	// 	if err.(*pgconn.PgError).Code == "23505" {
	// 		return c.JSON(http.StatusConflict, map[string]string{"message": "Email already exists"})
	// 	}
	// }
	if strings.Contains(err.Error(), "23505") {
		return c.JSON(http.StatusConflict, map[string]string{"message": "Email already exists"})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, RegisterResponse{UserName: user.Username, Email: user.Email})
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
	var req ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	if req.Email == "" || req.NewPassword == "" || req.Token == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email, new password and token are required"})
	}

	token, err := util.GetTokenFromDB(db, req.Token, req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
	}
	if token.IsExpired() {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token expired"})
	}

	user, err := models.GetUserByEmail(db, req.Email)
	if err != nil {
		logging.Fatal("Error getting user from db after using valid OTP token: ", err)
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Email not found"})
	}

	err = util.UseToken(db, req.Token, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to use token"})
	}

	err = user.UpdatePassword(db, req.NewPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update password"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Passowrd updated successfully"})
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
func SendResetPasswordEmail(c echo.Context) error { // Verify that the email exists in the database
	email := c.QueryParam("email")
	if !services.ValidateEmail(email) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid email"})
	}
	user, err := models.GetUserByEmail(db, email)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Email not found"})
	}

	token, err := util.NewOTPToken(db, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
	}
	if err := services.SendResetPasswordEmail(user.Email, user.Username, user.ID, token); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to send reset password email"})
	}
	obscuredEmail := services.ObscureEmail(email)
	return c.JSON(http.StatusOK, map[string]string{"message": "Reset password email sent to " + obscuredEmail})
}
