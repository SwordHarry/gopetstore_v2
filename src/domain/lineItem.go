package domain

// 订单项
type LineItem struct {
	OrderId    int     `db:"orderid"`
	LineNumber int     `db:"linenum"`
	Quantity   int     `db:"linequantity"`
	ItemId     string  `db:"lineitemid"`
	UnitPrice  float64 `db:"unitprice"`
	Total      float64
	*Item
}

func NewLineItem(lineNum int, cartItem *CartItem) *LineItem {
	li := &LineItem{
		LineNumber: lineNum,
		ItemId:     cartItem.Item.ItemId,
		UnitPrice:  cartItem.Item.ListPrice,
		Quantity:   cartItem.Quantity,
		Item:       cartItem.Item,
	}
	return li
}

func (li *LineItem) CalculateTotal() {
	if li != nil && li.Quantity > 0 {
		li.Total = li.UnitPrice * float64(li.Quantity)
	}
}
