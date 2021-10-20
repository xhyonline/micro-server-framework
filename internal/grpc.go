package internal

import (
	"fmt"
	"net"
	"os"

	"github.com/xhyonline/micro-server-framework/gen/golang"
	"github.com/xhyonline/micro-server-framework/rpc"

	"github.com/xhyonline/micro-server-framework/configs"

	"github.com/xhyonline/xutil/helper"
	"github.com/xhyonline/xutil/sig"
	"google.golang.org/grpc"

	// nolint
	. "github.com/xhyonline/micro-server-framework/component" // 忽略包名
)

type grpcInstance struct {
	*grpc.Server
	listener net.Listener
}

// GracefulClose 优雅停止
func (s *grpcInstance) GracefulClose() {
	Logger.Info("服务" + configs.Name + "接收到关闭通知")
	s.GracefulStop()
	Logger.Info("服务" + configs.Name + "已优雅停止")
}

// Run 启动
func (s *grpcInstance) Run() {
	go func() {
		if err := s.Serve(s.listener); err != nil {
			Logger.Errorf("服务 %s 启动失败 %s", configs.Name, err)
			os.Exit(1)
		}
	}()
}

func Run() <-chan struct{} {
	addr, err := helper.IntranetAddress()
	if err != nil {
		Logger.Errorf("获取内网地址失败 %s", err)
		os.Exit(1)
	}
	fmt.Printf("调试 %+v", addr)
	os.Exit(1)
	v, ok := addr["eth0"]
	if !ok {
		Logger.Errorf("未发现内网网卡 eth0 %s", err)
		Logger.Errorf("网卡信息 %+v", addr)
		os.Exit(1)
	}
	os.Exit(1)
	l, err := net.Listen("tcp", v.To4().String()+":0")
	if err != nil {
		Logger.Errorf("监听失败 %s", err)
	}
	s := grpc.NewServer()
	g := &grpcInstance{Server: s, listener: l}
	golang.RegisterRunnerServer(g.Server, &rpc.Service{})
	g.Run()
	ctx := sig.Get().RegisterClose(g)
	return ctx.Done()

}
