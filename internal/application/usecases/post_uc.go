package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/joisandresky/go-chi-clean-starter/internal/application/dto"
	"github.com/joisandresky/go-chi-clean-starter/internal/domain/entities"
	"github.com/joisandresky/go-chi-clean-starter/internal/domain/repositories"
)

type PostUsecase interface {
	GetAll(ctx context.Context) ([]entities.Post, error)
	GetById(ctx context.Context, id string) (*entities.Post, error)
	Create(ctx context.Context, req *dto.CreatePost) error
	UpdateById(ctx context.Context, id string, req *dto.CreatePost) error
	DeleteById(ctx context.Context, id string) error
}

type postUc struct {
	repo repositories.PostRepository
}

func NewPostUsecase(repo repositories.PostRepository) PostUsecase {
	return &postUc{repo: repo}
}

func (uc *postUc) GetAll(ctx context.Context) ([]entities.Post, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *postUc) GetById(ctx context.Context, id string) (*entities.Post, error) {
	return uc.repo.GetById(ctx, id)
}

func (uc *postUc) Create(ctx context.Context, req *dto.CreatePost) error {
	post := &entities.Post{
		ID:        uuid.New(),
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return uc.repo.Create(ctx, post)
}

func (uc *postUc) UpdateById(ctx context.Context, id string, req *dto.CreatePost) error {
	post, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	post.Title = req.Title
	post.Body = req.Body
	post.UpdatedAt = time.Now()

	return uc.repo.UpdateById(ctx, post)
}

func (uc *postUc) DeleteById(ctx context.Context, id string) error {
	post, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	return uc.repo.DeleteById(ctx, post.ID.String())
}
