package domain

import (
	"encoding/gob"
	"time"
)

type Order struct {
	// order
	OrderId   int       `db:"orderid"`
	OrderDate time.Time `db:"orderdate"`
	UserName  string    `db:"userid"`
	LineItems []*LineItem
	// ship
	ShipAddress1    string `db:"shipaddr1"`
	ShipAddress2    string `db:"shipaddr2"`
	ShipCity        string `db:"shipcity"`
	ShipState       string `db:"shipstate"`
	ShipZip         string `db:"shipzip"`
	ShipCountry     string `db:"shipcountry"`
	ShipToFirstName string `db:"shiptofirstname"`
	ShipToLastName  string `db:"shiptolastname"`
	// bill
	BillAddress1    string `db:"billaddr1"`
	BillAddress2    string `db:"billaddr2"`
	BillCity        string `db:"billcity"`
	BillState       string `db:"billstate"`
	BillZip         string `db:"billzip"`
	BillCountry     string `db:"billcountry"`
	BillToFirstName string `db:"billtofirstname"`
	BillToLastName  string `db:"billtolastname"`
	// other
	Courier    string  `db:"courier"`
	TotalPrice float32 `db:"totalprice"`
	CreditCard string  `db:"creditcard"`
	ExpiryDate string  `db:"expdate"`
	CardType   string  `db:"cardtype"`
	Locale     string  `db:"locale"`
	Status     string  `db:"status"`
	// 总货品数
	TotalLineNum int `db:"totallinenum"`
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
