package utils

import (
	"context"

	"github.com/laughing-nerd/file-hive/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserFound(user models.User) bool {
	var found models.User

	Collection.FindOne(context.Background(), bson.D{{"email", user.Email}}).Decode(&found)
	if found.Email == "" {
		return false
	}
	return true
}
