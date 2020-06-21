package persistence

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

const getCategoryListSQL = "SELECT CATID as catid,NAME as catname,descn FROM CATEGORY"
const getCategoryByIdSQL = "SELECT CATID as catid,NAME as catname,descn FROM CATEGORY WHERE CATID = ?"

// 通过 id 获取指定的 category
func GetCategory(categoryId string) (*domain.Category, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	if err != nil {
		return nil, err
	}
	c := new(domain.Category)
	err = d.Get(c, getCategoryByIdSQL, categoryId)
	if err != nil {
		return nil, err
	}
	return c, err
}

// 获取所有的 category
func GetCategoryList() ([]*domain.Category, error) {
	d, err := util.GetConnection()
	defer func() {
		_ = d.Close()
	}()
	var result []*domain.Category
	if err != nil {
		return result, err
	}
	// 获取所有 category，这里即使是 slice 也要取地址
	err = d.Get(&result, getCategoryListSQL)
	return result, err
}
