package config

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"springmars.com/http-mysql-viper/model"
)

var AppConf model.AppConf
var env string

const (
	port_str              = "%s.port"
	mysql_username        = "%s.mysql.username"
	mysql_password        = "%s.mysql.password"
	mysql_url             = "%s.mysql.url"
	mysql_database        = "%s.mysql.database"
	mysql_maxOpenconns    = "%s.mysql.maxOpenconns"
	mysql_maxIdleConns    = "%s.mysql.maxIdleConns"
	mysql_connMaxLifetime = "%s.mysql.connMaxLifetime"
)

func LoadConfig() {
	// 标准库Flag
	flag.String("env", "dev", "--env=dev 程序运行读取配置环境")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	//解析flag viper绑定pflag
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	// 读取程序运行读取配置环境
	env = viper.Get("env").(string)

	viper.SetConfigName("app")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("read config failed:" + err.Error())
	}
	parse(&AppConf)
	// watch
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//viper配置发生变化了 执行响应的操作
		fmt.Println("Config file changed:", e.Name)
		parse(&AppConf)
	})
}

// 解析
func parse(conf *model.AppConf) {
	conf.Env = env
	conf.Port = int(viper.Get(fmt.Sprintf(port_str, env)).(int64))
	conf.Mysql.Username = viper.Get(fmt.Sprintf(mysql_username, env)).(string)
	conf.Mysql.Password = viper.Get(fmt.Sprintf(mysql_password, env)).(string)
	conf.Mysql.Url = viper.Get(fmt.Sprintf(mysql_url, env)).(string)
	conf.Mysql.Database = viper.Get(fmt.Sprintf(mysql_database, env)).(string)
	conf.Mysql.MaxOpenconns = int(viper.Get(fmt.Sprintf(mysql_maxOpenconns, env)).(int64))
	conf.Mysql.MaxIdleConns = int(viper.Get(fmt.Sprintf(mysql_maxIdleConns, env)).(int64))
	conf.Mysql.ConnMaxLifetime = viper.Get(fmt.Sprintf(mysql_connMaxLifetime, env)).(int64)
	fmt.Printf("AppConf: %+v \n", AppConf)
}
