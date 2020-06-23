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

// file name
const (
	initOrderFile    = "initOrder.html"
	shipFormFile     = "shipForm.html"
	confirmOrderFile = "confirmOrder.html"
	viewOrderFile    = "viewOrder.html"
	listOrdersFile   = "listOrders.html"
)

// view init order
func ViewInitOrder(c *gin.Context) {
	account := util.GetAccountFromSession(c.Request)
	cart := util.GetCartFromSessionAndSave(c.Writer, c.Request, nil)
	o := domain.NewOrder(account, cart)
	s, err := util.GetSession(c.Request)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	if s != nil {
		err = s.Save(config.OrderKey, o, c.Writer, c.Request)
		if err != nil {
			util.ViewError(c, err)
			return
		}
	}
	c.HTML(http.StatusOK, initOrderFile, gin.H{
		"Account":         account,
		"Order":           o,
		"CreditCardTypes": []string{o.CardType},
	})
}

// confirm order step 1
func ConfirmOrderStep1(c *gin.Context) {
	order := util.GetOrderFromSessionAndSave(c.Writer, c.Request, func(order *domain.Order) {
		order.CardType = c.PostForm("cardType")
		order.CreditCard = c.PostForm("creditCard")
		order.ExpiryDate = c.PostForm("expiryDate")
		order.BillToFirstName = c.PostForm("firstName")
		order.BillToLastName = c.PostForm("lastName")
		order.BillAddress1 = c.PostForm("address1")
		order.BillAddress2 = c.PostForm("address2")
		order.BillCity = c.PostForm("city")
		order.BillState = c.PostForm("state")
		order.BillZip = c.PostForm("zip")
		order.BillCountry = c.PostForm("country")
	})
	dataMap := gin.H{
		"Order": order,
	}
	if len(c.PostForm("shippingAddressRequired")) > 0 {
		// view shipForm
		util.ViewWithAccount(c, shipFormFile, dataMap)
	} else {
		// view confirmOrder
		util.ViewWithAccount(c, confirmOrderFile, dataMap)
	}
}

// confirm ship
func ConfirmShip(c *gin.Context) {
	// get order from session
	order := util.GetOrderFromSessionAndSave(c.Writer, c.Request, func(o *domain.Order) {
		o.ShipToFirstName = c.PostForm("shipToFirstName")
		o.ShipToLastName = c.PostForm("shipToLastName")
		o.ShipAddress1 = c.PostForm("shipAddress1")
		o.ShipAddress2 = c.PostForm("shipAddress2")
		o.ShipCity = c.PostForm("shipCity")
		o.ShipState = c.PostForm("shipState")
		o.ShipZip = c.PostForm("shipZip")
		o.ShipCountry = c.PostForm("shipCountry")
	})
	util.ViewWithAccount(c, confirmOrderFile, gin.H{
		"Order": order,
	})
}

// create the final order
func ConfirmOrderStep2(c *gin.Context) {
	order := util.GetOrderFromSessionAndSave(c.Writer, c.Request, nil)
	err := service.InsertOrder(order)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	// 清空购物车
	s, err := util.GetSession(c.Request)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	err = s.Del(config.CartKey, c.Writer, c.Request)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	util.ViewWithAccount(c, viewOrderFile, gin.H{
		"Order": order,
	})
}

// list orders
func ListOrders(c *gin.Context) {
	a := util.GetAccountFromSession(c.Request)
	log.Print(a.UserName)
	orders, err := service.GetOrdersByUserName(a.UserName)
	if err != nil {
		util.ViewError(c, err)
	}
	util.ViewWithAccount(c, listOrdersFile, gin.H{
		"OrderList": orders,
	})
}

// check order
func CheckOrder(c *gin.Context) {
	orderIdStr := util.GetURLParam(c, "orderId")[0]
	orderId, err := strconv.Atoi(orderIdStr)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	o, err := service.GetOrderByOrderId(orderId)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	util.ViewWithAccount(c, viewOrderFile, gin.H{
		"Order": o,
	})
}
