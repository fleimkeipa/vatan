package pkg

import (
	"errors"

	"github.com/fleimkeipa/vatansoft/models"
	"gorm.io/gorm"
)

func QuantityCheck(r *gorm.DB, data *models.Invoice) error {
	data2 := models.Product{}
	if err := r.First(&data2, "id", data.ProductId).Error; err != nil {
		return err
	}

	if data2.Quantity < data.Quantity {
		return errors.New("quantity can't bigger stock")
	}

	data.Total = data2.Price * float32(data.Quantity)

	if err := r.Create(&data).Error; err != nil {
		return err
	}

	if err := r.Model(&data2).Where("id", data.ProductId).Update("quantity", data2.Quantity-data.Quantity).Error; err != nil {
		return err
	}
	return nil
}
