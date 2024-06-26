
package config

import (
    "os"
    "log"

    "github.com/joho/godotenv"
)

// Config stores all configuration of the application.
// The values are read by godotenv from a .env file or environment variables.
type Config struct {
    ServerPort string
    RedisHost  string
    RedisPort  string
}

// LoadConfig reads configuration from file or environment and returns it.
func LoadConfig() (*Config, error) {
    // Load .env file if it exists
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found, reading configuration from environment")
    }

    return &Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
        RedisHost:  getEnv("REDIS_HOST", "localhost"),
        RedisPort:  getEnv("REDIS_PORT", "6379"),
    }, nil
}

// getEnv reads an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
