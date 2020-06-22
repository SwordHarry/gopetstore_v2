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

func NewOrder(a *Account, c *Cart) *Order {
	o := &Order{
		OrderDate:  time.Now(),
		UserName:   a.UserName,
		LineItems:  []*LineItem{},
		TotalPrice: c.GetSubTotal(),
		// ship
		ShipAddress1:    a.Address1,
		ShipAddress2:    a.Address2,
		ShipCity:        a.City,
		ShipState:       a.State,
		ShipZip:         a.Zip,
		ShipCountry:     a.Country,
		ShipToFirstName: a.FirstName,
		ShipToLastName:  a.LastName,
		// bill
		BillAddress1:    a.Address1,
		BillAddress2:    a.Address2,
		BillCity:        a.City,
		BillState:       a.State,
		BillZip:         a.Zip,
		BillCountry:     a.Country,
		BillToFirstName: a.FirstName,
		BillToLastName:  a.LastName,
		// other
		Courier:    "UPS",
		CreditCard: "999 9999 9999 9999",
		ExpiryDate: "12/03",
		CardType:   "Visa",
		Locale:     "CA",
		Status:     "P",
	}
	for _, ci := range c.ItemList {
		li := NewLineItem(len(o.LineItems)+1, ci)
		o.LineItems = append(o.LineItems, li)
	}
	return o
}
