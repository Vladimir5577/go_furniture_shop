package model

import "database/sql"

type Furniture struct {
	Id          uint64          `json:"id"`
	CategoryId  sql.NullInt64   `json:"category_id"`
	Name        string          `json:"name"`
	Description sql.NullString  `json:"description"`
	Price       sql.NullFloat64 `json:"price"`
	Image       sql.NullString  `json:"image"`
	IsActive    sql.NullBool    `json:"is_active"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

type FurnitureSQL struct {
	Id          uint64          `json:"id"`
	CategoryId  sql.NullInt64   `json:"category_id"`
	Name        string          `json:"name"`
	Description sql.NullString  `json:"description"`
	Price       sql.NullFloat64 `json:"price"`
	Image       sql.NullString  `json:"image"`
	IsActive    sql.NullBool    `json:"is_active"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}
