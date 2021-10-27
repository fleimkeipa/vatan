package handlers

import (
	"net/http"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ProductReceiver struct {
	DB *gorm.DB
}

func (r *ProductReceiver) GetAll(c echo.Context) error {
	products := []models.Product{}
	categories, features := []models.Category{}, []models.Features{}

	r.DB.Find(&products)
	r.DB.Find(&categories)
	r.DB.Find(&features)

	for i, v := range products {
		//categoryId == categories.id
		for _, v2 := range categories {
			if v.CategoryId == int(v2.ID) {
				products[i].Category = v2
				break
			}
		}
		//featureId == features.id
		for _, v2 := range features {
			if v.FeaturesId == int(v2.ID) {
				products[i].Features = v2
				break
			}
		}
	}

	c.JSON(http.StatusOK, products)
	return nil
}

func (r *ProductReceiver) Get(c echo.Context) error {
	id := c.Param("id")
	categories, features := []models.Category{}, []models.Features{}

	r.DB.Find(&categories)
	r.DB.Find(&features)

	data := models.Product{}
	if err := r.DB.First(&data, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//categoryId == categories.id
	for _, v2 := range categories {
		if data.CategoryId == int(v2.ID) {
			data.Category = v2
			break
		}
	}
	//featureId == features.id
	for _, v2 := range features {
		if data.FeaturesId == int(v2.ID) {
			data.Features = v2
			break
		}
	}

	return c.JSON(http.StatusOK, data)
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
