package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-oauth2/configs"
	"github.com/wildanfaz/go-oauth2/internal/pkg"
	auth_router "github.com/wildanfaz/go-oauth2/internal/routers/auth-router"
	"github.com/wildanfaz/go-oauth2/internal/services/auth"
	"github.com/wildanfaz/go-oauth2/internal/services/health"
)

func InitEchoRouter() {
	// configs
	config := configs.InitConfig()

	// pkg
	log := pkg.InitLogger()

	// services
	authServices := auth.New(config, log)

	e := echo.New()

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/health", health.HealthCheck)

	auth := apiV1.Group("/auth")
	auth_router.AuthRouter(auth, authServices)

	e.Logger.Fatal(e.Start(config.AppPort))
}
