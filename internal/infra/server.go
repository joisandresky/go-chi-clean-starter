package infra

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/joisandresky/go-chi-clean-starter/internal/infra/configs"
)

type ServerBuilder struct {
	cfg         *configs.Config
	logger      *zap.SugaredLogger
	gormdb      *gorm.DB
	redisClient *redis.Client
	routers     *chi.Mux
}

func NewServer(
	cfg *configs.Config,
	logger *zap.SugaredLogger,
	gormdb *gorm.DB,
	redisClient *redis.Client,
	routers *chi.Mux,
) *ServerBuilder {
	return &ServerBuilder{
		cfg:         cfg,
		logger:      logger,
		gormdb:      gormdb,
		redisClient: redisClient,
		routers:     routers,
	}
}

func (srv *ServerBuilder) Run() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", srv.cfg.App.Port),
		Handler: srv.routers,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(
		sig,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()

			// Cleanly shutdown the server
			srv.logger.Info("shutting down server cause deadline exceeded..")
			if shutdownCtx.Err() == context.DeadlineExceeded {
				srv.logger.Error("graceful shutdown timed out.. forcing exit")
			}
		}()

		srv.logger.Info("preparing server shutdown...")
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			srv.logger.Error(err)
		}

		srv.logger.Info("server shutdown completed")
		serverStopCtx()
	}()

	srv.logger.Info("Server is ready to listen and serve on port ", srv.cfg.App.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		srv.logger.Error(err)
	}

	srv.logger.Info("shutting down..")

	<-serverCtx.Done()
}
