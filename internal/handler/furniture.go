package handler

import (
	"encoding/json"
	"furniture_shop/internal/model"
	"furniture_shop/internal/service"
	"furniture_shop/internal/utils"
	"net/http"
	"strconv"
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
		fileName, err := utils.UploadImageFromHTTPRequest(r)
		if err != nil {
			w.Write([]byte("Error upload image."))
			w.Write([]byte(err.Error()))
			return
		}

		if len(fileName) > 0 {
			errResizeImage := utils.ResizeImage(fileName)
			if errResizeImage != nil {
				errDeleteImage := utils.DeleteImage(fileName)
				if errDeleteImage != nil {
					w.Write([]byte("Error delete unused image."))
					w.Write([]byte(errDeleteImage.Error()))
					return
				}
				w.Write([]byte("Error resize image."))
				w.Write([]byte(errResizeImage.Error()))
				return
			}
		}

		var price float64
		priceString := r.FormValue("price")
		if len(priceString) > 0 {
			price, err = strconv.ParseFloat(priceString, 64)
			if err != nil {
				w.Write([]byte("Error parse price."))
				w.Write([]byte(err.Error()))
				return
			}
		}

		var categoryId uint64
		categoryIdString := r.FormValue("category_id")
		if len(categoryIdString) > 0 {
			categoryId, err = strconv.ParseUint(categoryIdString, 10, 64)
			if err != nil {
				w.Write([]byte("category_id should be numeric"))
				w.Write([]byte(err.Error()))
				return
			}
		}

		furniture := model.Furniture{
			CategoryId:  categoryId,
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price:       price,
			Image:       fileName,
		}
		res, errDB := c.Service.CreateFurniture(furniture)
		if errDB != nil {
			errDeleteImage := utils.DeleteImage(fileName)
			if errDeleteImage != nil {
				w.Write([]byte("Error insert in database and delete unused image."))
				w.Write([]byte(errDeleteImage.Error()))
				return
			}
			w.Write([]byte("Error inserting in database."))
			w.Write([]byte(errDB.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte("Created successfully."))
		json.NewEncoder(w).Encode(res)
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
