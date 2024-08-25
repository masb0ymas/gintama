package entity

type Order struct {
	Base
	PaymentMethod string      `json:"payment_method" gorm:"type:varchar; not null" validate:"required"`
	TaxPrice      float64     `json:"tax_price" gorm:"type:decimal(10,2); not null" default:"0"`
	ShippingPrice float64     `json:"shipping_price" gorm:"type:decimal(10,2); not null" default:"0"`
	TotalPrice    float64     `json:"total_price" gorm:"type:decimal(10,2); not null" default:"0"`
	OrderItem     []OrderItem `json:"order_items"`
}
