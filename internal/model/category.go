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
