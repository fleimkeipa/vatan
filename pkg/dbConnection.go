package pkg

import (
	"os"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := "" +
		os.Getenv("user") + ":" +
		os.Getenv("password") +
		"@tcp(" + os.Getenv("host") + ")/" +
		os.Getenv("db") +
		"?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error() + "failed to connection db")
	}
	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Invoice{}, &models.Feature{}, &models.Category{})
	db.AutoMigrate(&models.Product{})
}
