package util

import (
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jmoiron/sqlx"
	"gopetstore_v2/src/global"
)

func GetConnection() (*sqlx.DB, error) {
	return sqlx.Connect(global.DatabaseSetting.DriverName, global.DatabaseSetting.DataSourceName)
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
