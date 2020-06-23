package util

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"html/template"
	"net/http"
)

// about the gin.Context

// 获取 url 参数
func GetURLParam(c *gin.Context, key string) []string {
	return c.Request.URL.Query()[key]
}

// 发生错误时的渲染
func ViewError(c *gin.Context, err error) {
	a, _ := c.Get(config.AccountKey)
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"Account": a,
		"Message": err.Error(),
	})
}

// 将html片段完整输出并要求解析
func UnEscape(s string) template.HTML {
	return template.HTML(s)
}

// 携带用户信息进行渲染
func ViewWithAccount(c *gin.Context, viewFile string, dataMap map[string]interface{}) {
	a, _ := c.Get(config.AccountKey)
	dataMap[config.AccountKey] = a
	c.HTML(http.StatusOK, viewFile, dataMap)
}
