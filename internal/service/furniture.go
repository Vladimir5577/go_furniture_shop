package service

import (
	"database/sql"
	"errors"
	"fmt"
	"furniture_shop/internal/model"
	"furniture_shop/internal/repository"
	"furniture_shop/internal/utils"
)

type IFurnitureService interface {
	GetAllFurnitures(model.FurnitureQueryparams) (model.FurnituresResponse, error)
	GetFurnitureById(id uint64) (model.Furniture, error)
	CreateFurniture(model.FurnitureRequest) (int64, error)
	UpdateFurniture(model.FurnitureRequest) (int64, error)
	DeleteFurniture(id uint64) (int64, error)
}

type FurnitureService struct {
	Repository repository.IFurnitureRepository
}

func NewFurnitureService(repo repository.IFurnitureRepository) *FurnitureService {
	return &FurnitureService{
		Repository: repo,
	}
}

func (c *FurnitureService) GetAllFurnitures(queryParams model.FurnitureQueryparams) (model.FurnituresResponse, error) {
	var furnitures model.FurnituresResponse
	if queryParams.PageSize == 0 {
		return furnitures, errors.New("page size should be above 0")
	}
	if queryParams.Page == 0 {
		return furnitures, errors.New("page number should be above 0")
	}
	return c.Repository.GetAllFurnitures(queryParams)
}

func (c *FurnitureService) GetFurnitureById(id uint64) (model.Furniture, error) {
	var furniture model.Furniture
	furnitureRequest, err := c.Repository.GetFurnitureById(id)
	if err != nil {
		_ = furnitureRequest
		return furniture, err
	}

	furniture = model.Furniture{
		Id:          uint64(furnitureRequest.CategoryId.Int64),
		CategoryId:  uint64(furnitureRequest.CategoryId.Int64),
		Name:        furnitureRequest.Name.String,
		Description: furnitureRequest.Description.String,
		Price:       furnitureRequest.Price.Float64,
		Image:       furnitureRequest.Image.String,
		IsActive:    furnitureRequest.IsActive.Bool,
		CreatedAt:   furnitureRequest.CreatedAt.String,
		UpdatedAt:   furnitureRequest.UpdatedAt.String,
	}

	return furniture, nil
}

func (c *FurnitureService) CreateFurniture(furniture model.FurnitureRequest) (int64, error) {
	if !furniture.Name.Valid {
		return 0, errors.New("name is required")
	}
	if !furniture.IsActive.Valid {
		furniture.IsActive = sql.NullBool{
			Bool:  true,
			Valid: true,
		}
	}
	return c.Repository.CreateFurniture(furniture)
}

func (c *FurnitureService) UpdateFurniture(furniture model.FurnitureRequest) (int64, error) {
	if !furniture.Id.Valid {
		return 0, errors.New("id is required")
	}
	if furniture.Name.Valid && furniture.Name.String == "" {
		return 0, errors.New("name is required")
	}
	return c.Repository.UpdateFurniture(furniture)
}

func (c *FurnitureService) DeleteFurniture(id uint64) (int64, error) {
	existingFurniture, err := c.GetFurnitureById(id)
	if err != nil {
		return 0, err
	}

	rowNum, err := c.Repository.DeleteFurniture(id)
	if err != nil {
		return 0, err
	}

	err = utils.DeleteImage(existingFurniture.Image)
	if err != nil {
		return 0, fmt.Errorf("record deleted but old image not deleted")
	}
	return rowNum, nil
}
