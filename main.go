package main

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/route"
	"gopetstore_v2/src/util"
	"html/template"
	"path/filepath"
)

const port = ":8082"

func main() {
	r := gin.Default()
	r.Use(route.AccountLogin)
	// 注册自定义函数
	r.SetFuncMap(template.FuncMap{
		"unEscape": util.UnEscape,
	})
	// 设置静态文件加载
	for k, v := range config.StaticConfig {
		r.Static(k, v)
	}
	r.LoadHTMLGlob(filepath.Join("front", "web", "**", "*"))
	// 注册路由
	route.RegisterRoute(r)
	err := r.Run(port)
	if err != nil {
		panic(err)
	}
}
