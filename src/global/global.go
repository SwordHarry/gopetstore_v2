package global

import "gopetstore_v2/src/config"

var StaticConfig = map[string]string{
	"/css":    "front/css/",
	"/images": "front/images/",
}

const (
	AccountKey = "Account"
	ProductKey = "Product"
	CartKey    = "Cart"
	OrderKey   = "Order"
)

// 配置文件 setting
var (
	DatabaseSetting *config.DatabaseSettingS
	ServerSetting   *config.ServerSettingS
)
