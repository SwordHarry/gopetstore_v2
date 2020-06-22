package domain

import (
	"log"
	"testing"
)

func TestCart(t *testing.T) {
	cart := NewCart()
	cart.AddItem(&Item{
		ItemId:     "123",
		ProductId:  "123",
		ListPrice:  12,
		UnitCost:   12,
		SupplierId: 0,
		Status:     "ok",
		Attr1:      "123",
		Attr2:      "123",
		Attr3:      "123",
		Attr4:      "123",
		Attr5:      "123",
		Product:    nil,
		Quantity:   1,
	}, true)
	cart.SetQuantityByItemId("123", 3)
	log.Printf("item list and item map: %v", cart.ItemList[0].Quantity)
}
