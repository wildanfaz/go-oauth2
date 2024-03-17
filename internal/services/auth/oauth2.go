package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/go-template/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type ImplementServices struct {
	config *configs.Config
	log    *logrus.Logger
}

type Services interface {
	GoogleLogin(c echo.Context) error
	GoogleCallback(c echo.Context) error
}

func New(config *configs.Config, log *logrus.Logger) Services {
	return &ImplementServices{
		config: config,
		log:    log,
	}
}

func (s *ImplementServices) SetupConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     s.config.ClientID,
		ClientSecret: s.config.ClientSecret,
		RedirectURL:  s.config.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
