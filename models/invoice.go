package models

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	Quantity  int
	ProductId int `bson:"index"`
	Product   Product
}
