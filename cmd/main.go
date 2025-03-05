package main

import (
	"encoding/json"
	"fmt"
	"furniture_shop/internal/handler"
	"furniture_shop/internal/service"
	"net/http"
)

func main() {
	handlMux := http.NewServeMux()

	categoryService := service.NewCategoryService()
	categoryHandler := handler.NewCategoryHandler(categoryService)

	handlMux.HandleFunc("GET /categories", categoryHandler.GetAllCategories())
	handlMux.HandleFunc("GET /category", categoryHandler.GetCategoryById())
	handlMux.HandleFunc("POST /category", categoryHandler.CreateCategory())
	handlMux.HandleFunc("PUT /category", categoryHandler.UpdateCategory())
	handlMux.HandleFunc("DELETE /category", categoryHandler.DeleteCategory())

	handlMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Furniture web site.")
		// w.Write([]byte("Hello"))
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: handlMux,
	}

	fmt.Println("Server up and running on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		panic("Error running server !!!")
	}
}
