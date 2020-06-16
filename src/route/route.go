package route

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/controller"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", controller.ViewIndex)
	// catalog
	r.GET("/main", controller.ViewMain)
	r.GET("/viewCategory", controller.ViewCategory)
	r.GET("/viewProduct", SessionMiddleWare("product"), controller.ViewProduct)
	r.GET("/viewItem", SessionMiddleWare("product"), controller.ViewItem)
	r.POST("/searchProduct", controller.SearchProductList)
	// account
	r.GET("/login", controller.ViewLogin)
	r.GET("/register", controller.ViewRegister)
	// cart
	// order
}
