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

func (db *ProductReceiver) Get(c echo.Context) error {
	stocks := []models.Category{}
	db.DB.Find(&stocks)
	c.JSON(http.StatusOK, stocks)
	return nil
}
