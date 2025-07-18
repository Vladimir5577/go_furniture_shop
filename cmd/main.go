package main

import (
	"encoding/json"
	"fmt"
	"furniture_shop/internal/config"
	"furniture_shop/internal/handler"
	"furniture_shop/internal/repository"
	"furniture_shop/internal/service"
	"log"
	"net/http"
)

func main() {

	envConfigs, err := config.NewEnvConfig()
	if err != nil {
		log.Fatalf("Can't read .env file, %v\n", err)
	}
	db, err := config.ConnectDB(envConfigs)
	if err != nil {
		log.Fatalf("Can't connect to database, %v\n", err)
	}
	_ = db

	// os.Exit(123)

	handlMux := http.NewServeMux()
	// fs := http.FileServer(http.Dir("./uploads"))
	// handlMux.Handle("./uploads/*", http.StripPrefix("./uploads/", fs))

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// TO DO categories for users
	// handlMux.HandleFunc("GET /categories", categoryHandler.GetAllCategories())
	handlMux.HandleFunc("GET /category/{id}", categoryHandler.GetCategoryById())
	handlMux.HandleFunc("POST /category", categoryHandler.CreateCategory())
	handlMux.HandleFunc("PUT /category", categoryHandler.UpdateCategory())
	handlMux.HandleFunc("DELETE /category/{id}", categoryHandler.DeleteCategory())

	furnitureRepository := repository.NewFurnitureRepository(db)
	furnitureService := service.NewFurnitureService(furnitureRepository)
	furnitureHandler := handler.NewFurnitureHandler(furnitureService)

	handlMux.HandleFunc("GET /furnitures", furnitureHandler.GetAllFurnitures())
	handlMux.HandleFunc("GET /furniture/{id}", furnitureHandler.GetFurnitureById())
	handlMux.HandleFunc("POST /furniture", furnitureHandler.CreateFurniture())
	handlMux.HandleFunc("PUT /furniture", furnitureHandler.UpdateFurniture())
	handlMux.HandleFunc("DELETE /furniture/{id}", furnitureHandler.DeleteFurniture())

	// admin
	handlMux.HandleFunc("GET /admin/furnitures", furnitureHandler.AdminGetAllFurnitures())
	handlMux.HandleFunc("GET /admin/furniture/{id}", furnitureHandler.AdminGetFurnitureById())
	handlMux.HandleFunc("GET /admin/categories", categoryHandler.AdminGetAllCategories())

	handlMux.HandleFunc("GET /index", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode("Furniture web site.")
		// w.Write([]byte("Hello"))
	})

	handlMux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	server := http.Server{
		Addr:    ":8081",
		Handler: handlMux,
	}

	fmt.Println("Server up and running on port 8081")
	err = server.ListenAndServe()
	if err != nil {
		panic("Error running server !!!")
	}
}
