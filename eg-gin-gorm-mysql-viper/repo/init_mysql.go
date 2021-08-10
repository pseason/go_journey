package repo

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"springmars.com/gin-gorm-mysql-viper/config"
)

// 定义一个全局对象db
var MysqlDb *gorm.DB

// 初始化mysql连接
func InitDb() {
	var err error
	MysqlDb, err = gorm.Open(
		mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.AppConf.Mysql.Username, config.AppConf.Mysql.Password,
			config.AppConf.Mysql.Url, config.AppConf.Mysql.Database,
		)), &gorm.Config{})

	if err != nil {
		panic("mysql连接配置错误:" + err.Error())
	}
	sqlDB, e := MysqlDb.DB()
	if e != nil {
		panic("mysql连接配置错误:" + err.Error())
	}
	// 最大连接数
	sqlDB.SetMaxOpenConns(config.AppConf.Mysql.MaxOpenconns)
	// 闲置连接数
	sqlDB.SetMaxIdleConns(config.AppConf.Mysql.MaxIdleConns)
	// 最大连接周期
	sqlDB.SetConnMaxLifetime(time.Duration(config.AppConf.Mysql.ConnMaxLifetime))

	if err = sqlDB.Ping(); err != nil {
		panic("mysql连接异常:" + err.Error())
	}
}
