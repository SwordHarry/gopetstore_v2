package domain

import (
	"encoding/gob"
	"time"
)

type Order struct {
	// order
	OrderId   int
	OrderDate time.Time
	UserName  string
	LineItems []*LineItem
	// ship
	ShipAddress1    string
	ShipAddress2    string
	ShipCity        string
	ShipState       string
	ShipZip         string
	ShipCountry     string
	ShipToFirstName string
	ShipToLastName  string
	// bill
	BillAddress1    string
	BillAddress2    string
	BillCity        string
	BillState       string
	BillZip         string
	BillCountry     string
	BillToFirstName string
	BillToLastName  string
	// other
	Courier    string
	TotalPrice float32
	CreditCard string
	ExpiryDate string
	CardType   string
	Locale     string
	Status     string
}

// 注册序列化
func init() {
	gob.Register(&Order{})
}
