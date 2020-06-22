package controller

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/service"
	"gopetstore_v2/src/util"
	"log"
	"net/http"
	"strconv"
)

// 购物车控制器

// view cart
func ViewCart(c *gin.Context) {
	cart := util.GetCartFromSession(c.Writer, c.Request, nil)
	a, _ := c.Get(config.AccountKey)
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"Account": a,
		"Cart":    cart,
	})
}

// update cart
func UpdateCart(c *gin.Context) {
	cart := util.GetCartFromSession(c.Writer, c.Request, func(cart *domain.Cart) {
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
	c.HTML(http.StatusOK, "cart.html", gin.H{
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
	cart := util.GetCartFromSession(c.Writer, c.Request, func(cart *domain.Cart) {
		cart.AddItem(i, flag)
	})
	a, _ := c.Get(config.AccountKey)
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"Account": a,
		"Cart":    cart,
	})
}

// remove item from cart
func RemoveItemFromCart(c *gin.Context) {
	itemId := util.GetURLParam(c, "workingItemId")[0]
	cart := util.GetCartFromSession(c.Writer, c.Request, func(cart *domain.Cart) {
		cart.RemoveItemById(itemId)
	})
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"Cart": cart,
	})
}
