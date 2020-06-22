package persistence

import (
	"github.com/jmoiron/sqlx"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
	"log"
)

// DAO for order

const (
	getOrderByOrderIdSQL = `select BILLADDR1 AS billAddress1,BILLADDR2 AS billAddress2,BILLCITY,BILLCOUNTRY,BILLSTATE,BILLTOFIRSTNAME,BILLTOLASTNAME,BILLZIP,
SHIPADDR1 AS shipAddress1,SHIPADDR2 AS shipAddress2,SHIPCITY,SHIPCOUNTRY,SHIPSTATE,SHIPTOFIRSTNAME,SHIPTOLASTNAME,SHIPZIP,CARDTYPE,COURIER,CREDITCARD,
EXPRDATE AS expiryDate,LOCALE,ORDERDATE,ORDERS.ORDERID,TOTALPRICE,USERID AS username,STATUS FROM ORDERS, ORDERSTATUS 
WHERE ORDERS.ORDERID = ? AND ORDERS.ORDERID = ORDERSTATUS.ORDERID`
	getOrdersByUsernameSQL = `SELECT BILLADDR1 AS billAddress1, BILLADDR2 AS billAddress2, BILLCITY, BILLCOUNTRY, BILLSTATE, BILLTOFIRSTNAME, BILLTOLASTNAME, BILLZIP,
SHIPADDR1 AS shipAddress1, SHIPADDR2 AS shipAddress2, SHIPCITY, SHIPCOUNTRY, SHIPSTATE, SHIPTOFIRSTNAME, SHIPTOLASTNAME, SHIPZIP, CARDTYPE, COURIER, CREDITCARD, EXPRDATE AS expiryDate,LOCALE,
ORDERDATE, ORDERS.ORDERID, TOTALPRICE, USERID AS username,STATUS FROM ORDERS, ORDERSTATUS WHERE ORDERS.USERID = ? AND ORDERS.ORDERID = ORDERSTATUS.ORDERID ORDER BY ORDERDATE`
	insertOrderSQL = `INSERT INTO ORDERS (ORDERID, USERID, ORDERDATE, SHIPADDR1, SHIPADDR2, SHIPCITY, SHIPSTATE, SHIPZIP, SHIPCOUNTRY,
BILLADDR1, BILLADDR2, BILLCITY, BILLSTATE, BILLZIP, BILLCOUNTRY, COURIER, TOTALPRICE, BILLTOFIRSTNAME, BILLTOLASTNAME, SHIPTOFIRSTNAME, SHIPTOLASTNAME, CREDITCARD, EXPRDATE, CARDTYPE, LOCALE) 
VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	insertOrderStatusSQL = `INSERT INTO ORDERSTATUS (ORDERID, LINENUM, TIMESTAMP, STATUS) VALUES (?, ?, ?, ?)`
)

// get order by order id
func GetOrderByOrderId(orderId int) (*domain.Order, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	o := new(domain.Order)
	err = d.Get(o, getOrderByOrderIdSQL, orderId)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// get orders by user name
func GetOrdersByUserName(userName string) ([]*domain.Order, error) {
	d, err := util.GetConnection()
	defer func() {
		if d != nil {
			d.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	var result []*domain.Order
	err = d.Select(&result, getOrdersByUsernameSQL, userName)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// insert order
func InsertOrder(o *domain.Order) error {
	// 这里的插入使用事务，插入订单出错则回滚报错
	return util.ExecTransaction(func(tx *sqlx.Tx) error {
		for _, li := range o.LineItems {
			// update inventory by item id
			_, err := tx.NamedExec(updateInventoryByItemIdSQl, li)
			if err != nil {
				log.Printf("service InsertOrder UpdateInventoryQuantity error: %v", err.Error())
				continue
			}
		}
		// insert order info
		_, err := tx.NamedExec(insertOrderSQL, o)
		if err != nil {
			tx.Rollback()
			return err
		}

		// insert order status
		_, err = tx.NamedExec(insertOrderStatusSQL, o)
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, li := range o.LineItems {
			li.OrderId = o.OrderId
			// insert line item
			_, err := tx.NamedExec(insertLineItemSQL, li)
			if err != nil {
				log.Printf("service InsertOrder InsertLineItem error: %v", err.Error())
				continue
			}
		}
		return nil
	})
}
