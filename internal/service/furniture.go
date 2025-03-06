package service

import "furniture_shop/internal/repository"

type IFurnitureService interface {
	GetAllFurnitures() string
	GetFurnitureById() string
	CreateFurniture() string
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

func (c *FurnitureService) CreateFurniture() string {
	return c.Repository.CreateFurniture()
}

func (c *FurnitureService) UpdateFurniture() string {
	return c.Repository.UpdateFurniture()
}

func (c *FurnitureService) DeleteFurniture() string {
	return c.Repository.DeleteFurniture()
}
