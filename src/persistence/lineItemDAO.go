package persistence

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

// DAO for line item of an order

const (
	getLineItemsByOrderIdSQL = `SELECT ORDERID, LINENUM AS lineNumber, ITEMID, QUANTITY, UNITPRICE FROM LINEITEM WHERE ORDERID = ?`
	insertLineItemSQL        = `INSERT INTO LINEITEM (ORDERID, LINENUM, ITEMID, QUANTITY, UNITPRICE) VALUES (?, ?, ?, ?, ?)`
)

// get line item by order id
func GetLineItemsByOrderId(orderId int) ([]*domain.LineItem, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	var result []*domain.LineItem
	err = d.Select(&result, getLineItemsByOrderIdSQL, orderId)
	if err != nil {
		return result, err
	}
	return result, nil
}
