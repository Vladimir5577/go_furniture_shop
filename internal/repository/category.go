package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"furniture_shop/internal/model"
	"time"

	"github.com/Masterminds/squirrel"
)

type ICategoryRepository interface {
	GetAllCategories(page, pageSize uint64) (model.CategoriesResponse, error)
	GetCategoryById(id uint64) (model.Category, error)
	CreateCategory(model.Category) (int64, error)
	UpdateCategory(model.Category) (int64, error)
	DeleteCategory(id uint64) (int64, error)
}

type CategoryRepository struct {
	Db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		Db: db,
	}
}

func (c *CategoryRepository) CountRows(isAdmin bool) (uint64, error) {
	var count uint64
	sqlQuery := "SELECT COUNT(id) FROM category "
	if !isAdmin {
		sqlQuery += "WHERE is_active = true"
	}
	err := c.Db.QueryRow(sqlQuery).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (c *CategoryRepository) GetAllCategories(page, pageSize uint64) (model.CategoriesResponse, error) {
	var (
		category           model.Category
		categories         []model.Category
		categoriesResponse model.CategoriesResponse
	)
	builder := squirrel.Select(idColumn, nameColumn, descriptionColumn, imageColumn, isActiveColumn, createdAtColumn, updatedAtColumn).
		From(categoryTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(isActiveColumn)

	if page != 0 && pageSize != 0 {
		builder = builder.
			Limit(pageSize).
			Offset(pageSize * (page - 1))
	}
	query, args, err := builder.ToSql()

	if err != nil {
		return categoriesResponse, err
	}

	rows, err := c.Db.Query(query, args...)
	if err != nil {
		return categoriesResponse, err
	}

	for rows.Next() {
		err = rows.Scan(&category.Id, &category.Name, &category.Description, &category.Image, &category.IsActive, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return categoriesResponse, err
		}
		categories = append(categories, category)
	}

	numRows, err := c.CountRows(true)
	if err != nil {
		return categoriesResponse, err
	}

	pageCount := numRows / pageSize
	if numRows%pageSize > 0 {
		pageCount++
	}

	pages := make([]uint64, pageCount)
	var i uint64
	for i = 0; i < pageCount; i++ {
		pages[i] = i + 1
	}

	categoriesResponse = model.CategoriesResponse{
		Categories: categories,
		Count:      numRows,
		Page:       page,
		PageSize:   pageSize,
		PageCount:  pageCount,
		Pages:      pages,
	}

	return categoriesResponse, nil
}

func (c *CategoryRepository) GetCategoryById(id uint64) (model.Category, error) {
	var category model.Category
	query, args, err := squirrel.Select(idColumn, nameColumn, descriptionColumn, imageColumn, isActiveColumn, createdAtColumn, updatedAtColumn).
		From(categoryTableName).
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
	// TO DO check for existing name manually
	query, args, err := squirrel.Insert(categoryTableName).
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
	existingCategory, err := c.GetCategoryById(category.Id)
	if err != nil {
		return 0, fmt.Errorf("category with id = %d not found", category.Id)
	}

	builder := squirrel.
		Update(categoryTableName).
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
	builder = builder.Set(updatedAtColumn, time.Now().Format("2006-01-02 15:04:05"))
	query, args, err := builder.
		Where((fmt.Sprintf("%s = ?", idColumn)), category.Id).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, errors.New(err.Error())
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

func (c *CategoryRepository) DeleteCategory(id uint64) (int64, error) {
	query, args, err := squirrel.
		Delete(categoryTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", idColumn)), id).
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
	if rowAffected == 0 {
		return rowAffected, fmt.Errorf("record not deleted. Maybe unexisting id")
	}
	return rowAffected, nil
}
