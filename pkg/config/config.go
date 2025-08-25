package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv    string
	Port      string
	DBUser    string
	DBPass    string
	DBHost    string
	DBPort    string
	DBName    string
	JWTSecret string
}

func Load() *Config {
	_ = godotenv.Load() // no-op if .env missing
	cfg := &Config{
		AppEnv:    get("APP_ENV", "development"),
		Port:      get("PORT", "3000"),
		DBUser:    get("DB_USER", "youruser"),
		DBPass:    get("DB_PASS", "youruserpassword"),
		DBHost:    get("DB_HOST", "127.0.0.1"),
		DBPort:    get("DB_PORT", "3306"),
		DBName:    get("DB_NAME", "yourdb"),
		JWTSecret: get("JWT", "thisisasecret"),
	}
	return cfg
}

func get(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
