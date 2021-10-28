package models

type ProductFeature struct {
	ProductId string `gorm:"foreignKey:productId"`
	FeatureId string `gorm:"foreignKey:invoceId"`
}
