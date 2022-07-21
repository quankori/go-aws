package services

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
func S3(files *multipart.FileHeader) string {
	// config, _ := configs.LoadConfig()

	// Session S3
	session, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatal(err)
	}

	// File name custom
	// Destination
	spl := strings.Split(files.Filename, ".")
	uploadedFileName := strings.Join(spl, ".")
	readFile, _ := files.Open()
	var fileSize int64 = files.Size
	fileBuffer := make([]byte, fileSize)

	bug, err := s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String("new_root/" + uploadedFileName),
		// ACL:                  aws.String("private"),
		Body:        readFile,
		ContentType: aws.String(http.DetectContentType(fileBuffer)),
		// ContentDisposition:   aws.String("attachment"),
		// ServerSideEncryption: aws.String("AES256"),
	})
	fmt.Println(bug)
	return ""
}
