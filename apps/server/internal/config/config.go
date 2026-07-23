package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Dsn  string

	JwtSecretKey string
	JwtDuration  string

	FrontendUrl string

	SuperAdminName     string
	SuperAdminEmail    string
	SuperAdminPassword string

	SmtpHost      string
	SmtpPort      string
	SmtpUsername  string
	SmtpPassword  string
	SmtpFromEmail string

	CloudinaryCloudName string
	CloudinaryApiKey    string
	CloudinaryApiSecret string
}

var AppConfig *Config

func LoadEnv() (*Config, error) {
	_ = godotenv.Load()
	_ = godotenv.Load("../.env")
	_ = godotenv.Load("../../.env")
	_ = godotenv.Load("../../../.env")

	AppConfig = &Config{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DSN"),

		JwtSecretKey: os.Getenv("JWT_SECRETKEY"),
		JwtDuration:  os.Getenv("JWT_DURATION"),

		FrontendUrl: os.Getenv("FRONTEND_URL"),

		SuperAdminName:     os.Getenv("SUPER_ADMIN_NAME"),
		SuperAdminEmail:    os.Getenv("SUPER_ADMIN_EMAIL"),
		SuperAdminPassword: os.Getenv("SUPER_ADMIN_PASSWORD"),

		SmtpHost:      os.Getenv("SMTP_HOST"),
		SmtpPort:      os.Getenv("SMTP_PORT"),
		SmtpUsername:  os.Getenv("SMTP_USERNAME"),
		SmtpPassword:  os.Getenv("SMTP_PASSWORD"),
		SmtpFromEmail: os.Getenv("SMTP_FROM_EMAIL"),

		CloudinaryCloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
	}

	return AppConfig, nil
}
