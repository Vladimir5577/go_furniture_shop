package service

import (
	"errors"
	"furniture_shop/internal/model"
	"furniture_shop/internal/repository"
)

type IFurnitureService interface {
	GetAllFurnitures() string
	GetFurnitureById() string
	CreateFurniture(model.Furniture) (int64, error)
	UpdateFurniture() string
	DeleteFurniture() string
}

type FurnitureService struct {
	Repository repository.IFurnitureRepository
}

func NewFurnitureService(repo repository.IFurnitureRepository) *FurnitureService {
	return &FurnitureService{
		Repository: repo,
	}
}

func (c *FurnitureService) GetAllFurnitures() string {
	return c.Repository.GetAllFurnitures()
}

func (c *FurnitureService) GetFurnitureById() string {
	return c.Repository.GetFurnitureById()
}

func (c *FurnitureService) CreateFurniture(furniture model.Furniture) (int64, error) {
	if furniture.Name == "" {
		return 0, errors.New("name is required")
	}
	return c.Repository.CreateFurniture(furniture)
}

func (c *FurnitureService) UpdateFurniture() string {
	return c.Repository.UpdateFurniture()
}

func (c *FurnitureService) DeleteFurniture() string {
	return c.Repository.DeleteFurniture()
}
