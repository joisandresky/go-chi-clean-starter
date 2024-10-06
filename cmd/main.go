package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/joisandresky/go-chi-clean-starter/internal/infra"
	"github.com/joisandresky/go-chi-clean-starter/internal/infra/configs"
	"github.com/joisandresky/go-chi-clean-starter/pkg/postgresdb"
	"github.com/joisandresky/go-chi-clean-starter/pkg/redisstore"
)

func main() {
	// init logger
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync() // flushes buffer, if any

	logger := zapLogger.Sugar()

	// Config
	cfg, err := configs.LoadConfig()
	if err != nil {
		logger.Errorf("failed to load config: %v", err)
		os.Exit(1)
	}

	// postgres gorm
	gormdb := postgresdb.InitGormPostgres(
		cfg.App.Environment,
		logger,
		cfg.DB.DBUsername,
		cfg.DB.DBPassword,
		cfg.DB.DBHost,
		cfg.DB.DBPort,
		cfg.DB.DBDatabase,
	)

	// redis
	redisClient := redisstore.InitRedis(
		logger,
		cfg.Storage.RedisHost,
		cfg.Storage.RedisPort,
		cfg.Storage.RedisPassword,
		cfg.Storage.RedisDB,
	)

	server := infra.NewServer(cfg, logger, gormdb, redisClient)

	server.Run()
}
