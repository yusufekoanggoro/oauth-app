package config

import "os"

type Config struct {
	ClientID  string
	SecretKey string
}

func LoadConfig() *Config {
	return &Config{
		ClientID:  os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}
