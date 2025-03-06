package repository

import "database/sql"

type IFurnitureRepository interface {
	GetAllFurnitures() string
	GetFurnitureById() string
	CreateFurniture() string
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

func (c *FurnitureRepository) CreateFurniture() string {
	return "Create in repository"
}

func (c *FurnitureRepository) UpdateFurniture() string {
	return "Update in repository"
}

func (c *FurnitureRepository) DeleteFurniture() string {
	return "Delete in repository"
}
