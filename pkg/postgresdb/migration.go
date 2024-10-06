package postgresdb

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func AutoMigration(dbUrl string, logger *zap.SugaredLogger) error {
	dbUrl = fmt.Sprintf("%s?sslmode=disable", dbUrl)
	logger.Info("Prepareing DB Migration for NON-Local ENV")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		logger.Errorf("an error occurred when opening connection: %v", err)
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		logger.Errorf("an error occurred when opening connection driver: %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)
	if err != nil {
		logger.Errorf(
			"an error occurred when connecting migration folder and database instance: %v",
			err,
		)
		return err
	}

	logger.Info("Doing Migration Up!")
	return m.Up()
}
