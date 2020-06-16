package util

import "github.com/jinzhu/gorm"


const (
	userName   = "root"
	password   = "root"
	dbName     = "gopetstore"
	driverName = "mysql"
	charset    = "charset=utf8"
	local      = "loc=Local"
	tcpPort    = "@tcp(localhost:3306)/"
	parseTime  = "parseTime=true" // 用以解析 数据库 中的 date 类型，否则会解析成 []uint8 不能隐式转为 string
	dataSourceName = userName + ":" + password + tcpPort + dbName + "?" + charset + "&" + local + "&" + parseTime
)

func GetConnection() (*gorm.DB,error)  {
	return gorm.Open(driverName,dataSourceName)
}
