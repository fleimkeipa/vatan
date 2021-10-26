package models

type InvoceProduct struct {
	Quantity  int
	InvoceId  int `gorm:"foreignKey:invoceId"`
	ProductId int `gorm:"foreignKey:productId"`
}
