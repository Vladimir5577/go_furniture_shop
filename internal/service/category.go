package service

import (
	"errors"
	"fmt"
	"furniture_shop/internal/model"
	"furniture_shop/internal/repository"
	"os"
)

type ICategoryService interface {
	GetAllCategories(page, pageSize uint64) (model.CategoriesResponse, error)
	GetCategoryById(id uint64) (model.Category, error)
	CreateCategory(model.Category) (int64, error)
	UpdateCategory(model.Category) (int64, error)
	DeleteCategory(uint64) (int64, error)
}

type CategoryService struct {
	Repository repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		Repository: repo,
	}
}

func (c *CategoryService) GetAllCategories(page, pageSize uint64) (model.CategoriesResponse, error) {
	var categories model.CategoriesResponse
	if pageSize == 0 {
		return categories, errors.New("page size should be above 0")
	}
	if page == 0 {
		return categories, errors.New("page number should be above 0")
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

	existingCategory, err := c.Repository.GetCategoryById(category.Id)
	if err != nil {
		return 0, err
	}

	if len(category.Image) > 0 &&
		len(existingCategory.Image) > 0 &&
		category.Image != existingCategory.Image {
		err := os.Remove("./uploads/" + existingCategory.Image)
		if err != nil {
			return 0, err
		}
		fmt.Println("old image deleted")
	}

	return c.Repository.UpdateCategory(category)
}

func (c *CategoryService) DeleteCategory(id uint64) (int64, error) {
	existingCategory, err := c.Repository.GetCategoryById(id)
	if err != nil {
		return 0, err
	}

	if len(existingCategory.Image) > 0 {
		err := os.Remove("./uploads/" + existingCategory.Image)
		if err != nil {
			return 0, err
		}
		fmt.Println("image deleted")
	}
	return c.Repository.DeleteCategory(id)
}
