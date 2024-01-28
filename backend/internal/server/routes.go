package server

import (
	"net/http"

	"github.com/IainMcl/HereWeGo/internal/logging"
	"github.com/IainMcl/HereWeGo/internal/server/controllers/auth"
	"github.com/IainMcl/HereWeGo/internal/settings"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	if settings.ServerSettings.RunMode != "debug" {
		e.Logger.SetOutput(logging.F)
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

	api := e.Group("/api")
	api.GET("/", s.HelloWorldHandler)
	api.GET("/health", s.healthHandler)

	// authresp := e.Group("/auth")
	// authresp.POST("/login", auth.Login)
	// Add routes under /auth
	api = auth.Setup(api)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}
