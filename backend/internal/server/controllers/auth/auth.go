package auth

import (
	"net/http"

	_ "github.com/IainMcl/HereWeGo/docs"
	models "github.com/IainMcl/HereWeGo/internal/models/user"
	"github.com/IainMcl/HereWeGo/internal/util"
	"github.com/labstack/echo/v4"
)

func Setup(a, r *echo.Group) (*echo.Group, *echo.Group) {
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
	UserName string `json:"username"`
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

// Login godoc
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
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request:  email and password required")
	}

	u, err := models.AuthenticateUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	token, err := util.GenerateToken(u.Username, u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: token})
}

// Register godoc
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
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	err := user.CreateUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}
	return c.JSON(http.StatusOK, RegisterResponse{UserName: user.Username})
}

// Logout godoc
//	@Summary		Logout
//	@Description	Logout
//	@Tags			Auth
//	@Success		200
//	@Router			/auth/logout [post]
func Logout(c echo.Context) error {
	return c.JSON(http.StatusOK, "logout")
}
