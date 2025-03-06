package repository

import "database/sql"

type ICategoryRepository interface {
	GetAllCategories() string
	GetCategoryById() string
	CreateCategory() string
	UpdateCategory() string
	DeleteCategory() string
}

type CategoryRepository struct {
	Db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		Db: db,
	}
}

func (c *CategoryRepository) GetAllCategories() string {
	return "Get all categories from repository"
}

func (c *CategoryRepository) GetCategoryById() string {
	return "Get by id repository"
}

func (c *CategoryRepository) CreateCategory() string {
	return "Create in repository"
}

func (c *CategoryRepository) UpdateCategory() string {
	return "Update in repository"
}

func (c *CategoryRepository) DeleteCategory() string {
	return "Delete in repository"
}
