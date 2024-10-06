package postgresdb

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormPostgres(
	app_env string,
	logger *zap.SugaredLogger,
	DBUsername string,
	DBPassword string,
	DBHost string,
	DBPort string,
	DBDatabase string,
) *gorm.DB {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		DBUsername,
		DBPassword,
		DBHost,
		DBPort,
		DBDatabase,
	)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect with database", err)
		os.Exit(1)
	}

	logger.Info("Database Connected!")

	// Do Auto Migration only for NON-LOCAL ENV
	// for LOCAL ENV please use golang migrate tool: https://github.com/golang-migrate/migrate
	// and do Migration using Makefile for convinient
	if app_env != "local" {
		if err := AutoMigration(dbUrl, logger); err != nil {
			logger.Errorf("failed to do Migration Up: %v", err)
		} else {
			logger.Info("Migration Up Done!")
		}
	}

	return db
}
