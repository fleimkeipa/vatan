package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Quantity   int
	Price      float32
	Name       string
	CategoryId int `json:"categoryId" gorm:"foreignKey=categoryId"`
	FeaturesId int `json:"featuresId" gorm:"foreignKey=featuresId"`
	Category   Category
	Features   Features
}
