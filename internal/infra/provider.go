package infra

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/joisandresky/go-chi-clean-starter/internal/infra/configs"
	"github.com/joisandresky/go-chi-clean-starter/pkg/postgresdb"
	"github.com/joisandresky/go-chi-clean-starter/pkg/redisstore"
)

func LoggerProvider() *zap.SugaredLogger {
	// init logger
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync() // flushes buffer, if any

	return zapLogger.Sugar()
}

func ConfigProvider() *configs.Config {
	cfg, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}

func PgGormProvider(cfg *configs.Config, logger *zap.SugaredLogger) *gorm.DB {
	return postgresdb.InitGormPostgres(
		cfg.App.Environment,
		logger,
		cfg.DB.DBUsername,
		cfg.DB.DBPassword,
		cfg.DB.DBHost,
		cfg.DB.DBPort,
		cfg.DB.DBDatabase,
	)
}

func RedisProvider(cfg *configs.Config, logger *zap.SugaredLogger) *redis.Client {
	return redisstore.InitRedis(
		logger,
		cfg.Storage.RedisHost,
		cfg.Storage.RedisPort,
		cfg.Storage.RedisPassword,
		cfg.Storage.RedisDB,
	)
}
