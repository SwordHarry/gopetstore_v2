package persistence

import (
	"errors"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

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
	r := d.Where("category = ?", categoryId).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return result, nil
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
	r := d.Where("productid = ?", productId).Find(p)
	if r.Error != nil {
		return nil, r.Error
	}
	if r.RecordNotFound() {
		return nil, errors.New("can not find the product by this product id")
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
	r := d.Where("lower(name) like ?", keyword).Find(&result)
	if r.Error != nil {
		return result, r.Error
	}
	return result, nil
}
