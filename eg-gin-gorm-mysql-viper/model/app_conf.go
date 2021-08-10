package model

type AppConf struct {
	Env   string
	Port  int
	Mysql MysqlConfig
}

type MysqlConfig struct {
	Username        string
	Password        string
	Url             string
	Database        string
	MaxOpenconns    int
	MaxIdleConns    int
	ConnMaxLifetime int64
}
