package handler

import (
	"furniture_shop/internal/service"
	"net/http"
)

type CategoryHandler struct {
	Service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{
		Service: service,
	}
}

func (c *CategoryHandler) GetAllCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.GetAllCategories()))
	}
}

func (c *CategoryHandler) GetCategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.GetCategoryById()))
	}
}

func (c *CategoryHandler) CreateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.CreateCategory()))
	}
}

func (c *CategoryHandler) UpdateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.UpdateCategory()))
	}
}

func (c *CategoryHandler) DeleteCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.DeleteCategory()))
	}
}
