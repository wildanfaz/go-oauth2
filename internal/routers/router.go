package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-template/configs"
	"github.com/wildanfaz/go-template/internal/pkg"
	oauth2_router "github.com/wildanfaz/go-template/internal/routers/oauth2-router"
	"github.com/wildanfaz/go-template/internal/services/auth"
	"github.com/wildanfaz/go-template/internal/services/health"
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
	oauth2_router.Oauth2Router(auth, authServices)

	e.Logger.Fatal(e.Start(config.AppPort))
}
