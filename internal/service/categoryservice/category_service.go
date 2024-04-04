package categoryservice

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/wiliamvj/api-users-golang/internal/dto"
	"github.com/wiliamvj/api-users-golang/internal/entity"
	"github.com/wiliamvj/api-users-golang/internal/handler/response"
)

func (s *service) CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error {
	categoryEntity := entity.CategoryEntity{
		ID:        uuid.New().String(),
		Title:     u.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.repo.CreateCategory(ctx, &categoryEntity)
	if err != nil {
		return errors.New("error to create category")
	}
	return nil
}

func (s *service) FindManyCategories(ctx context.Context) (*response.ManyCategoriesResponse, error) {
	findManyCategories, err := s.repo.FindManyCategories(ctx)
	if err != nil {
		slog.Error("error to find many categories", "err", err, slog.String("package", "categorieservice"))
		return nil, err
	}
	categories := response.ManyCategoriesResponse{}
	for _, category := range findManyCategories {
		categoryResponse := response.CategoryResponse{
			ID:        category.ID,
			Title:     category.Title,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
		categories.Categories = append(categories.Categories, categoryResponse)
	}
	return &categories, nil
}
