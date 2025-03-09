package repository

import (
	"database/sql"
	"errors"
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
	GetAllCategories(page, pageSize uint64) ([]model.Category, error)
	GetCategoryById(id uint64) (model.Category, error)
	CreateCategory(model.Category) (int64, error)
	UpdateCategory(model.Category) (int64, error)
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

func (c *CategoryRepository) GetAllCategories(page, pageSize uint64) ([]model.Category, error) {
	var (
		category   model.Category
		categories []model.Category
	)
	builder := squirrel.Select(idColumn, nameColumn, descriptionColumn, imageColumn, isActiveColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(isActiveColumn)

	if page != 0 && pageSize != 0 {
		builder = builder.
			Limit(pageSize).
			Offset(pageSize * (page - 1))
	}
	query, args, err := builder.ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := c.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.Description, &category.Image, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c *CategoryRepository) GetCategoryById(id uint64) (model.Category, error) {
	var category model.Category
	query, args, err := squirrel.Select(idColumn, nameColumn, descriptionColumn, imageColumn, isActiveColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", idColumn)), id).
		Limit(1).
		ToSql()
	if err != nil {
		return category, err
	}

	row := c.Db.QueryRow(query, args...)
	err = row.Scan(&category.Id, &category.Name, &category.Description, &category.Image, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (c *CategoryRepository) CreateCategory(category model.Category) (int64, error) {
	query, args, err := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameColumn, descriptionColumn, imageColumn).
		Values(category.Name, category.Description, category.Image).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	res, err := c.Db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

func (c *CategoryRepository) UpdateCategory(category model.Category) (int64, error) {
	existingCategory, err := c.GetCategoryById(uint64(category.Id))
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Category with id = %d not found.", category.Id))
	}

	builder := squirrel.
		Update(tableName).
		PlaceholderFormat(squirrel.Dollar)
	if category.Name != "" {
		builder = builder.Set(nameColumn, category.Name)
	}
	if category.Description != "" {
		builder = builder.Set(descriptionColumn, category.Description)
	}
	if category.Image != "" {
		builder = builder.Set(imageColumn, category.Image)
	}
	if existingCategory.IsActive != category.IsActive {
		builder = builder.Set(isActiveColumn, category.IsActive)
	}
	query, args, err := builder.
		Where((fmt.Sprintf("%s = ?", idColumn)), category.Id).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, errors.New(err.Error())
	}

	fmt.Println(query, args)

	res, err := c.Db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowAffected, nil
}

func (c *CategoryRepository) DeleteCategory() string {
	return "Delete in repository"
}
