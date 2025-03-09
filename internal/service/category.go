package service

import (
	"errors"
	"furniture_shop/internal/model"
	"furniture_shop/internal/repository"
)

type ICategoryService interface {
	GetAllCategories(page, pageSize uint64) ([]model.Category, error)
	GetCategoryById(id uint64) (model.Category, error)
	CreateCategory(model.Category) (int64, error)
	UpdateCategory(model.Category) (int64, error)
	DeleteCategory() string
}

type CategoryService struct {
	Repository repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		Repository: repo,
	}
}

func (c *CategoryService) GetAllCategories(page, pageSize uint64) ([]model.Category, error) {
	if pageSize == 0 {
		return nil, errors.New("page size should be above 0")
	}
	if page == 0 {
		return nil, errors.New("page number should be above 0")
	}
	return c.Repository.GetAllCategories(page, pageSize)
}

func (c *CategoryService) GetCategoryById(id uint64) (model.Category, error) {
	return c.Repository.GetCategoryById(id)
}

func (c *CategoryService) CreateCategory(category model.Category) (int64, error) {
	if category.Name == "" {
		return 0, errors.New("name is required")
	}
	return c.Repository.CreateCategory(category)
}

func (c *CategoryService) UpdateCategory(category model.Category) (int64, error) {
	if category.Id == 0 {
		return 0, errors.New("id required")
	}
	if category.Name == "" {
		return 0, errors.New("name required")
	}

	return c.Repository.UpdateCategory(category)
}

func (c *CategoryService) DeleteCategory() string {
	return c.Repository.DeleteCategory()
}
