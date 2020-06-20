package domain

type Item struct {
	ItemId     string  `db:"itemid"`
	ProductId  string  `db:"productid"`
	ListPrice  float32 `db:"listprice"`
	UnitCost   float32 `db:"uintcost"`
	SupplierId int     `db:"supplier"`
	Status     string  `db:"status"`
	Attr1      string  `db:"attr1"`
	Attr2      string  `db:"attr2"`
	Attr3      string  `db:"attr3"`
	Attr4      string  `db:"attr4"`
	Attr5      string  `db:"attr5"`
	Product    *Product
	Quantity   int
}

func (*Item) TableName() string {
	return "item"
}
