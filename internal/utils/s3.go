package utils

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Uploader struct { client *s3.S3 bucket string }

func NewS3Uploader() (*S3Uploader, error) { sess, err := session.NewSession(&aws.Config{ Region: aws.String("us-east-1"), Credentials: credentials.NewStaticCredentials("your-access-key", "your-secret-key", ""), }) if err != nil { return nil, err } return &S3Uploader{ client: s3.New(sess), bucket: "your-bucket-name", }, nil }

func (u *S3Uploader) Upload(fileName string, file []byte) (string, error) { _, err := u.client.PutObject(&s3.PutObjectInput{ Bucket: aws.String(u.bucket), Key: aws.String(fileName), Body: bytes.NewReader(file), ACL: aws.String("public-read"), }) if err != nil { return "", err } return "https://" + u.bucket + ".s3.amazonaws.com/" + fileName, nil }