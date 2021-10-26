package main

import (
	"github.com/fleimkeipa/vatansoft/cmd/handlers"
	"github.com/fleimkeipa/vatansoft/pkg"
	"github.com/labstack/echo"
)

func main() {
	db := pkg.InitDB()
	//pkg.Migration(db.DB)

	product := handlers.ProductReceiver{
		DB: db,
	}
	category := handlers.CategoryReceiver{
		DB: db,
	}

	r := echo.New()

	r.GET("/stocks", product.Get)
	r.GET("/stocks/filter", nil) // query params kullanılmalıdır.
	r.POST("/stock/insert", nil)
	r.PUT("/stock/:id/update", nil)
	r.DELETE("/stock/:id/delete", nil)
	r.GET("/stock/:id", nil) // tekil bir ürün getirilecek.

	r.POST("/stock/category/insert", nil)
	r.DELETE("/stock/:id/category/:id/delete", nil)
	r.DELETE("/stock/:id/category/delete", nil)

	r.GET("/categories", category.GetAll)
	r.GET("/category/:id", category.Get)
	r.POST("/category/insert", category.Insert)
	r.DELETE("/category/:id/delete", category.Delete)
	r.PUT("/category/:id/update", category.Update)

	r.Logger.Fatal(r.Start(":8080"))
}

/*
	db.AutoMigrate(&models.Category{})
	db.Create(&models.Category{
		Name: "category1",
	})
*/
