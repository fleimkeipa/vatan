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
	feature := handlers.FeatureReceiver{
		DB: db,
	}
	invoce := handlers.InvoceReceiver{
		DB: db,
	}

	r := echo.New()

	r.GET("/stocks", product.GetAll)
	r.GET("/stock/:id", product.Get)

	r.GET("/stocks/filter", product.Filter)
	r.POST("/stock/insert", product.Insert)
	r.PUT("/stock/:id/update", product.Update)
	r.DELETE("/stock/:id/delete", product.Delete)

	r.POST("/stock/category/insert", nil)
	r.DELETE("/stock/:id/category/:id/delete", nil)
	r.DELETE("/stock/:id/category/delete", nil)

	r.GET("/invoices", invoce.GetAll)
	r.GET("/invoice/:id", invoce.Get)
	r.POST("/invoice/insert", invoce.Insert)
	r.PUT("/invoice/:id/update", invoce.Update)
	r.DELETE("/invoice/:id/delete", invoce.Delete)

	r.GET("/categories", category.GetAll)
	r.GET("/category/:id", category.Get)
	r.POST("/category/insert", category.Insert)
	r.PUT("/category/:id/update", category.Update)
	r.DELETE("/category/:id/delete", category.Delete)

	r.GET("/feature", feature.GetAll)
	r.GET("/feature/:id", feature.Get)
	r.POST("/feature/insert", feature.Insert)
	r.PUT("/feature/:id/update", feature.Update)
	r.DELETE("/feature/:id/delete", feature.Delete)

	r.Logger.Fatal(r.Start(":8080"))
}
