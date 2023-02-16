package config

import "os"

// DatabaseConfig is the struct to hold the parameters to the db connection
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// Config holds the app configuration
type Config struct {
	Db DatabaseConfig
	MessageBroker string
}

// GetConfig fills the configuration and return a Config object
func GetConfig() *Config {
	return &Config{
		Db: DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Host:  os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		},
		MessageBroker: os.Getenv("MB_URL"),
	}
}
