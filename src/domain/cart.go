package domain

import "encoding/gob"

// 由于购物车没有数据库模块，故维护一个集合，加快查询速度
// list 和 map 需要同步，故不暴露
type Cart struct {
	ItemList []*CartItem
}

// 注册学序列化存到 session
func init() {
	gob.Register(&Cart{})
}

// 构造函数
func NewCart() *Cart {
	return &Cart{
		//ItemMap: map[string]*CartItem{},
		ItemList: []*CartItem{},
	}
}

// 判断购物车中是否有指定id商品
func (c *Cart) ContainItem(itemId string) (*CartItem, bool) {
	//_, ok := c.ItemMap[itemId]
	for _, ci := range c.ItemList {
		if ci.Item.ItemId == itemId {
			return ci, true
		}
	}
	return nil, false
}

// 往购物车添加商品
func (c *Cart) AddItem(item Item, isInStock bool) {
	if ci, ok := c.ContainItem(item.ItemId); ok {
		ci.IncrementQuantity()
	} else {
		ci := &CartItem{
			Item:     &item,
			Quantity: 1,
			InStock:  isInStock,
			Total:    0,
		}
		ci.CalculateTotal()
		c.ItemList = append(c.ItemList, ci)
	}
}

// 移除购物车项
func (c *Cart) RemoveItemById(itemId string) *Item {
	for i, v := range c.ItemList {
		if v.Item.ItemId == itemId {
			c.ItemList = append(c.ItemList[:i], c.ItemList[i+1:]...)
			// 删除中间1个元素
			return v.Item
		}
	}
	return nil
}

// 获取item数量
func (c *Cart) GetNumberOfItems() int {
	return len(c.ItemList)
}

// 通过item id 增加数量
func (c *Cart) IncrementQuantityByItemId(itemId string) {
	ci, ok := c.ContainItem(itemId)
	if ok {
		ci.IncrementQuantity()
	}
}

// 直接设置该项的数量
func (c *Cart) SetQuantityByItemId(itemId string, quantity int) {
	ci, _ := c.ContainItem(itemId)
	ci.Quantity = quantity
	ci.CalculateTotal()
}

// 获取购物车总价格
func (c *Cart) GetSubTotal() (subTotal float32) {
	for _, ci := range c.ItemList {
		ci.CalculateTotal()
		subTotal += ci.Total
	}
	return
}
