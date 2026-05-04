package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User          string
	Password      string
	Host          string
	Port          int
	Name          string
	EnableSSLMode bool
	ConnectionURL string
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

var config *Config

func loadConfig() {
	_ = godotenv.Load() // Ignore error if .env is missing

	version := os.Getenv("VERSION")
	if version == "" {
		version = "1.0.0" // Default version
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		serviceName = "ECOMMERCE" // Default name
	}

	// Support both HTTP_PORT and PORT (used by Render)
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		httpPortStr = os.Getenv("PORT")
	}
	if httpPortStr == "" {
		httpPortStr = "3000" // Default port
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		fmt.Println("HTTP_PORT must be a number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		jwtSecretKey = "default_secret_key" // Should be changed in production
	}

	// --- Database config ---
	dbURL := os.Getenv("DATABASE_URL")
	var dbConfig *DBConfig

	if dbURL != "" {
		// If DATABASE_URL is provided, we use it directly
		// We'll store it in the Host field or add a new field if needed.
		// For now, let's just make sure NewConnection can handle it.
		dbConfig = &DBConfig{
			ConnectionURL: dbURL,
		}
	} else {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPortStr := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		if dbUser == "" || dbHost == "" || dbName == "" {
			fmt.Println("Database configuration (DB_USER, DB_HOST, DB_NAME) or DATABASE_URL is required")
			// We don't exit here to allow the app to start, but DB calls will fail
		}

		dbPort := 5432
		if dbPortStr != "" {
			dbPort, _ = strconv.Atoi(dbPortStr)
		}

		dbSSLStr := os.Getenv("DB_ENABLE_SSL_MODE")
		enableSSL := false
		if strings.ToLower(dbSSLStr) == "true" {
			enableSSL = true
		}

		dbConfig = &DBConfig{
			User:          dbUser,
			Password:      dbPassword,
			Host:          dbHost,
			Port:          dbPort,
			Name:          dbName,
			EnableSSLMode: enableSSL,
		}
	}

	config = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     httpPort,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config {
	if config == nil {
		loadConfig()
	}
	return config
}
