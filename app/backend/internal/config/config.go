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
		MongoURI:     getOr("MONGO_URI", "mongodb://localhost:27017/fitconnect"),
		JWTSecret:    getOr("JWT_SECRET", "dev-secret"),
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
