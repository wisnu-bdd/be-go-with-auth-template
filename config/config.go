package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	// JWTSecret is used for signing JWT tokens
	JWTSecret []byte

	// Port is the port your app will listen on
	Port string

	// AllowedOrigins is the list of CORS-allowed origins
	AllowedOrigins []string

	// MongoDB database name
	MongoDatabaseName string

	// MongoDB collection string
	MongoConnectionString string
)

// Load loads environment variables into config variables
func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	JWTSecret = []byte(getEnv("JWT_SECRET", "default-secret-change-this"))
	Port = getEnv("PORT", "8080")

	origins := getEnv("ALLOWED_ORIGINS", "http://localhost:5173")
	AllowedOrigins = parseOrigins(origins)

	MongoDatabaseName = getEnv("MONGODB_NAME", "")

	MongoConnectionString = getEnv("MONGODB_CONNECTION_STRING", "")

	if string(JWTSecret) == "default-secret-change-this" {
		log.Println("Warning: Using default JWT secret. Set JWT_SECRET in .env for production.")
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func parseOrigins(origins string) []string {
	parts := strings.Split(origins, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}
