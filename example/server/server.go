package main

import (
	"github.com/xhyonline/micro-server-framework/component"
	"github.com/xhyonline/micro-server-framework/configs"
	"github.com/xhyonline/micro-server-framework/gen/golang"
	"github.com/xhyonline/micro-server-framework/rpc"
	xgrpc "github.com/xhyonline/xutil/grpc"
	"google.golang.org/grpc"
)

func main() {
	// 初始化配置
	configs.Init(configs.WithETCD())
	// 初始化微服务组件
	component.Init(component.RegisterETCD())

	xgrpc.StartGRPCServer(func(server *grpc.Server) {
		golang.RegisterRunnerServer(server, &rpc.Service{})
		// 自定义配置 xgrpc.WithPort(8080) 、xgrpc.WithIP("0.0.0.0)
	}, xgrpc.WithAppName(configs.Name), xgrpc.WithETCD(component.Instance.ETCD))
}
