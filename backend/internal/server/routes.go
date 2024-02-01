package server

import (
	"fmt"
	"net/http"

	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/server/controllers/auth"
	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/IainMcl/HereWeGo/internal/util"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/IainMcl/HereWeGo/docs"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	if settings.ServerSettings.RunMode != "debug" {
		e.Logger.SetOutput(logging.F)
	} else {
		e.Debug = true
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	enableCors := settings.AppSettings.EnableCors
	if enableCors {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		}))
	}

	// Restricted routes
	r := e.Group("/api")
	config := echojwt.Config{
		SigningKey: util.JwtSecret,
	}
	r.Use(echojwt.WithConfig(config))

	r.GET("/test", func(c echo.Context) error {
		// Get the claims from the jwt
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims) // Use jwt.MapClaims
		name := claims["name"].(string)       // Access the "name" field from jwt.MapClaims
		return c.String(http.StatusOK, fmt.Sprintf("Welcome %s!", name))
	})

	// Anonymous routes
	a := e.Group("/api")
	a.GET("/", s.HelloWorldHandler)
	a.GET("/health", s.healthHandler)

	auth.Setup(a, r)
	if settings.ServerSettings.RunMode == "debug" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	return e
}

// HelloWorldHandler godoc
// @Summary Returns a hello world message
// @Description Returns a hello world message
// @Tags System
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router / [get]
func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

// HealthHandler godoc
// @Summary Returns the health of the database server
// @Description Returns the health of the database server
// @Tags System
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
