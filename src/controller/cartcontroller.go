package controller

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/service"
	"gopetstore_v2/src/util"
	"log"
	"strconv"
)

// 购物车控制器

// view cart
func ViewCart(c *gin.Context) {
	cart := util.GetCartFromSessionAndSave(c.Writer, c.Request, nil)
	util.ViewWithAccount(c, "cart.html", gin.H{
		"Cart": cart,
	})
}

// update cart
func UpdateCart(c *gin.Context) {
	cart := util.GetCartFromSessionAndSave(c.Writer, c.Request, func(cart *domain.Cart) {
		for _, ci := range cart.ItemList {
			quantityStr := c.PostForm(ci.ItemId)
			var quantity int
			var err error
			if len(quantityStr) != 0 {
				if quantity, err = strconv.Atoi(quantityStr); err != nil {
					log.Printf("UpdateCart error: %v", err.Error())
					continue
				}
			}
			cart.SetQuantityByItemId(ci.ItemId, quantity)
		}
	})
	util.ViewWithAccount(c, "cart.html", gin.H{
		"Cart": cart,
	})
}

// add item to cart
func AddItemToCart(c *gin.Context) {
	itemId := util.GetURLParam(c, "workingItemId")[0]
	i, err := service.GetItem(itemId)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	flag := service.IsItemInStock(itemId)
	cart := util.GetCartFromSessionAndSave(c.Writer, c.Request, func(cart *domain.Cart) {
		cart.AddItem(i, flag)
	})

	util.ViewWithAccount(c, "cart.html", gin.H{
		"Cart": cart,
	})
}

// remove item from cart
func RemoveItemFromCart(c *gin.Context) {
	itemId := util.GetURLParam(c, "workingItemId")[0]
	cart := util.GetCartFromSessionAndSave(c.Writer, c.Request, func(cart *domain.Cart) {
		cart.RemoveItemById(itemId)
	})
	util.ViewWithAccount(c, "cart.html", gin.H{
		"Cart": cart,
	})
}
