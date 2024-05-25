package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func InitDB(e *echo.Echo) {
	mongo_uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))
	if err != nil {
		e.Logger.Fatal(err)
	}

	Collection = client.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	fmt.Println("Connected to MongoDB")
}
