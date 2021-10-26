package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Quantity   int
	Price      float32
	Name       string
	CategoryId string `gorm:"foreignKey:id"`
	Category   Category
}
