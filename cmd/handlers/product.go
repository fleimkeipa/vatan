package handlers

import (
	"net/http"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/fleimkeipa/vatansoft/pkg"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ProductReceiver struct {
	DB *gorm.DB
}

func (r *ProductReceiver) GetAll(c echo.Context) error {
	products := []models.Product{}
	r.DB.Preload("Feature").Preload("Category").Find(&products)

	return c.JSON(http.StatusOK, products)
}

func (r *ProductReceiver) Get(c echo.Context) error {
	id := c.Param("id")

	data := models.Product{}
	r.DB.Preload("Feature").Preload("Category").First(&data, "id", id)

	return c.JSON(http.StatusOK, data)
}

func (r *ProductReceiver) Filter(c echo.Context) error {
	params := pkg.TakeAllParams(c.QueryParams())

	products, product := []models.Product{}, models.Product{}
	invoices, invoice := []models.Invoice{}, models.Invoice{}
	products2 := []models.Product{}

	if err := r.DB.Find(&invoices).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	counter := 0

	for i := range params {
		if i == "error" {
			return c.JSON(http.StatusBadRequest, "convert error, filter value must be int")
		}

		if i == "feature" || i == "category" {
			counter++
		}
	}

	if counter == 1 && len(params) == 1 {
		for i, v := range params {
			if i == "feature" || i == "category" { //feature_id || category_id
				if err := r.DB.Preload("Category").Preload("Feature").Find(&products2, i+"_id", v).Error; err != nil {
					return c.JSON(http.StatusBadRequest, err.Error())
				}
			}
		}
	} else if counter == 2 && len(params) == 2 {
		if err := r.DB.Preload("Category").Preload("Feature").Where("feature_id", params["feature"]).Where("category_id", params["category"]).Find(&products2).Error; err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		for i, v := range params {
			if i == "selled" {
				if err := r.DB.Preload("Category").Preload("Feature").First(&invoice, "product_id", v).Error; err != nil {
					return c.JSON(http.StatusBadRequest, err.Error())
				}

				if err := r.DB.Preload("Category").Preload("Feature").First(&product, "id", invoice.ProductId).Error; err != nil {
					return c.JSON(http.StatusBadRequest, err.Error())
				}
				products = append(products, product)
			}

			if i == "deleted" {
				if err := r.DB.Preload("Category").Preload("Feature").Unscoped().Where("deleted_at IS NOT NULL").Find(&product).Error; err != nil {
					return c.JSON(http.StatusBadRequest, err.Error())
				}

				products = append(products, product)
			}

			if i == "feature" || i == "category" { //feature_id || category_id
				if err := r.DB.Preload("Category").Preload("Feature").First(&product, i+"_id", v).Error; err != nil {
					return c.JSON(http.StatusBadRequest, err.Error())
				}
				products = append(products, product)
			}
		}
		if len(products) > 1 {
			temp := products[0]
			for _, v := range products[1:] {
				if temp.ID != v.ID {
					return c.JSON(http.StatusBadRequest, "record not found")
				} else {
					products2 = append(products2, v)
				}
			}
		} else {
			products2 = products
		}
	}

	if products2[0].ID == 0 {
		return c.JSON(http.StatusBadRequest, "record not found")
	}

	return c.JSON(http.StatusOK, products2)
}

func (r *ProductReceiver) Insert(c echo.Context) error {
	data := models.Product{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Create(&data).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Product")
}

func (r *ProductReceiver) Update(c echo.Context) error {
	id := c.Param("id")
	data := models.Product{}
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Model(&data).Where("id", id).Updates(data).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Updated "+id+"'s Product")
}

func (r *ProductReceiver) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := r.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted "+id+"'s Product")
}
