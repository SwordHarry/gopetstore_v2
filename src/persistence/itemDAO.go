package persistence

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

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
	d.Where("itemid = ?", itemId).Find(i)
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
	r := d.Where("productid = ?", productId).Find(&result)
	if r.Error != nil {
		return result, r.Error
	}
	return result, nil
}
