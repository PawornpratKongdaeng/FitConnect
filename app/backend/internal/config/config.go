package config

import (
	"os"
)

type Config struct {
	MongoURI     string
	JWTSecret    string
	Port         string
	ClientOrigin string
}

func Load() Config {
	return Config{
		MongoURI:     os.Getenv("MONGO_URI"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		Port:         getOr("PORT", "8080"),
		ClientOrigin: getOr("CLIENT_ORIGIN", "http://localhost:5173"),
	}
}

func getOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
