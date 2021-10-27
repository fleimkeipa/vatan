package handlers

import (
	"net/http"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type FeaturesReceiver struct {
	DB *gorm.DB
}

func (r *FeaturesReceiver) GetAll(c echo.Context) error {
	features := []models.Features{}
	if err := r.DB.Find(&features).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, features)
}

func (r *FeaturesReceiver) Get(c echo.Context) error {
	id := c.Param("id")

	features := models.Features{}
	if err := r.DB.First(&features, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, features)
}

func (r *FeaturesReceiver) Insert(c echo.Context) error {
	data := models.Features{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Create(&data).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Features")
}

func (r *FeaturesReceiver) Update(c echo.Context) error {
	id := c.Param("id")

	features := models.Features{}
	if err := c.Bind(&features); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Model(&features).Where("id", id).Update("name", features.Desc).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Updated "+id+"'s Features")
}

func (r *FeaturesReceiver) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := r.DB.Delete(&models.Features{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted "+id+"'s Features")
}
