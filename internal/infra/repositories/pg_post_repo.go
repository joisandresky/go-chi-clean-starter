package repositories

import (
	"context"

	"gorm.io/gorm"

	"github.com/joisandresky/go-chi-clean-starter/internal/domain/entities"
	repo "github.com/joisandresky/go-chi-clean-starter/internal/domain/repositories"
)

type pgPostRepo struct {
	db *gorm.DB
}

func NewPgPostRepository(db *gorm.DB) repo.PostRepository {
	return &pgPostRepo{db: db}
}

func (pg *pgPostRepo) GetAll(ctx context.Context) ([]entities.Post, error) {
	var posts []entities.Post
	if err := pg.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pg *pgPostRepo) GetById(ctx context.Context, id string) (*entities.Post, error) {
	var post entities.Post
	if err := pg.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

func (pg *pgPostRepo) Create(ctx context.Context, post *entities.Post) error {
	if err := pg.db.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

func (pg *pgPostRepo) UpdateById(ctx context.Context, post *entities.Post) error {
	if err := pg.db.Save(&post).Error; err != nil {
		return err
	}

	return nil
}

func (pg *pgPostRepo) DeleteById(ctx context.Context, id string) error {
	if err := pg.db.Where("id = ?", id).Delete(&entities.Post{}).Error; err != nil {
		return err
	}

	return nil
}
