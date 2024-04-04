package categoryservice

import (
	"context"

	"github.com/wiliamvj/api-users-golang/internal/dto"
	"github.com/wiliamvj/api-users-golang/internal/handler/response"
	"github.com/wiliamvj/api-users-golang/internal/repository/categoryrepository"
)

func NewCategoryService(repo categoryrepository.CategoryRepository) CategoryService {
	return &service{
		repo,
	}
}

type service struct {
	repo categoryrepository.CategoryRepository
}

type CategoryService interface {
	CreateCategory(ctx context.Context, u dto.CreateCategoryDto) error
	FindManyCategories(ctx context.Context) (*response.ManyCategoriesResponse, error)
}
