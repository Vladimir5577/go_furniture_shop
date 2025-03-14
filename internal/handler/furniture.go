package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func (c *FurnitureHandler) GetAllFurnitures() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.GetAllFurnitures()))
	}
}

func (c *FurnitureHandler) GetFurnitureById() http.HandlerFunc {
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
		res, err := c.Service.GetFurnitureById(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
	}
}

func (c *FurnitureHandler) CreateFurniture() http.HandlerFunc {
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

		// var categoryId uint64
		// categoryIdString := r.FormValue("category_id")
		// if len(categoryIdString) > 0 {
		// 	categoryId, err = strconv.ParseUint(categoryIdString, 10, 64)
		// 	if err != nil {
		// 		w.Write([]byte("category_id should be numeric"))
		// 		w.Write([]byte(err.Error()))
		// 		return
		// 	}
		// }

		furniture := model.Furniture{
			CategoryId:  sql.NullInt64{Int64: 123},
			Name:        r.FormValue("name"),
			Description: sql.NullString{String: r.FormValue("description")},
			Price:       sql.NullFloat64{Float64: price},
			Image:       sql.NullString{String: fileName},
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

func (c *FurnitureHandler) UpdateFurniture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TO DO update iumage available but not delete
		idString := r.FormValue("id")
		if idString == "" {
			w.Write([]byte("id required"))
			return
		}
		id, err := strconv.ParseUint(r.FormValue("id"), 10, 64)
		if err != nil {
			w.Write([]byte("id should be numeric"))
			w.Write([]byte(err.Error()))
			return
		}

		// var categoryId uint64
		// categoryIdString := r.FormValue("category_id")
		// if len(categoryIdString) > 0 {
		// 	categoryId, err = strconv.ParseUint(categoryIdString, 10, 64)
		// 	if err != nil {
		// 		w.Write([]byte("category_id should be numeric"))
		// 		w.Write([]byte(err.Error()))
		// 		return
		// 	}
		// }

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

		existingFurniture, err := c.Service.GetFurnitureById(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		var isActive bool
		if r.FormValue("is_active") != "" {
			isActive, err = strconv.ParseBool(r.FormValue("is_active"))
			if err != nil {
				w.Write([]byte("is_active incorrect value should be boolean"))
				w.Write([]byte(err.Error()))
				return
			}
		}

		var fileName string
		r.ParseMultipartForm(10 << 20) //10 MB
		file, fileHeader, err := r.FormFile("image")
		// if file not sended it is an error
		if err != nil {
			// if file not sended - then delete existed, why ?
			// errDeleteImage := utils.DeleteImage(fileName)
			// if errDeleteImage != nil {
			// 	w.Write([]byte("Error delete unused image."))
			// 	w.Write([]byte(errDeleteImage.Error()))
			// 	return
			// }
			// w.Write([]byte("Error form data to retrive image."))
			// w.Write([]byte(err.Error()))
			// return
		}

		if file != nil {
			if existingFurniture.Image.String != fileHeader.Filename {
				fmt.Println("exist image != upload image")
				fileName, err = utils.UploadImageFromHTTPRequest(r)
				if err != nil {
					w.Write([]byte("Error upload image."))
					w.Write([]byte(err.Error()))
					return
				}

				if len(fileName) > 0 {
					errResizeImage := utils.ResizeImage(fileName)
					if errResizeImage != nil {
						w.Write([]byte("Error resize image."))
						w.Write([]byte(errResizeImage.Error()))
						return
					}
				}
			} else {
				fmt.Println("exist image == upload image")
				fileName = existingFurniture.Image.String
			}
		}

		furniture := model.Furniture{
			Id:          id,
			CategoryId:  sql.NullInt64{},
			Name:        r.FormValue("name"),
			Description: sql.NullString{String: r.FormValue("description")},
			Price:       sql.NullFloat64{Float64: price},
			Image:       sql.NullString{String: fileName},
			IsActive:    sql.NullBool{Bool: isActive},
		}
		res, errDb := c.Service.UpdateFurniture(furniture)
		if errDb != nil {
			w.Write([]byte("Error inserting in database."))
			w.Write([]byte(errDb.Error()))
			return
		}
		fmt.Println(res)

		w.Write([]byte("Updated successfully."))
	}
}

func (c *FurnitureHandler) DeleteFurniture() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(c.Service.DeleteFurniture()))
	}
}
