package initializer

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port string
	ENV string
}

type DBConfig struct {
	Host string
	Port string
	User string
	Password string
	Name string
	SSLMode string
}

type JWTConfig struct {
	Secret string
	ExpiryHour int
}

type Config struct {
	Server ServerConfig
	DB DBConfig
	JWT JWTConfig
}

// Global variable untuk menampung config
var AppConfig Config

func LoadConfig() {
	// Gunakan Load tanpa argument jika .env ada di root
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	AppConfig.Server.Port = getEnv("APP_PORT", "3000")
	AppConfig.Server.ENV = getEnv("APP_ENV", "development")

	AppConfig.DB.Host = getEnv("DB_HOST", "localhost")
	AppConfig.DB.Port = getEnv("DB_PORT", "5432")
	AppConfig.DB.User = getEnv("DB_USER", "postgres")
	AppConfig.DB.Password = getEnv("DB_PASSWORD", "")
	AppConfig.DB.Name = getEnv("DB_NAME", "algora_db")
	AppConfig.DB.SSLMode = getEnv("DB_SSLMODE", "disable")

	AppConfig.JWT.Secret = os.Getenv("JWT_SECRET")
	if AppConfig.JWT.Secret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	expiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOUR", "24"))
	AppConfig.JWT.ExpiryHour = expiry
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
