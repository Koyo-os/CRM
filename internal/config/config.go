package config

import (
	"os"
)

func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }

    return defaultVal
}

type Config struct{
	Port string
	Host string
	MongoURL string
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Host: getEnv("HOST", "localhost"),
		MongoURL: getEnv("MONGO_URL", "mongodb://localhost:27017"),
	}
}