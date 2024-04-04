package categoryrepository

import (
	"context"

	"github.com/wiliamvj/api-users-golang/internal/database/sqlc"
	"github.com/wiliamvj/api-users-golang/internal/entity"
)

func (r *repository) CreateCategory(ctx context.Context, c *entity.CategoryEntity) error {
	err := r.queries.CreateCategory(ctx, sqlc.CreateCategoryParams{
		ID:        c.ID,
		Title:     c.Title,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindManyCategories(ctx context.Context) ([]entity.CategoryEntity, error) {
	categories, err := r.queries.FindManyCategories(ctx)
	if err != nil {
		return nil, err
	}
	var categoriesEntity []entity.CategoryEntity
	for _, category := range categories {
		categoryEntity := entity.CategoryEntity{
			ID:        category.ID,
			Title:     category.Title,
			UpdatedAt: category.UpdatedAt,
		}
		categoriesEntity = append(categoriesEntity, categoryEntity)
	}
	return categoriesEntity, nil
}
