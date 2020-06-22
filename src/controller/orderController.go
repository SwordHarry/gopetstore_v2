package controller

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/util"
	"net/http"
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
	c.HTML(http.StatusOK, "initOrder.html", gin.H{
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
		util.ViewWithAccount(c, "shipForm.html", dataMap)
	} else {
		// view confirmOrder
		util.ViewWithAccount(c, "confirmOrder.html", dataMap)
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
	util.ViewWithAccount(c, "confirmOrder.html", gin.H{
		"Order": order,
	})
}

// create the final order
func ConfirmOrderStep2(c *gin.Context) {
	order := util.GetOrderFromSessionAndSave(c.Writer, c.Request, nil)

}
