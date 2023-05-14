package models

import "electronics-store-go/internal/app/models/templates"

type Category struct {
	templates.UUIDModel
	Name string `gorm:"not null"`
}
