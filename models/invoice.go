package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Quantity  int
	Total     float32
	ProductId int `bson:"index"`
	Product   Product
}
