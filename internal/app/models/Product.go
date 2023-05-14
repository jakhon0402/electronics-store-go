package models

import "electronics-store-go/internal/app/models/templates"

type Product struct {
	templates.UUIDModel
	Name          string  `gorm:"not null" json:"name"`
	Title         string  `gorm:"not null" json:"title"`
	Price         float32 `gorm:"not null" json:"price"`
	Specification string  `gorm:"not null" json:"specification"`
}
