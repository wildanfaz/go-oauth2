package oauth2_router

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-template/internal/services/auth"
)

func Oauth2Router(g *echo.Group, authServices auth.Services) {
	g.GET("/google/login", authServices.GoogleLogin)
	g.GET("/google/callback", authServices.GoogleCallback)
}
