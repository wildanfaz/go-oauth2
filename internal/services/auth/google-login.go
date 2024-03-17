package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *ImplementServices) GoogleLogin(c echo.Context) error {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	conf := s.SetupConfig()

	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("random")

	return c.Redirect(http.StatusTemporaryRedirect, url)
}
