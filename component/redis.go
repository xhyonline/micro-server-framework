package component

import (
	"strconv"

	"github.com/xhyonline/micro-server-framework/configs"

	"github.com/xhyonline/xutil/kv"
)

// RegisterMySQL 注册 Redis 服务
func RegisterRedis() Option {
	return func() {
		Instance.Redis = kv.NewRedisClient(kv.Config{
			Host:         configs.Instance.Redis.Host,
			Port:         strconv.Itoa(configs.Instance.Redis.Port),
			Password:     configs.Instance.Redis.Password,
			DB:           configs.Instance.Redis.DB,
			PoolSize:     configs.Instance.Redis.MaxConnNum,
			MinIdleConns: configs.Instance.Redis.IdleConnNum,
		})
	}
}
