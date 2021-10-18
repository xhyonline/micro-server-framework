package main

import (
	"net"

	"github.com/xhyonline/micro-server-framework/configs"
	"github.com/xhyonline/micro-server-framework/gen/golang"
	"github.com/xhyonline/micro-server-framework/rpc"
	"google.golang.org/grpc"
)

func main() {
	// 初始化配置
	configs.Init(configs.WithRedis(), configs.WithMySQL())
	// 初始化 mysql 、redis 等服务组件
	//Init(RegisterRedis(), RegisterMySQL(), RegisterLogger())=
	g := grpc.NewServer()
	golang.RegisterRunnerServer(g, rpc.NewService())
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	panic(g.Serve(l))
}
