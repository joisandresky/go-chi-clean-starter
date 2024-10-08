package usecases

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/joisandresky/go-chi-clean-starter/internal/application/dto"
	"github.com/joisandresky/go-chi-clean-starter/internal/domain/entities"
	"github.com/joisandresky/go-chi-clean-starter/internal/domain/repositories"
	"github.com/joisandresky/go-chi-clean-starter/pkg/guy"
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
	posts, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, guy.NewAppError(
			http.StatusInternalServerError,
			"failed to get posts",
			err.Error(),
		)
	}

	return posts, nil
}

func (uc *postUc) GetById(ctx context.Context, id string) (*entities.Post, error) {
	if id == "" {
		return nil, guy.NewAppError(
			http.StatusBadRequest,
			"invalid post id",
			"",
		)
	}

	post, err := uc.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, guy.NewAppError(
				http.StatusNotFound,
				fmt.Sprintf("post with id %s not found", id),
				"",
			)
		}

		return nil, guy.NewAppError(
			http.StatusInternalServerError,
			"failed to get post",
			err.Error(),
		)
	}
	return post, nil
}

func (uc *postUc) Create(ctx context.Context, req *dto.CreatePost) error {
	post := &entities.Post{
		ID:        uuid.New(),
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.repo.Create(ctx, post); err != nil {
		return guy.NewAppError(http.StatusInternalServerError, "failed to create post", err.Error())
	}

	return nil
}

func (uc *postUc) UpdateById(ctx context.Context, id string, req *dto.CreatePost) error {
	post, err := uc.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	post.Title = req.Title
	post.Body = req.Body
	post.UpdatedAt = time.Now()

	if err := uc.repo.UpdateById(ctx, post); err != nil {
		return guy.NewAppError(http.StatusInternalServerError, "failed to update post", err.Error())
	}

	return nil
}

func (uc *postUc) DeleteById(ctx context.Context, id string) error {
	post, err := uc.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return guy.NewAppError(
				http.StatusNotFound,
				fmt.Sprintf("post with id %s not found", id),
				"",
			)
		}

		return guy.NewAppError(
			http.StatusInternalServerError,
			"failed to get post",
			err.Error(),
		)
	}

	if err := uc.repo.DeleteById(ctx, post.ID.String()); err != nil {
		return guy.NewAppError(http.StatusInternalServerError, "failed to delete post", err.Error())
	}

	return nil
}
