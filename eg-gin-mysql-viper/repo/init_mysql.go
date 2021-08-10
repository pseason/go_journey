package repo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"springmars.com/gin-mysql-viper/config"
)

// 定义一个全局对象db
var MysqlDb *sql.DB

// 初始化mysql连接
func InitDb() {
	var err error

	MysqlDb, err = sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8",
			config.AppConf.Mysql.Username,
			config.AppConf.Mysql.Password,
			config.AppConf.Mysql.Url,
			config.AppConf.Mysql.Database,
		))
	if err != nil {
		panic("mysql连接配置错误:" + err.Error())
	}
	// 最大连接数
	MysqlDb.SetMaxOpenConns(config.AppConf.Mysql.MaxOpenconns)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(config.AppConf.Mysql.MaxIdleConns)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(time.Duration(config.AppConf.Mysql.ConnMaxLifetime))

	if err = MysqlDb.Ping(); err != nil {
		panic("mysql连接异常:" + err.Error())
	}
}
