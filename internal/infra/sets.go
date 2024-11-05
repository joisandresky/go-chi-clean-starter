package infra

import (
	"github.com/google/wire"

	"github.com/joisandresky/go-chi-clean-starter/internal/application/usecases"
	"github.com/joisandresky/go-chi-clean-starter/internal/infra/repositories"
	"github.com/joisandresky/go-chi-clean-starter/internal/presentation/api"
	"github.com/joisandresky/go-chi-clean-starter/internal/presentation/middleware"
)

var BaseSet = wire.NewSet(
	ConfigProvider,
	LoggerProvider,
	PgGormProvider,
	RedisProvider,
)

var RepositorySet = wire.NewSet(
	repositories.NewPgPostRepository,
)

var UsecaseSet = wire.NewSet(
	usecases.NewPostUsecase,
)

var MiddlewareSet = wire.NewSet(
	middleware.NewTestMiddleware,
)

var HandlerSet = wire.NewSet(
	api.NewPostHttpApi,

	api.SetupRoutes,
)

var AppSet = wire.NewSet(
	BaseSet,
	RepositorySet,
	UsecaseSet,
	MiddlewareSet,
	HandlerSet,

	NewServer,
)
