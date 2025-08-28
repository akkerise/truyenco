package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port       string
    DatabaseURL string
    RedisAddr  string
    KafkaAddr  string
    JWTSecret  string
}

func Load() Config {
    _ = godotenv.Load()
    return Config{
        Port:        getEnv("API_PORT", "8080"),
        DatabaseURL: getEnv("DB_DSN", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
        RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
        KafkaAddr:   getEnv("KAFKA_ADDR", "localhost:9092"),
        JWTSecret:   getEnv("JWT_SECRET", "secret"),
    }
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}

