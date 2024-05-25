package utils

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var MinioErrr error

func InitMinio(e *echo.Echo) {
	secure := true
	if os.Getenv("ENV") == "dev" {
		secure = false
	}

	MinioClient, MinioErrr = minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: secure,
	})

	if MinioErrr != nil {
		e.Logger.Fatal(MinioErrr)
	}
	fmt.Println("Connected to Minio")
}
