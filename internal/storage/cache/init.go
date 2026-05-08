package cache

import (
	"log/slog"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
	"github.com/gofiber/storage/redis/v3"
)

func InitSession(cfg *config.Config, log *slog.Logger) (*redis.Storage, error) {

	storage := redis.New(redis.Config{
		Host: cfg.Redis.REDIS_HOST,
		Port: cfg.Redis.REDIS_PORT,
	})

	log.Info("Redis session storage initialized")

	return storage, nil
}
