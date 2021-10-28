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
	FeatureId  int `json:"featureId" gorm:"foreignKey=featureId"`
	Category   Category
	Feature    Feature
}
