package route

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/global"
	"gopetstore_v2/src/util"
)

// 是否登录，从 session 中获取 并往 context 中存放 account 指针
func AccountLogin(c *gin.Context) {
	a := util.GetAccountFromSession(c.Request)
	c.Set(global.AccountKey, a)
	c.Next()
}
