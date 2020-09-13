package service

import (
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/persistence"
	"log"
	"sync"
)

const orderNum = "ordernum"

// get order by order id
func GetOrderByOrderId(orderId int) (*domain.Order, error) {
	o, err := persistence.GetOrderByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	o.LineItems, err = persistence.GetLineItemsByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	for _, li := range o.LineItems {
		item, err := persistence.GetItem(li.ItemId)
		if err != nil {
			log.Printf("service GetOrderByOrderId GetItem error: %v", err.Error())
			continue
		}
		item.Quantity, err = persistence.GetInventoryQuantity(li.ItemId)
		if err != nil {
			log.Printf("service GetOrderByOrderId GetInventoryQuantity error: %v", err.Error())
			continue
		}
		li.Item = item
		li.CalculateTotal()
	}
	return o, nil
}

// get orders by username
func GetOrdersByUserName(useName string) ([]*domain.Order, error) {
	return persistence.GetOrdersByUserName(useName)
}

// insert order
func InsertOrder(o *domain.Order) error {
	orderId, err := getNextId(orderNum)
	if err != nil {
		return err
	}
	o.OrderId = orderId
	return persistence.InsertOrder(o)
}

var serviceMutex sync.Mutex

// update the sequence and next id
func getNextId(name string) (int, error) {
	// 在并发场景下，这里需要锁
	serviceMutex.Lock()
	defer serviceMutex.Unlock()
	s, err := persistence.GetSequence(name)
	if err != nil {
		return -1, err
	}
	s.NextId++
	err = persistence.UpdateSequence(s)
	if err != nil {
		return -1, err
	}
	return s.NextId, nil
}
