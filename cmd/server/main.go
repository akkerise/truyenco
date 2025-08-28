package main

import (
    "log"

    "github.com/gin-gonic/gin"

    "truyenco/internal/cache"
    "truyenco/internal/config"
    "truyenco/internal/database"
    "truyenco/internal/handlers"
    kaf "truyenco/internal/kafka"
)

func main() {
    cfg := config.Load()
    db := database.Connect(cfg.DatabaseURL)
    cacheClient := cache.New(cfg.RedisAddr)
    _ = cacheClient
    producer := kaf.NewWriter(cfg.KafkaAddr, "events")
    defer func() { _ = producer.Close() }()

    r := gin.Default()
    handlers.RegisterRoutes(r, db, cfg.JWTSecret)

    log.Fatal(r.Run(":" + cfg.Port))
}

