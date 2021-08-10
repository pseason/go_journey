package repo

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var MysqlDb *sql.DB

// 初始化mysql连接
func InitDb() {
	var err error
	MysqlDb, err = sql.Open("mysql", "root:game@tcp(192.168.50.183:3306)/test?charset=utf8")
	if err != nil {
		panic("mysql连接配置错误:" + err.Error())
	}
	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(6 * time.Second)

	if err = MysqlDb.Ping(); err != nil {
		panic("mysql连接异常:" + err.Error())
	}
}
