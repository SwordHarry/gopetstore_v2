package main

import (
	"github.com/gin-gonic/gin"
	"gopetstore_v2/src/config"
	"gopetstore_v2/src/global"
	"gopetstore_v2/src/route"
	"gopetstore_v2/src/util"
	"html/template"
	"log"
	"path/filepath"
)

func main() {
	r := gin.Default()
	setFrontConfig(r)
	// 注册路由
	route.RegisterRoute(r)
	err := r.Run(":" + global.ServerSetting.Port)
	if err != nil {
		panic(err)
	}
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatal(err)
	}
}

func setFrontConfig(r *gin.Engine) {
	// 注册自定义函数
	r.SetFuncMap(template.FuncMap{
		"unEscape": util.UnEscape,
	})
	// 设置静态文件加载
	for k, v := range global.StaticConfig {
		r.Static(k, v)
	}
	r.LoadHTMLGlob(filepath.Join("front", "web", "**", "*"))
}

func setupSetting() error {
	newSetting, err := config.NewSetting()
	if err != nil {
		return err
	}
	if err := newSetting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	global.DatabaseSetting.DataSourceName = global.DatabaseSetting.UserName +
		":" + global.DatabaseSetting.Password + "@tcp(" +
		global.DatabaseSetting.Domain + ":" +
		global.DatabaseSetting.Port + ")/" +
		global.DatabaseSetting.DBName + "?" +
		"charset=" + global.DatabaseSetting.Charset + "&" +
		"loc=" + global.DatabaseSetting.Local + "&" +
		"parseTime=" + global.DatabaseSetting.ParseTime
	if err := newSetting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	return nil
}
