package repositories

import (
	"context"

	"github.com/joisandresky/go-chi-clean-starter/internal/domain/entities"
)

type PostRepository interface {
	GetAll(ctx context.Context) ([]entities.Post, error)
	GetById(ctx context.Context, id string) (*entities.Post, error)
	Create(ctx context.Context, post *entities.Post) error
	UpdateById(ctx context.Context, post *entities.Post) error
	DeleteById(ctx context.Context, id string) error
}
