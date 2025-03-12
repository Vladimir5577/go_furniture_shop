package repository

import (
	"database/sql"
	"fmt"
	"furniture_shop/internal/model"

	"github.com/Masterminds/squirrel"
)

type IFurnitureRepository interface {
	GetAllFurnitures() string
	GetFurnitureById() string
	CreateFurniture(model.Furniture) (int64, error)
	UpdateFurniture() string
	DeleteFurniture() string
}

type FurnitureRepository struct {
	Db *sql.DB
}

func NewFurnitureRepository(db *sql.DB) *FurnitureRepository {
	return &FurnitureRepository{
		Db: db,
	}
}

func (c *FurnitureRepository) GetAllFurnitures() string {
	return "Get all furniture from repository"
}

func (c *FurnitureRepository) GetFurnitureById() string {
	return "Get by id repository"
}

func (c *FurnitureRepository) CreateFurniture(furniture model.Furniture) (int64, error) {
	var cId *uint64
	if furniture.CategoryId != 0 {
		cId = &furniture.CategoryId
	}
	var pr *float64
	if furniture.Price != 0 {
		pr = &furniture.Price
	}

	query, args, err := squirrel.Insert(furnitureTableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(categoryIdColumn, nameColumn, descriptionColumn, priceColumn, imageColumn).
		Values(cId, furniture.Name, furniture.Description, pr, furniture.Image).
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

func (c *FurnitureRepository) UpdateFurniture() string {
	return "Update in repository"
}

func (c *FurnitureRepository) DeleteFurniture() string {
	return "Delete in repository"
}
