package handler

import (
	"furniture_shop/internal/service"
	"net/http"
)

type FurnitureHandler struct {
	Service service.IFurnitureService
}

func NewFurnitureHandler(service service.IFurnitureService) *FurnitureHandler {
	return &FurnitureHandler{
		Service: service,
	}
}

func (c *FurnitureHandler) GetAllCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.GetAllFurnitures()))
	}
}

func (c *FurnitureHandler) GetCategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.GetFurnitureById()))
	}
}

func (c *FurnitureHandler) CreateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.CreateFurniture()))
	}
}

func (c *FurnitureHandler) UpdateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.UpdateFurniture()))
	}
}

func (c *FurnitureHandler) DeleteCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.DeleteFurniture()))
	}
}
