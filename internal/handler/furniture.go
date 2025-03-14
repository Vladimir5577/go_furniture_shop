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

		res, err := c.Service.GetAllFurnitures(page, pageSize)
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(res)
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

		var categoryId sql.NullInt64
		categoryIdString := r.FormValue("category_id")
		if len(categoryIdString) > 0 {
			cId, err := strconv.ParseInt(categoryIdString, 10, 64)
			if err != nil {
				w.Write([]byte("Category id should be numeric"))
				w.Write([]byte(err.Error()))
				return
			}
			if cId > 0 {
				categoryId = sql.NullInt64{
					Int64: cId,
					Valid: true,
				}
			}
		}

		var name sql.NullString
		nameString := r.FormValue("name")
		if len(nameString) > 0 {
			name = sql.NullString{
				String: nameString,
				Valid:  true,
			}
		}

		var description sql.NullString
		descriptionString := r.FormValue("description")
		if len(descriptionString) > 0 {
			description = sql.NullString{
				String: descriptionString,
				Valid:  true,
			}
		}

		var price sql.NullFloat64
		priceString := r.FormValue("price")
		if len(priceString) > 0 {
			pr, err := strconv.ParseFloat(priceString, 64)
			if err != nil {
				w.Write([]byte("price should be numeric."))
				w.Write([]byte(err.Error()))
				return
			}
			if pr > 0 {
				price = sql.NullFloat64{
					Float64: pr,
					Valid:   true,
				}
			}
		}

		var image sql.NullString
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
			image = sql.NullString{
				String: fileName,
				Valid:  true,
			}
		}

		var isActive sql.NullBool
		isActiveString := r.FormValue("is_active")
		if len(isActiveString) > 0 {
			ac, err := strconv.ParseBool(isActiveString)
			if err != nil {
				w.Write([]byte("is_active incorrect value should be boolean"))
				w.Write([]byte(err.Error()))
				return
			}
			isActive = sql.NullBool{
				Bool:  ac,
				Valid: true,
			}
		}

		furniture := model.FurnitureRequest{
			CategoryId:  categoryId,
			Name:        name,
			Description: description,
			Price:       price,
			Image:       image,
			IsActive:    isActive,
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
		var id sql.NullInt64
		idString := r.FormValue("id")
		if len(idString) > 0 {
			idF, err := strconv.ParseInt(idString, 10, 64)
			if err != nil {
				w.Write([]byte("id should be numeric"))
				w.Write([]byte(err.Error()))
				return
			}
			if idF > 0 {
				id = sql.NullInt64{
					Int64: idF,
					Valid: true,
				}
			}
		}

		var categoryId sql.NullInt64
		categoryIdString := r.FormValue("category_id")
		if len(categoryIdString) > 0 {
			cId, err := strconv.ParseInt(categoryIdString, 10, 64)
			if err != nil {
				w.Write([]byte("Category id should be numeric"))
				w.Write([]byte(err.Error()))
				return
			}
			if cId >= 0 {
				categoryId = sql.NullInt64{
					Int64: cId,
					Valid: true,
				}
			}
		}

		var name sql.NullString
		nameString := r.FormValue("name")
		if len(nameString) > 0 {
			name = sql.NullString{
				String: nameString,
				Valid:  true,
			}
		}

		var description sql.NullString
		descriptionString := r.FormValue("description")
		if len(descriptionString) > 0 || descriptionString == "" {
			description = sql.NullString{
				String: descriptionString,
				Valid:  true,
			}
		}

		var price sql.NullFloat64
		priceString := r.FormValue("price")
		if len(priceString) > 0 || priceString == "0" {
			pr, err := strconv.ParseFloat(priceString, 64)
			if err != nil {
				w.Write([]byte("price should be numeric."))
				w.Write([]byte(err.Error()))
				return
			}
			if pr > 0 {
				price = sql.NullFloat64{
					Float64: pr,
					Valid:   true,
				}
			}
		}

		var isActive sql.NullBool
		isActiveString := r.FormValue("is_active")
		if len(isActiveString) > 0 {
			ac, err := strconv.ParseBool(isActiveString)
			if err != nil {
				w.Write([]byte("is_active incorrect value should be boolean"))
				w.Write([]byte(err.Error()))
				return
			}
			isActive = sql.NullBool{
				Bool:  ac,
				Valid: true,
			}
		}

		existingFurniture, err := c.Service.GetFurnitureById(uint64(id.Int64))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		var image sql.NullString
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			// image not sended
		}
		if file != nil && len(fileHeader.Filename) > 0 {
			if existingFurniture.Image != fileHeader.Filename {
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
					image = sql.NullString{
						String: fileName,
						Valid:  true,
					}
					errDeleteImage := utils.DeleteImage(existingFurniture.Image)
					if errDeleteImage != nil {
						w.Write([]byte("Error delete unused image."))
						w.Write([]byte(errDeleteImage.Error()))
						return
					}
				}
			} else {
				image = sql.NullString{
					String: "",
					Valid:  false,
				}
			}
		}

		fmt.Println("---", image)
		furniture := model.FurnitureRequest{
			Id:          id,
			CategoryId:  categoryId,
			Name:        name,
			Description: description,
			Price:       price,
			Image:       image,
			IsActive:    isActive,
		}

		res, errDb := c.Service.UpdateFurniture(furniture)
		if errDb != nil {
			w.Write([]byte("Error updating in database."))
			w.Write([]byte(errDb.Error()))
			return
		}

		fmt.Println(res)

		w.Write([]byte("Updated successfully."))
	}
}

func (c *FurnitureHandler) DeleteFurniture() http.HandlerFunc {
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

		res, err := c.Service.DeleteFurniture(id)
		if err != nil {
			w.Write([]byte("Error delete furniture."))
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte("Record furniture deleted."))
		w.Write([]byte(fmt.Sprintf("Rows affected: %d", res)))
	}
}
