package persistence

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const getItemByIdSQL = `select I.ITEMID as itemid,LISTPRICE as listprice,UNITCOST as unitcost,SUPPLIER AS supplier,I.PRODUCTID AS productid,
NAME AS name,DESCN AS descn,CATEGORY AS category,STATUS as status,
IFNULL(ATTR1, "") AS attr1,IFNULL(ATTR2, "") AS attr2,IFNULL(ATTR3, "") AS attr3,
IFNULL(ATTR4, "") AS attr4,IFNULL(ATTR5, "") AS attr5,QTY AS quantity from ITEM I, INVENTORY V, PRODUCT P 
where P.PRODUCTID = I.PRODUCTID and I.ITEMID = V.ITEMID and I.ITEMID=?`

const getItemListByProductIdSQL = `SELECT I.ITEMID as itemid,LISTPRICE as listprice,UNITCOST as unitcost,SUPPLIER AS supplier,I.PRODUCTID AS productid,
NAME AS name,DESCN AS descn,CATEGORY AS category,STATUS as status,
IFNULL(ATTR1, "") AS attr1,IFNULL(ATTR2, "") AS attr2,IFNULL(ATTR3, "") AS attr3,
IFNULL(ATTR4, "") AS attr4,IFNULL(ATTR5, "") AS attr5 FROM ITEM I, PRODUCT P 
WHERE P.PRODUCTID = I.PRODUCTID AND I.PRODUCTID = ?`
const getInventoryByItemIdSQL = `SELECT QTY AS quantity FROM INVENTORY WHERE ITEMID = ?`
const updateInventoryByItemIdSQl = `UPDATE INVENTORY SET QTY = QTY - :linequantity WHERE ITEMID = :lineitemid`

// get item by id
func GetItem(itemId string) (*domain.Item, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return nil, err
	}
	i := new(domain.Item)
	p := new(domain.Product)
	i.Product = p
	err = d.Get(i, getItemByIdSQL, itemId)
	if err != nil {
		return nil, err
	}
	return i, nil
}

// get item list by product id
func GetItemListByProduct(productId string) ([]*domain.Item, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	var result []*domain.Item
	if err != nil {
		return result, err
	}
	// 查询外键相关的实体，item 中 有 product
	err = d.Select(&result, getItemListByProductIdSQL, productId)
	return result, err
}

// get inventory by item id
func GetInventoryQuantity(itemId string) (result int, err error) {
	d, err := util.GetConnection()
	if err != nil {
		return -1, err
	}
	r, err := d.Queryx(getInventoryByItemIdSQL, itemId)
	if err != nil {
		return -1, err
	}
	defer r.Close()
	if r.Next() {
		err = r.Scan(&result)
		if err != nil {
			return -1, err
		}
	}
	return
}
