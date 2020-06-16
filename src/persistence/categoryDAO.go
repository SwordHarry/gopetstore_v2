package persistence

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
)

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
	r := d.Where("catid = ?", categoryId).Find(c)
	if r.Error != nil {
		return nil, r.Error
	}
	return c, nil
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
	r := d.Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}

	return result, err
}
