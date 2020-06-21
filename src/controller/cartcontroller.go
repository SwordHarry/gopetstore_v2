package controller

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/util"
	"net/http"
)

// 购物车控制器

func ViewCart(c *gin.Context) {
	cart := util.GetCartFromSession(c.Writer, c.Request, nil)
	a := util.GetAccountFromSession(c.Request)
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"Account": a,
		"Cart":    cart,
	})
}
