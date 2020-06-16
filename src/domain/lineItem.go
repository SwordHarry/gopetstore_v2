package domain

// 订单项
type LineItem struct {
	OrderId    int
	LineNumber int
	Quantity   int
	ItemId     string
	UnitPrice  float32
	Total      float32
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
		li.Total = li.UnitPrice * float32(li.Quantity)
	}
}
