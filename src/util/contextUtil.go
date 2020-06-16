package util

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// about the gin.Context

func GetURLParam(c *gin.Context, key string) []string {
	return c.Request.URL.Query()[key]
}

func ViewError(c *gin.Context, err error) {
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"Message": err.Error(),
	})
}

// 将html片段完整输出并要求解析
func UnEscape(s string) template.HTML {
	return template.HTML(s)
}
