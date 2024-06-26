
package config

import (
    "database/sql"
    "github.com/go-redis/redis/v8"
    "github.com/joho/godotenv"
    "log"
    "os"
)

type Config struct {
    ServerPort string
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    RedisHost  string
    RedisPort  string
}

func LoadConfig() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }

    return &Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "user"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "gocool"),
        RedisHost:  getEnv("REDIS_HOST", "localhost"),
        RedisPort:  getEnv("REDIS_PORT", "6379"),
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func NewDBClient(cfg *Config) *sql.DB {
    dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
                       cfg.DBUser, cfg.DBPassword, cfg.DBName)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    return db
}

func NewRedisClient(cfg *Config) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
        Password: "",  // no password set
        DB:       0,  // use default DB
    })
    return rdb
}
