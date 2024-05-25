package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/laughing-nerd/file-hive/config"
	"github.com/laughing-nerd/file-hive/models"
	"github.com/laughing-nerd/file-hive/utils"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthGoogle(c echo.Context) error {
	state := c.QueryParam("state")
	if state != os.Getenv("RANDOM_STATE") {
		return c.JSON(http.StatusBadRequest, "Invalid state")
	}

	conf := config.SetupGoogleConfig()
	info, err := conf.Exchange(context.Background(), c.QueryParam("code"))
	utils.HandleErr(err)

	res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + info.AccessToken)
	utils.HandleErr(err)
	defer res.Body.Close()
	var user models.User
	jsonErr := json.NewDecoder(res.Body).Decode(&user)
	utils.HandleErr(jsonErr)

	found := utils.UserFound(user)

	if !found {
		inserted, insertedErr := utils.Collection.InsertOne(context.Background(), user)
		utils.HandleErr(insertedErr)

		id := inserted.InsertedID.(primitive.ObjectID).Hex()
		err := utils.MinioClient.MakeBucket(context.Background(), id, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			utils.HandleErr(err)
		}
	}

	c.SetCookie(&http.Cookie{
		Name:   "token",
		Value:  info.AccessToken,
		Path:   "/",
		Domain: "localhost",
	})
	return c.Redirect(http.StatusFound, "http://localhost:3000/")
}
