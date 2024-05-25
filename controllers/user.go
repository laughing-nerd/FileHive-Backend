package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/laughing-nerd/file-hive/models"
	"github.com/laughing-nerd/file-hive/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")
	res, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + token)
	utils.HandleErr(err)
	defer res.Body.Close()

	var user models.User
	jsonErrr := json.NewDecoder(res.Body).Decode(&user)
	utils.HandleErr(jsonErrr)

	var found models.User

	utils.Collection.FindOne(context.Background(), bson.D{{"email", user.Email}}).Decode(&found)

	return c.JSON(http.StatusOK, found)
}
