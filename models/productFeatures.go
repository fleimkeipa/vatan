package models

type ProductFeatures struct {
	ProductId  string `gorm:"foreignKey:productId"`
	FeaturesId string `gorm:"foreignKey:invoceId"`
}
