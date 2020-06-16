package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/service"
	"gopetstore_v2/src/util"
	"net/http"
)

// about View
// 欢迎页
func ViewIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 跳转 主页
func ViewMain(c *gin.Context) {
	a, _ := c.Get("account")
	c.HTML(http.StatusOK, "main.html", gin.H{
		"Account": a,
	})
}

// 跳转 category 分类页
func ViewCategory(c *gin.Context) {
	categoryId := util.GetURLParam(c, "categoryId")[0]
	category, err := service.GetCategory(categoryId)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	products, err := service.GetProductListByCategory(categoryId)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	a, _ := c.Get("account")
	c.HTML(http.StatusOK, "category.html", gin.H{
		"Account":     a,
		"Category":    category,
		"ProductList": products,
	})
}

// 跳转 product 商品页
func ViewProduct(c *gin.Context) {
	productId := util.GetURLParam(c, "productId")[0]
	p, err := service.GetProduct(productId)
	// 将 product 存到 context 中，供中间件进行 session 存储
	c.Set("product", p)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	itemList, err := service.GetItemListByProduct(productId)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	a, _ := c.Get("account")
	c.HTML(http.StatusOK, "product.html", gin.H{
		"Account":  a,
		"Product":  p,
		"ItemList": itemList,
	})
}

// 跳转 item 商品详情页
func ViewItem(c *gin.Context) {
	itemId := util.GetURLParam(c, "itemId")[0]
	item, err := service.GetItem(itemId)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	// 从 中间件中 到 context 中获取 product 并放入页面
	p, ok := c.Get("product")
	if !ok {
		util.ViewError(c, errors.New("can not get the product from session"))
		return
	}
	a, _ := c.Get("account")
	c.HTML(http.StatusOK, "item.html", gin.H{
		"Account": a,
		"Product": p,
		"Item":    item,
	})
}

func SearchProductList(c *gin.Context) {
	keyword := c.PostForm("keyword")
	products, err := service.SearchProductList("%" + keyword + "%")
	if err != nil {
		util.ViewError(c, err)
		return
	}
	a, _ := c.Get("account")
	c.HTML(http.StatusOK, "searchProduct.html", gin.H{
		"Account":     a,
		"ProductList": products,
	})
}
