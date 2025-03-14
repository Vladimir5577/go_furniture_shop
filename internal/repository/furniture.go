package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"furniture_shop/internal/model"
	"time"

	"github.com/Masterminds/squirrel"
)

type IFurnitureRepository interface {
	GetAllFurnitures(page, pageSize uint64) ([]model.Furniture, error)
	GetFurnitureById(id uint64) (model.FurnitureRequest, error)
	CreateFurniture(model.FurnitureRequest) (int64, error)
	UpdateFurniture(model.FurnitureRequest) (int64, error)
	DeleteFurniture(id uint64) (int64, error)
}

type FurnitureRepository struct {
	Db *sql.DB
}

func NewFurnitureRepository(db *sql.DB) *FurnitureRepository {
	return &FurnitureRepository{
		Db: db,
	}
}

func (c *FurnitureRepository) GetAllFurnitures(page, pageSize uint64) ([]model.Furniture, error) {
	var (
		furniture  model.FurnitureRequest
		furnitures []model.Furniture
	)
	builder := squirrel.Select(idColumn, categoryIdColumn, nameColumn, descriptionColumn, priceColumn, imageColumn, isActiveColumn, createdAtColumn, updatedAtColumn).
		From(furnitureTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(isActiveColumn).
		OrderBy(fmt.Sprintf("%s %s", idColumn, orderByAsc))

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
		err = rows.Scan(&furniture.Id, &furniture.CategoryId, &furniture.Name, &furniture.Description, &furniture.Price, &furniture.Image, &furniture.IsActive, &furniture.CreatedAt, &furniture.UpdatedAt)
		if err != nil {
			return nil, err
		}

		f := model.Furniture{
			Id:          uint64(furniture.Id.Int64),
			CategoryId:  uint64(furniture.CategoryId.Int64),
			Name:        furniture.Name.String,
			Description: furniture.Description.String,
			Price:       furniture.Price.Float64,
			Image:       furniture.Image.String,
			IsActive:    furniture.IsActive.Bool,
			CreatedAt:   furniture.CreatedAt.String,
			UpdatedAt:   furniture.UpdatedAt.String,
		}

		furnitures = append(furnitures, f)
	}

	return furnitures, nil
}

func (c *FurnitureRepository) GetFurnitureById(id uint64) (model.FurnitureRequest, error) {
	var furniture model.FurnitureRequest
	query, args, err := squirrel.Select(idColumn, categoryIdColumn, nameColumn, descriptionColumn, priceColumn, imageColumn, isActiveColumn, createdAtColumn, updatedAtColumn).
		From(furnitureTableName).
		PlaceholderFormat(squirrel.Dollar).
		Where((fmt.Sprintf("%s = ?", idColumn)), id).
		Limit(1).
		ToSql()
	if err != nil {
		return furniture, err
	}

	row := c.Db.QueryRow(query, args...)
	err = row.Scan(&furniture.Id, &furniture.CategoryId, &furniture.Name, &furniture.Description, &furniture.Price, &furniture.Image, &furniture.IsActive, &furniture.CreatedAt, &furniture.UpdatedAt)
	if err != nil {
		return furniture, err
	}

	return furniture, nil
}

func (c *FurnitureRepository) CreateFurniture(furniture model.FurnitureRequest) (int64, error) {
	query, args, err := squirrel.Insert(furnitureTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(categoryIdColumn, nameColumn, descriptionColumn, priceColumn, imageColumn, isActiveColumn).
		Values(furniture.CategoryId, furniture.Name, furniture.Description, furniture.Price, furniture.Image, furniture.IsActive).
		Suffix("RETURNING id").
		ToSql()

	fmt.Println("==", query, args)

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

func (c *FurnitureRepository) UpdateFurniture(furniture model.FurnitureRequest) (int64, error) {
	existingFurniture, err := c.GetFurnitureById(uint64(furniture.Id.Int64))
	if err != nil {
		return 0, fmt.Errorf("furniture with id = %d not found", furniture.Id.Int64)
	}

	builder := squirrel.
		Update(furnitureTableName).
		PlaceholderFormat(squirrel.Dollar)
	if furniture.CategoryId.Valid {
		if furniture.CategoryId.Int64 == 0 {
			builder = builder.Set(categoryIdColumn, nil)
		} else {
			builder = builder.Set(categoryIdColumn, furniture.CategoryId.Int64)
		}
	}
	if furniture.Name.Valid {
		builder = builder.Set(nameColumn, furniture.Name)
	}
	if furniture.Description.Valid {
		builder = builder.Set(descriptionColumn, furniture.Description)
	}
	if furniture.Price.Valid {
		builder = builder.Set(priceColumn, furniture.Price)
	}
	if furniture.Image.Valid {
		builder = builder.Set(imageColumn, furniture.Image)
	}
	if furniture.IsActive.Valid {
		if existingFurniture.IsActive.Bool != furniture.IsActive.Bool {
			builder = builder.Set(isActiveColumn, furniture.IsActive)
		}
	}

	builder = builder.Set(updatedAtColumn, time.Now().Format("2006-01-02 15:04:05"))
	query, args, err := builder.
		Where((fmt.Sprintf("%s = ?", idColumn)), furniture.Id).
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

func (c *FurnitureRepository) DeleteFurniture(id uint64) (int64, error) {
	query, args, err := squirrel.
		Delete(furnitureTableName).
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
