package redisstore

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis(logger *zap.SugaredLogger, host, port, password string, dbNumber int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			host,
			port,
		), // Replace with the actual address of your Redis server
		Password: password, // no password set
		DB:       dbNumber, // use default DB
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		logger.Errorf("failed to connect into redis: %v", err)
		os.Exit(1)
	} else {
		logger.Info("Successfully connected to Redis")
	}

	return redisClient
}
