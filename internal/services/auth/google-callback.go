package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/go-oauth2/internal/helpers"
	"golang.org/x/oauth2"
)

func (s *ImplementServices) GoogleCallback(c echo.Context) error {
	var (
		resp = helpers.NewResponse()
	)

	state := c.Request().URL.Query().Get("state")
	if state != "random" {
		s.log.Warnln("Invalid state")
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid state"))
	}

	code := c.Request().URL.Query().Get("code")

	conf := s.SetupConfig()

	token, err := conf.Exchange(c.Request().Context(), code)
	if err != nil {
		s.log.Errorln("Failed to exchange token:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Invalid credentials"))
	}

	response, err := oauth2.NewClient(c.Request().Context(), conf.TokenSource(c.Request().Context(), token)).Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		s.log.Errorln("Failed to get user info:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Failed to get user info"))
	}

	userInfo, err := io.ReadAll(response.Body)
	if err != nil {
		s.log.Errorln("Failed to read user info:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Failed to read user info"))
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(userInfo, &data)
	if err != nil {
		s.log.Errorln("Failed to unmarshal user info:", err)
		return c.JSON(http.StatusBadRequest, resp.AsError().
			WithMessage("Failed to unmarshal user info"))
	}

	return c.JSON(http.StatusOK, resp.WithMessage("Success").
		WithData(data))
}
