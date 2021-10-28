package handlers

import (
	"net/http"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/fleimkeipa/vatansoft/pkg"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type InvoceReceiver struct {
	DB *gorm.DB
}

func (r *InvoceReceiver) GetAll(c echo.Context) error {
	invoces := []models.Invoice{}

	if err := r.DB.Preload("Product").Preload("Product.Category").Preload("Product.Feature").Find(&invoces).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, invoces)
}

func (r *InvoceReceiver) Get(c echo.Context) error {
	id := c.Param("id")

	invoice := models.Invoice{}

	if err := r.DB.Preload("Product").Preload("Product.Category").Preload("Product.Feature").First(&invoice, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, invoice)
}

func (r *InvoceReceiver) Insert(c echo.Context) error {
	data := models.Invoice{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := pkg.QuantityCheck(r.DB, &data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Invoice")
}

func (r *InvoceReceiver) Update(c echo.Context) error {
	id := c.Param("id")

	invoice := models.Invoice{}
	if err := c.Bind(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Model(&invoice).Where("id", id).Updates(invoice).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Updated "+id+"'s Invoice")
}

func (r *InvoceReceiver) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := r.DB.Delete(&models.Invoice{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted "+id+"'s Invoice")
}
