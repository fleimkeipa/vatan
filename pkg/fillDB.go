package pkg

import (
	"strconv"

	"github.com/fleimkeipa/vatansoft/models"
	"gorm.io/gorm"
)

func FillDB(db *gorm.DB) {
	category := [4]models.Category{}
	db.Exec("DELETE FROM categories")
	for i := 0; i < 4; i++ {
		category[i].Name = "name " + strconv.Itoa(i)
	}
	db.Create(&category)

	features := [4]models.Features{}
	db.Exec("DELETE FROM features")
	for i := 0; i < 4; i++ {
		features[i].Desc = "desc " + strconv.Itoa(i)
	}
	db.Create(&features)
}
