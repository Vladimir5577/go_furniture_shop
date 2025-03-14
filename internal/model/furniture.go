package model

import "database/sql"

type Furniture struct {
	Id          uint64  `json:"id"`
	CategoryId  uint64  `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type FurnitureRequest struct {
	Id          sql.NullInt64   `json:"id"`
	CategoryId  sql.NullInt64   `json:"category_id"`
	Name        sql.NullString  `json:"name"`
	Description sql.NullString  `json:"description"`
	Price       sql.NullFloat64 `json:"price"`
	Image       sql.NullString  `json:"image"`
	IsActive    sql.NullBool    `json:"is_active"`
	CreatedAt   sql.NullString  `json:"created_at"`
	UpdatedAt   sql.NullString  `json:"updated_at"`
}
