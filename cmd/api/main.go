package main

import (
	"github.com/fleimkeipa/vatansoft/cmd/handlers"
	"github.com/fleimkeipa/vatansoft/pkg"
	"github.com/labstack/echo"
)

func main() {
	db := pkg.InitDB()

	//pkg.Migration(db)
	//pkg.FillDB(db)

	product := handlers.ProductReceiver{
		DB: db,
	}
	category := handlers.CategoryReceiver{
		DB: db,
	}
	features := handlers.FeaturesReceiver{
		DB: db,
	}

	r := echo.New()

	r.GET("/stocks", product.GetAll)
	r.GET("/stock/:id", product.Get) // tekil bir ürün getirilecek.
	r.GET("/stocks/filter", nil)     // query params kullanılmalıdır.
	r.POST("/stock/insert", product.Insert)
	r.PUT("/stock/:id/update", product.Update)
	r.DELETE("/stock/:id/delete", product.Delete)

	r.POST("/stock/category/insert", nil)
	r.DELETE("/stock/:id/category/:id/delete", nil)
	r.DELETE("/stock/:id/category/delete", nil)

	r.GET("/categories", category.GetAll)
	r.GET("/category/:id", category.Get)
	r.POST("/category/insert", category.Insert)
	r.PUT("/category/:id/update", category.Update)
	r.DELETE("/category/:id/delete", category.Delete)

	r.GET("/features", features.GetAll)
	r.GET("/features/:id", features.Get)
	r.POST("/features/insert", features.Insert)
	r.PUT("/features/:id/update", features.Update)
	r.DELETE("/features/:id/delete", features.Delete)

	r.Logger.Fatal(r.Start(":8080"))
}
