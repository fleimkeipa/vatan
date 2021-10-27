package handlers

import (
	"net/http"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type CategoryReceiver struct {
	DB *gorm.DB
}

func (r *CategoryReceiver) GetAll(c echo.Context) error {
	categories := []models.Category{}
	if err := r.DB.Find(&categories).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, categories)
}

func (r *CategoryReceiver) Get(c echo.Context) error {
	id := c.Param("id")

	category := models.Category{}
	if err := r.DB.First(&category, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, category)
}

func (r *CategoryReceiver) Insert(c echo.Context) error {
	data := models.Category{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Create(&data).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Category")
}

func (r *CategoryReceiver) Update(c echo.Context) error {
	id := c.Param("id")

	category := models.Category{}
	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Model(&category).Where("id", id).Update("name", category.Name).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Updated "+id+"'s Category")
}

func (r *CategoryReceiver) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := r.DB.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted "+id+"'s Category")
}
