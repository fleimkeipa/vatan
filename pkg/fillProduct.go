package pkg

import (
	"strconv"

	"github.com/fleimkeipa/vatansoft/models"
	"gorm.io/gorm"
)

func FillProduct(r *gorm.DB, idCame string) []models.Product {
	products := []models.Product{}
	categories, feature := []models.Category{}, []models.Feature{}

	id, err := strconv.Atoi(idCame)
	if err != nil || idCame == "" {
		return nil
	}

	r.Find(&categories)
	r.Find(&feature)

	if id == 0 {
		if err := r.Find(&products).Error; err != nil {
			return nil
		}

		for i, v := range products {
			for _, v2 := range categories {
				if v.CategoryId == int(v2.ID) {
					products[i].Category = v2
					break
				}
			}

			for _, v2 := range feature {
				if v.FeatureId == int(v2.ID) {
					products[i].Feature = v2
					break
				}
			}
		}
	} else {
		if err := r.First(&products, id).Error; err != nil {
			return nil
		}
		for _, v2 := range categories {
			if products[0].CategoryId == int(v2.ID) {
				products[0].Category = v2
				break
			}
		}
		//featureId == feature.id
		for _, v2 := range feature {
			if products[0].FeatureId == int(v2.ID) {
				products[0].Feature = v2
				break
			}
		}
	}
	return products
}

func FillProductForInvoice(r *gorm.DB, invoice []models.Invoice) []models.Invoice {
	products := []models.Product{}
	categories, feature := []models.Category{}, []models.Feature{}

	if err := r.Find(&products).Error; err != nil {
		return nil
	}
	r.Find(&categories)
	r.Find(&feature)

	for i, v := range invoice {
		for _, v1 := range products {
			if v.ProductId == int(v1.ID) {
				invoice[i].Product = v1
				for _, v2 := range categories {
					if v1.CategoryId == int(v2.ID) {
						invoice[i].Product.Category = v2
						break
					}
				}

				for _, v2 := range feature {
					if v1.FeatureId == int(v2.ID) {
						invoice[i].Product.Feature = v2
						break
					}
				}
				break
			}
		}
	}
	return invoice
}
