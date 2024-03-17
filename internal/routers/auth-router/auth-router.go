package auth_router

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-oauth2/internal/services/auth"
)

func AuthRouter(g *echo.Group, authServices auth.Services) {
	g.GET("/google/login", authServices.GoogleLogin)
	g.GET("/google/callback", authServices.GoogleCallback)
}
