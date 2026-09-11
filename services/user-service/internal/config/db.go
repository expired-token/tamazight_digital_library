package config

import (
	"fmt"
	"os"
	"services/user-service/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
	SSLMode  string
}

func LoadDatabaseConfig() (DatabaseConfig, error) {
	cfg := DatabaseConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	if cfg.User == "" || cfg.Host == "" || cfg.DBName == "" {
		return DatabaseConfig{}, fmt.Errorf("Missing required database configs; DB_USER || DB_HOST || POSTGRES_DB")
	}

	return cfg, nil
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", c.User, c.Password, c.DBName, c.Host, c.Port, c.SSLMode)
}

func InitDB() (*gorm.DB, error) {
	cfg, err := LoadDatabaseConfig()

	if err != nil {
		return nil, fmt.Errorf("Failed to load database config: %w", err)
	}

	dsn := cfg.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, fmt.Errorf("Failed to run migrations: %w", err)
	}

	return db, nil
}
