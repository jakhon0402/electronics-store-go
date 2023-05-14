package models

import "electronics-store-go/internal/app/models/templates"

type Product struct {
	templates.UUIDModel
	Name          string  `gorm:"not null"`
	Title         string  `gorm:"not null"`
	Price         float32 `gorm:"not null"`
	Specification string  `gorm:"not null"`
}
