package model

import "database/sql"

type Category struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CategoryRequest struct {
	Id          uint64         `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Image       sql.NullString `json:"image"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

type CategoryResponse struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type CategoriesResponse struct {
	Categories  []Category `json:"categories"`
	Count       uint64     `json:"count"`
	Page        uint64     `json:"page"`
	PageSize    uint64     `json:"page_size"`
	PageCount   uint64     `json:"page_count"`
	Pages       []uint64
	CurrentPage uint64
}
