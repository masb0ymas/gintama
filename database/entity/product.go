package entity

import "time"

type Product struct {
	Base
	Name         string  `json:"name" gorm:"type:varchar; not null" validate:"required"`
	Image        string  `json:"image" gorm:"type:varchar; not null" validate:"required"`
	Category     string  `json:"category" gorm:"type:varchar; not null" validate:"required"`
	Description  string  `json:"description" gorm:"type:varchar; not null" validate:"required"`
	Rating       int64   `json:"rating" gorm:"type:int; not null" default:"0"`
	NumReviews   int64   `json:"num_reviews" gorm:"type:int; not null" default:"0"`
	Price        float64 `json:"price" gorm:"type:decimal(10,2); not null" default:"0"`
	CountInStock int64   `json:"count_in_stock" gorm:"type:int; not null" default:"0"`
}

type ProductReq struct {
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
	Rating       int64   `json:"rating"`
	NumReviews   int64   `json:"num_reviews"`
	Price        float64 `json:"price"`
	CountInStock int64   `json:"count_in_stock"`
}

type ProductRes struct {
	ID           int64      `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Name         string     `json:"name"`
	Image        string     `json:"image"`
	Category     string     `json:"category"`
	Description  string     `json:"description"`
	Rating       int64      `json:"rating"`
	NumReviews   int64      `json:"num_reviews"`
	Price        float64    `json:"price"`
	CountInStock int64      `json:"count_in_stock"`
}
