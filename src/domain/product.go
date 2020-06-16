package domain

import "encoding/gob"

type Product struct {
	ProductId   string `gorm:"column:productid"`
	CategoryId  string `gorm:"column:category;foreignkey:catid"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:descn"`
}

func (*Product) TableName() string {
	return "product"
}

func init() {
	gob.Register(&Product{})
}
