package server

import (
	"net/http"

	"github.com/IainMcl/HereWeGo/internal/logging"
	custommiddleware "github.com/IainMcl/HereWeGo/internal/middleware"
	"github.com/IainMcl/HereWeGo/internal/server/controllers/auth"
	"github.com/IainMcl/HereWeGo/internal/settings"
	"github.com/IainMcl/HereWeGo/internal/util"
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
	r.Use(custommiddleware.JWTWithInvalidationCheck)

	r.GET("/health", s.healthHandler)
	r.Use(echojwt.WithConfig(config))

	// Anonymous routes
	a := e.Group("/api")
	a.GET("/ping", s.ping)

	auth.Setup(s.db.Db(), a, r)
	if settings.ServerSettings.RunMode == "debug" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	return e
}

// HealthHandler godoc
//
//	@Summary		Returns the health of the database server
//	@Description	Returns the health of the database server
//	@Tags			System
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Security		ApiKeyAuth
//	@Router			/health [get]
func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

// Ping godoc
//
//	@Summary		Ping
//	@Description	Ping
//	@Tags			System
//	@Produce		json
//	@Success		200
//	@Router			/ping [get]
func (s *Server) ping(c echo.Context) error {
	logging.Debug("ping")
	return c.String(http.StatusOK, "pong")
}
