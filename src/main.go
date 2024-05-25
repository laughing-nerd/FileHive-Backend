package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laughing-nerd/file-hive/controllers"
	"github.com/laughing-nerd/file-hive/utils"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {

	e := echo.New()
	go utils.InitDB(e)
	go utils.InitMinio(e)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	group := e.Group("/api/v1")

	// Routes
	group.GET("/", controllers.Home)
	group.GET("/user", controllers.GetUser)
  group.POST("/upload/:id", controllers.UploadFile)

	// Google Login
	group.GET("/auth/google", controllers.GoogleLogin)
	group.GET("/google/callback", controllers.AuthGoogle)

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
