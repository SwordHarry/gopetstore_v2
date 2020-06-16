package service

// service for category product and item, as catalogService
import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/persistence"
)

// get category by id
func GetCategory(categoryId string) (*domain.Category, error) {
	return persistence.GetCategory(categoryId)
}

// get product list by category id
func GetProductListByCategory(categoryId string) ([]*domain.Product, error) {
	return persistence.GetProductListByCategory(categoryId)
}

// get product by id
func GetProduct(productId string) (*domain.Product, error) {
	return persistence.GetProduct(productId)
}

// get item list by product id
func GetItemListByProduct(productId string) ([]*domain.Item, error) {
	return persistence.GetItemListByProduct(productId)
}

// get item by item id
func GetItem(itemId string) (*domain.Item, error) {
	return persistence.GetItem(itemId)
}

// search product list by keyword
func SearchProductList(keyword string) ([]*domain.Product, error) {
	return persistence.SearchProductList(keyword)
}
