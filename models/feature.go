package models

import "gorm.io/gorm"

type Feature struct {
	gorm.Model
	Desc string
}
