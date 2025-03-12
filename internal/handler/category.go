package handler

import (
	"encoding/json"
	"fmt"
	"furniture_shop/internal/model"
	"furniture_shop/internal/service"
	"furniture_shop/internal/utils"
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

		category := model.Category{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Image:       fileName,
		}
		res, errDB := c.Service.CreateCategory(category)
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

func (c *CategoryHandler) UpdateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TO DO update iumage available but not deleet
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

		existingCategory, err := c.Service.GetCategoryById(id)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		var isActive bool
		if r.FormValue("is_active") != "" {
			isActive, err = strconv.ParseBool(r.FormValue("is_active"))
			if err != nil {
				w.Write([]byte("is_active incorrect value"))
				w.Write([]byte(err.Error()))
				return
			}
		} else {
			isActive = existingCategory.IsActive
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
			if existingCategory.Image != fileHeader.Filename {
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
				fileName = existingCategory.Image
			}
		}

		category := model.Category{
			Id:          id,
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Image:       fileName,
			IsActive:    isActive,
		}
		res, errDb := c.Service.UpdateCategory(category)
		if errDb != nil {
			w.Write([]byte("Error inserting in database."))
			w.Write([]byte(errDb.Error()))
			return
		}
		fmt.Println(res)

		w.Write([]byte("Updated successfully."))
	}
}

func (c *CategoryHandler) DeleteCategory() http.HandlerFunc {
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

		res, err := c.Service.DeleteCategory(id)
		if err != nil {
			fmt.Println("Error occured", err)
			return
		}

		fmt.Println(res)

		w.Write([]byte("Deleted successfully"))
	}
}
