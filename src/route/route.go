package route

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/controller"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/", controller.ViewIndex)
	// 路由分组，他们需要进行 用户判断
	g := r.Group("", AccountLogin)
	{
		// catalog
		g.GET("/main", controller.ViewMain)
		g.GET("/viewCategory", controller.ViewCategory)
		g.GET("/viewProduct", controller.ViewProduct)
		g.GET("/viewItem", controller.ViewItem)
		g.POST("/searchProduct", controller.SearchProductList)
		// cart
		// order
	}
	// account
	r.GET("/login", controller.ViewLogin)
	r.GET("/register", controller.ViewRegister)
	r.POST("/newAccount", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/signOut", controller.SignOut)
	r.GET("/editAccount", controller.ViewEditAccount)
}
