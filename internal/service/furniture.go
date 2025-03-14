package service

import (
	"errors"
	"furniture_shop/internal/model"
	"furniture_shop/internal/repository"
)

type IFurnitureService interface {
	GetAllFurnitures() string
	GetFurnitureById(id uint64) (model.Furniture, error)
	CreateFurniture(model.Furniture) (int64, error)
	UpdateFurniture(model.Furniture) (int64, error)
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

func (c *FurnitureService) GetFurnitureById(id uint64) (model.Furniture, error) {
	return c.Repository.GetFurnitureById(id)
}

func (c *FurnitureService) CreateFurniture(furniture model.Furniture) (int64, error) {
	if furniture.Name == "" {
		return 0, errors.New("name is required")
	}
	return c.Repository.CreateFurniture(furniture)
}

func (c *FurnitureService) UpdateFurniture(furniture model.Furniture) (int64, error) {
	if furniture.Name == "" {
		return 0, errors.New("name is required")
	}
	return c.Repository.UpdateFurniture(furniture)
}

func (c *FurnitureService) DeleteFurniture() string {
	return c.Repository.DeleteFurniture()
}
