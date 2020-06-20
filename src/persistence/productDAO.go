package persistence

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const getProductListByCategorySQL = "SELECT PRODUCTID,NAME,DESCN as description,CATEGORY as categoryId FROM PRODUCT WHERE CATEGORY = ?"
const getProductByIdSQL = "SELECT PRODUCTID,NAME,DESCN as description,CATEGORY as categoryId FROM PRODUCT WHERE PRODUCTID = ?"
const getProductListByKeyword = "select PRODUCTID,NAME,DESCN as description,CATEGORY as categoryId from PRODUCT WHERE lower(name) like ?"

// get product list by category id
func GetProductListByCategory(categoryId string) ([]*domain.Product, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	var result []*domain.Product
	if err != nil {
		return result, err
	}
	err = d.Get(&result, getProductListByCategorySQL, categoryId)
	return result, err
}

// get product by product id
func GetProduct(productId string) (*domain.Product, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return nil, err
	}
	p := new(domain.Product)
	err = d.Get(p, getProductByIdSQL, productId)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// 通过关键字获取 product
func SearchProductList(keyword string) ([]*domain.Product, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	var result []*domain.Product
	if err != nil {
		return result, err
	}
	err = d.Get(&result, getProductListByKeyword, keyword)
	return result, err
}
