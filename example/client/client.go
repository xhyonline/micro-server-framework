package main

import (
	"context"
	"fmt"
	"time"

	"github.com/xhyonline/micro-server-framework/component"
	"github.com/xhyonline/micro-server-framework/configs"
	"github.com/xhyonline/micro-server-framework/gen/golang"
	"github.com/xhyonline/micro-server-framework/gen/golang/basic"
	xgrpc "github.com/xhyonline/xutil/grpc"
	"github.com/xhyonline/xutil/logger"
)

func main() {
	// 初始化配置
	configs.Init(configs.WithBaseConfig(), configs.WithRedis(), configs.WithMySQL(), configs.WithETCD())
	// 初始化微服务组件
	component.Init(component.RegisterETCD())
	// 启动 grpc
	conn := xgrpc.NewGRPCClient("myapp", component.Instance.ETCD)
	client := golang.NewRunnerClient(conn)
	for {
		resp, err := client.Hello(context.Background(), &basic.Empty{})
		if err != nil {
			logger.Error("发生错误了" + err.Error())
			continue
		}
		fmt.Println(resp.GetData())
		time.Sleep(time.Second)
	}
}
