package controllers

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/laughing-nerd/file-hive/models"
	"github.com/laughing-nerd/file-hive/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func UploadFile(c echo.Context) error {

	c.Request().ParseMultipartForm(1 << 30) // 1 GB limit

	id := c.Param("id")
	fmt.Println(id)

	file, err := c.FormFile("file")
	utils.HandleErr(err)

	var user models.User
	utils.Collection.FindOne(context.Background(), bson.D{{"_id",id}}).Decode(&user)

  if (file.Size / (1024 * 1024)) <= int64(user.TotalSize) 

	return nil
}
