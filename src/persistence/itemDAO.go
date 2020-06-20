package persistence

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const getItemByIdSQL = `select I.ITEMID,LISTPRICE,UNITCOST,SUPPLIER AS supplierId,I.PRODUCTID AS productId,
NAME AS productName,DESCN AS productDescription,CATEGORY AS CategoryId,STATUS,
IFNULL(ATTR1, "") AS attribute1,IFNULL(ATTR2, "") AS attribute2,IFNULL(ATTR3, "") AS attribute3,
IFNULL(ATTR4, "") AS attribute4,IFNULL(ATTR5, "") AS attribute5,QTY AS quantity from ITEM I, INVENTORY V, PRODUCT P 
where P.PRODUCTID = I.PRODUCTID and I.ITEMID = V.ITEMID and I.ITEMID=?`

const getItemListByProductIdSQL = `SELECT I.ITEMID,LISTPRICE,UNITCOST,SUPPLIER AS supplierId,I.PRODUCTID AS productId,
NAME AS productName,DESCN AS productDescription,CATEGORY AS categoryId,STATUS,
IFNULL(ATTR1, "") AS attribute1,IFNULL(ATTR2, "") AS attribute2,IFNULL(ATTR3, "") AS attribute3,
IFNULL(ATTR4, "") AS attribute4,IFNULL(ATTR5, "") AS attribute5 FROM ITEM I, PRODUCT P 
WHERE P.PRODUCTID = I.PRODUCTID AND I.PRODUCTID = ?`
const getInventoryByItemIdSQL = `SELECT QTY AS QUANTITY FROM INVENTORY WHERE ITEMID = ?`
const updateInventoryByItemIdSQl = `UPDATE INVENTORY SET QTY = QTY - ? WHERE ITEMID = ?`

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
	// 查询外键相关的实体
	err = d.Get(&result, getItemListByProductIdSQL, productId)
	return result, err
}
