package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/domain"
	"gopetstore_v2/src/service"
	"gopetstore_v2/src/util"
	"log"
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
	log.Printf("%+v", a)
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
	// 将 product 存到 session 中
	s, err := util.GetSession(c.Request)
	if err != nil {
		log.Printf("ViewProduct get session error: %v", err.Error())
	}
	if s != nil {
		err = s.Save(config.ProductKey, p, c.Writer, c.Request)
		if err != nil {
			log.Printf("ViewProduct session save error: %v", err.Error())
		}
	}

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
	// 从 session 中获取 product 并存入页面
	s, err := util.GetSession(c.Request)
	if err != nil {
		util.ViewError(c, err)
		return
	}
	var p *domain.Product
	if s != nil {
		r, ok := s.Get(config.ProductKey)
		if ok {
			p = r.(*domain.Product)
		} else {
			util.ViewError(c, errors.New("ViewItem: type translation to *domain.Product is failed"))
			return
		}
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
