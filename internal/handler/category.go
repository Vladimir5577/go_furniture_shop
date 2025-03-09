package handler

import (
	"encoding/json"
	"fmt"
	"furniture_shop/internal/model"
	"furniture_shop/internal/service"
	"net/http"
	"strconv"
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
		var page, pageSize uint64
		// default pagination
		page = 1
		pageSize = 20
		var err error
		pageString := r.URL.Query().Get("page")
		pageSizeString := r.URL.Query().Get("pageSize")

		if pageString != "" {
			page, err = strconv.ParseUint(pageString, 10, 64)
			if err != nil {
				w.Write([]byte("Page number should be positive numeric"))
				return
			}
			// if page number sended than make sure page size also sended
			if pageSizeString == "" {
				w.Write([]byte("Page size required"))
				return
			}
		}

		if pageSizeString != "" {
			pageSize, err = strconv.ParseUint(pageSizeString, 10, 64)
			if err != nil {
				w.Write([]byte("PageSize should be positive numeric"))
				return
			}
		}

		fmt.Println("Page and page size =>", page, pageSize)
		res, err := c.Service.GetAllCategories(page, pageSize)
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}

func (c *CategoryHandler) GetCategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		if idString == "" {
			w.Write([]byte("id required"))
			return
		}
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("id should be positive numeric"))
			return
		}

		// w.Write([]byte(id))
		res, err := c.Service.GetCategoryById(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}

func (c *CategoryHandler) CreateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category model.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			w.Write([]byte("Error parsing json."))
			w.Write([]byte(err.Error()))
			return
		}

		res, err := c.Service.CreateCategory(category)
		if err != nil {
			w.Write([]byte("Error inserting in database."))
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("Created successfully."))
		json.NewEncoder(w).Encode(res)
	}
}

func (c *CategoryHandler) UpdateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var category model.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			w.Write([]byte("Error parsing json."))
			w.Write([]byte(err.Error()))
			return
		}

		_, err = c.Service.UpdateCategory(category)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte("Update."))
	}
}

func (c *CategoryHandler) DeleteCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.DeleteCategory()))
	}
}
