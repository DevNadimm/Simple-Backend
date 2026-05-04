package db

import (
	"fmt"
	"test/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cfg *config.DBConfig) string {
	if cfg.ConnectionURL != "" {
		return cfg.ConnectionURL
	}

	sslMode := "disable"
	if cfg.EnableSSLMode {
		sslMode = "require" // "require" is better for most cloud providers than "enable"
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		sslMode,
	)

	return connStr
}

func NewConnection(cfg *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cfg)
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
