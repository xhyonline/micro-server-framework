package component

import (
	"strconv"

	"github.com/xhyonline/micro-server-framework/configs"

	"github.com/xhyonline/xutil/db"
)

// RegisterMySQL 注册 MySQL 服务
func RegisterMySQL() Option {
	return func() {
		Instance.MySQL = db.NewDataBase(&db.Config{
			Host:          configs.Instance.MySQL.Host,
			Port:          strconv.Itoa(configs.Instance.MySQL.Port),
			User:          configs.Instance.MySQL.User,
			Password:      configs.Instance.MySQL.Password,
			Name:          configs.Instance.MySQL.DB,
			Lifetime:      3600,
			MaxActiveConn: configs.Instance.MySQL.MaxConnNum,
			MaxIdleConn:   configs.Instance.MySQL.IdleConnNum,
		})
	}
}
