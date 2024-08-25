package entity

import "github.com/google/uuid"

type OrderItem struct {
	Base
	OrderId   uuid.UUID `json:"order_id" gorm:"type:uuid; not null" validate:"required"`
	ProductId uuid.UUID `json:"product_id" gorm:"type:uuid; not null" validate:"required"`
	Name      string    `json:"name" gorm:"type:varchar; not null" validate:"required"`
	Quantity  int64     `json:"quantity" gorm:"type:int; not null" default:"0"`
	Image     string    `json:"image" gorm:"type:varchar; not null" validate:"required"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2); not null" default:"0"`
	Order     Order     `json:"order"`
	Product   Product   `json:"product"`
}
