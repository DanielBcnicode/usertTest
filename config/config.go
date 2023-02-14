package config

import "os"

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Db DatabaseConfig
	MessageBroker string
}

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
