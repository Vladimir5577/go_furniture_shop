package service

import "furniture_shop/internal/repository"

type ICategoryService interface {
	GetAllCategories() string
	GetCategoryById() string
	CreateCategory() string
	UpdateCategory() string
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

func (c *CategoryService) GetAllCategories() string {
	return c.Repository.GetAllCategories()
}

func (c *CategoryService) GetCategoryById() string {
	return c.Repository.GetCategoryById()
}

func (c *CategoryService) CreateCategory() string {
	return c.Repository.CreateCategory()
}

func (c *CategoryService) UpdateCategory() string {
	return c.Repository.UpdateCategory()
}

func (c *CategoryService) DeleteCategory() string {
	return c.Repository.DeleteCategory()
}
