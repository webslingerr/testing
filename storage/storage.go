package storage

import (
	"app/models"

	"golang.org/x/net/context"
)

type StorageI interface {
	CloseDB()
	Category() CategoryRepoI
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetById(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest)  (*models.GetListCategoryResponse, error)
	Update(context.Context, *models.UpdateCategory) (int64, error)
	Delete(context.Context, *models.CategoryPrimaryKey) (int64, error)
}
