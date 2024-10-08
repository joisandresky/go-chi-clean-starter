package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/joisandresky/go-chi-clean-starter/internal/application/usecases"
	"github.com/joisandresky/go-chi-clean-starter/internal/infra/configs"
	"github.com/joisandresky/go-chi-clean-starter/internal/infra/repositories"
	"github.com/joisandresky/go-chi-clean-starter/internal/presentation/api"
	"github.com/joisandresky/go-chi-clean-starter/internal/presentation/middleware"
)

func Inject(
	r *chi.Mux,
	cfg *configs.Config,
	logger *zap.SugaredLogger,
	gormdb *gorm.DB,
	redisClient *redis.Client,
) {
	/* Add your dependencies here */

	// repos
	postRepo := repositories.NewPgPostRepository(gormdb)

	// usecases
	postUc := usecases.NewPostUsecase(postRepo)

	// middlewares
	testMw := middleware.NewTestMiddleware(logger)

	// Routes Registration
	api.NewPostHttpApi(postUc).RegisterRoutes(r, testMw)
}
