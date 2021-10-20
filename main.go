package main

import (
	"github.com/xhyonline/micro-server-framework/configs"
	"github.com/xhyonline/micro-server-framework/internal"

	// nolint
	. "github.com/xhyonline/micro-server-framework/component" // 忽略包名
)

func main() {
	// 初始化配置
	configs.Init(configs.WithBaseConfig(),
		configs.WithRedis(),
		configs.WithMySQL(), configs.WithETCD())
	// 初始化组件
	Init(RegisterLogger(), RegisterETCD())
	// 启动 grpc
	<-internal.Run()
}
