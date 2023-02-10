package config

import "os"

type DatabaseConfig struct {
	User     string
	Password string
	Database string
	Url      string
}

type Config struct {
	Db DatabaseConfig
}

func GetConfig() *Config {
	return &Config{Db: DatabaseConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Url:      os.Getenv("DB_URL"),
	}}
}
