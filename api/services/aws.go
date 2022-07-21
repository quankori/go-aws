package services

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/quankori/go-aws/configs"
)

const (
	S3_REGION = "ap-south-1"
	S3_BUCKET = "test-node-kori"
	// S3_ACL    = "public-read"
)

type S3Handler struct {
	Session *session.Session
	Bucket  string
}

// Hostname get the host name
func S3(files *multipart.FileHeader) *s3.PutObjectOutput {
	config, _ := configs.LoadConfig()

	// Session S3
	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(S3_REGION),
			Credentials: credentials.NewStaticCredentials(
				config.AWSPublicId,
				config.AWSSecretKey,
				""),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// File name custom
	spl := strings.Split(files.Filename, ".")
	uploadedFileName := strings.Join(spl, ".")
	readFile, _ := files.Open()

	// Header
	var size int64 = files.Size
	buffer := make([]byte, size)
	readFile.Read(buffer)

	file, err := s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(S3_BUCKET),
		Key:         aws.String("new_root/" + uploadedFileName),
		Body:        bytes.NewReader(buffer),
		ContentType: aws.String(http.DetectContentType(buffer)),
	})
	return file
}
