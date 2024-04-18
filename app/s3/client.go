package s3

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewS3Client() *minio.Client {
	if err := godotenv.Load("aws.env"); err != nil {
		log.Fatal("error loading .env file:", err)
	}

	awsEndpoint := os.Getenv("AWS_ENDPOINT")
	awsRegion := os.Getenv("AWS_REGION")
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	useSsl := true

	client, err := minio.New(awsEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(awsAccessKeyID, awsSecretAccessKey, ""),
		Region: awsRegion,
		Secure: useSsl,
	})
	if err != nil {
		log.Println(err)
	}
	return client
}
