package handlers

import (
	"net/http"

	"github.com/fleimkeipa/vatansoft/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type FeatureReceiver struct {
	DB *gorm.DB
}

func (r *FeatureReceiver) GetAll(c echo.Context) error {
	feature := []models.Feature{}
	if err := r.DB.Find(&feature).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, feature)
}

func (r *FeatureReceiver) Get(c echo.Context) error {
	id := c.Param("id")

	feature := models.Feature{}
	if err := r.DB.First(&feature, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, feature)
}

func (r *FeatureReceiver) Insert(c echo.Context) error {
	data := models.Feature{}
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Create(&data).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Inserted Feature")
}

func (r *FeatureReceiver) Update(c echo.Context) error {
	id := c.Param("id")

	feature := models.Feature{}
	if err := c.Bind(&feature); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := r.DB.Model(&feature).Where("id", id).Update("name", feature.Desc).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Updated "+id+"'s Feature")
}

func (r *FeatureReceiver) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := r.DB.Delete(&models.Feature{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Deleted "+id+"'s Feature")
}
