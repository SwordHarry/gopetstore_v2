package domain

type Item struct {
	ItemId     string   `gorm:"column:itemid"`
	ProductId  string   `gorm:"column:productid"`
	ListPrice  float32  `gorm:"column:listprice"`
	UnitCost   float32  `gorm:"column:uintcost"`
	SupplierId int      `gorm:"column:supplier"`
	Status     string   `gorm:"column:status"`
	Attr1      string   `gorm:"column:attr1"`
	Attr2      string   `gorm:"column:attr2"`
	Attr3      string   `gorm:"column:attr3"`
	Attr4      string   `gorm:"column:attr4"`
	Attr5      string   `gorm:"column:attr5"`
	Product    *Product `gorm:"foreignKey:ProductId"`
	Quantity   int
}

func (*Item) TableName() string {
	return "item"
}
