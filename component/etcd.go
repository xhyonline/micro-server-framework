package component

import (
	"fmt"
	"os"

	"github.com/xhyonline/micro-server-framework/configs"

	"github.com/xhyonline/xutil/etcd"
)

func RegisterETCD() Option {
	return func() {
		address := fmt.Sprintf("%s:%d", configs.Instance.ETCD.Host,
			configs.Instance.ETCD.Port)
		client, err := etcd.New(address)
		if err != nil {
			Logger.Errorf("etcd 启动失败,地址:%s %s", address, err)
			os.Exit(1)
		}
		Instance.ETCD = client
	}
}
