package repository

import (
	"database/sql"
	"fmt"
	"furniture_shop/internal/model"

	"github.com/Masterminds/squirrel"
)

const (
	tableName = "category"

	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	imageColumn       = "image"
	isActiveColumn    = "is_active"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

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
	var (
		category   model.Category
		categories []model.Category
	)
	rows, err := c.Db.Query(
		"select * from category",
	)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.Description, &category.Image, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}
		categories = append(categories, category)
	}
	fmt.Printf("Structure: %+v\n", categories)
	// fmt.Println(categories)
	return "Get all categories from repository"
}

func (c *CategoryRepository) GetCategoryById() string {
	return "Get by id repository"
}

func (c *CategoryRepository) CreateCategory() string {
	query, args, err := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameColumn, descriptionColumn, imageColumn).
		Values("Sofa", "Good sofa", "http://some_image.com").
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		fmt.Println(err)
	}

	res, err := c.Db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.RowsAffected())
	return "Create in repository"
}

func (c *CategoryRepository) UpdateCategory() string {
	return "Update in repository"
}

func (c *CategoryRepository) DeleteCategory() string {
	return "Delete in repository"
}
