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
	GetAllFurnitures() string
	GetFurnitureById(id uint64) (model.Furniture, error)
	CreateFurniture(model.Furniture) (int64, error)
	UpdateFurniture(model.Furniture) (int64, error)
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

func (c *FurnitureRepository) GetFurnitureById(id uint64) (model.Furniture, error) {
	var furniture model.Furniture
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

func (c *FurnitureRepository) CreateFurniture(furniture model.Furniture) (int64, error) {
	// var cId *uint64
	// if furniture.CategoryId != 0 {
	// 	cId = &furniture.CategoryId
	// }
	// var pr *float64
	// if furniture.Price != 0 {
	// 	pr = &furniture.Price
	// }

	// query, args, err := squirrel.Insert(furnitureTableName).
	// 	PlaceholderFormat(squirrel.Dollar).
	// 	Columns(categoryIdColumn, nameColumn, descriptionColumn, priceColumn, imageColumn).
	// 	Values(cId, furniture.Name, furniture.Description, pr, furniture.Image).
	// 	Suffix("RETURNING id").
	// 	ToSql()

	// fmt.Println("==", query, args)

	// if err != nil {
	// 	return 0, err
	// }

	// res, err := c.Db.Exec(query, args...)
	// if err != nil {
	// 	return 0, err
	// }
	// rowAffected, err := res.RowsAffected()
	// if err != nil {
	// 	return 0, err
	// }
	// return rowAffected, nil
	return 1, nil
}

func (c *FurnitureRepository) UpdateFurniture(furnit model.Furniture) (int64, error) {
	// existingFurniture, err := c.GetFurnitureById(furniture.Id)
	// if err != nil {
	// 	return 0, fmt.Errorf("furniture with id = %d not found", furniture.Id)
	// }
	_ = furnit

	furniture := model.Furniture{
		Id:          4,
		CategoryId:  sql.NullInt64{Int64: 0, Valid: false},
		Name:        "New",
		Description: sql.NullString{String: "bar", Valid: false},
		Price:       sql.NullFloat64{Float64: 0.123, Valid: false},
		Image:       sql.NullString{String: "Imagesssss", Valid: false},
		IsActive:    sql.NullBool{Bool: false, Valid: false},
	}

	builder := squirrel.
		Update(furnitureTableName).
		PlaceholderFormat(squirrel.Dollar)
	if furniture.Name != "" {
		builder = builder.Set(nameColumn, furniture.Name)
	}
	// if furniture.Description.String != "" {
	builder = builder.Set(descriptionColumn, furniture.Description)
	// }
	// if furniture.Image.String != "" {
	builder = builder.Set(imageColumn, furniture.Image)
	// }
	// if existingFurniture.IsActive != furniture.IsActive {
	builder = builder.Set(isActiveColumn, furniture.IsActive)
	// }
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

func (c *FurnitureRepository) DeleteFurniture() string {
	return "Delete in repository"
}
