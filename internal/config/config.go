package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"styleai/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct { DB *gorm.DB S3Uploader *utils.S3Uploader }

func LoadConfig() (*Config, error) { // Подключение к PostgreSQL dsn := "host=localhost user=postgres password=yourpassword dbname=styleai port=5432 sslmode=disable" db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) if err != nil { return nil, err }

	// Инициализация S3
	s3Uploader, err := utils.NewS3Uploader()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB:        db,
		S3Uploader: s3Uploader,
	}, nil

}