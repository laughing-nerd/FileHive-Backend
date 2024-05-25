package controllers

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/laughing-nerd/file-hive/config"
)

func GoogleLogin(c echo.Context) error {

	googleConfig := config.SetupGoogleConfig()
	url := googleConfig.AuthCodeURL(os.Getenv("RANDOM_STATE"))
	return c.Redirect(http.StatusTemporaryRedirect, url)
}
