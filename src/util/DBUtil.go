package util

import (
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jmoiron/sqlx"
)

const (
	userName       = "root"
	password       = "root"
	dbName         = "gopetstore"
	driverName     = "mysql"
	charset        = "charset=utf8"
	local          = "loc=Local"
	tcpPort        = "@tcp(localhost:3306)/"
	parseTime      = "parseTime=true" // 用以解析 数据库 中的 date 类型，否则会解析成 []uint8 不能隐式转为 string
	dataSourceName = userName + ":" + password + tcpPort + dbName + "?" + charset + "&" + local + "&" + parseTime
)

func GetConnection() (*sqlx.DB, error) {
	return sqlx.Connect(driverName, dataSourceName)
}

// 事务：函数式编程 sqlx 事务
func ExecTransaction(callback func(tx *sqlx.Tx) error) error {
	d, err := GetConnection()
	defer func() {
		if d != nil {
			_ = d.Close()
		}
	}()
	if err != nil {
		return err
	}
	tx, err := d.Beginx()

	if err != nil {
		return err
	}
	err = callback(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
