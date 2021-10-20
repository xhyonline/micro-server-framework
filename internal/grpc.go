package internal

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/xhyonline/xutil/micro"

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
	l, err := net.Listen("tcp", internalAddress())
	if err != nil {
		Logger.Errorf("监听失败 %s", err)
		return nil
	}
	addr := l.Addr().(*net.TCPAddr)
	port := addr.Port
	ip := addr.IP.String()
	s := grpc.NewServer()
	g := &grpcInstance{Server: s, listener: l}
	golang.RegisterRunnerServer(g.Server, &rpc.Service{})
	g.Run()
	ctx := sig.Get().RegisterClose(g)

	// 服务注册
	if err := micro.NewMicroServiceRegister(Instance.ETCD, configs.Instance.ETCD.Prefix, 10).
		Register(configs.Name, &micro.Node{
			Host: ip,
			Port: strconv.Itoa(port),
		}); err != nil {
		Logger.Errorf("服务注册失败 %s", err)
		return nil
	}

	Logger.Info("服务"+configs.Name, "已启动,启动地址:"+fmt.Sprintf("%s:%d", ip, port))
	return ctx.Done()
}

// internalAddress 获取服务地址
func internalAddress() string {
	addr, err := helper.IntranetAddress()
	if err != nil {
		Logger.Errorf("获取内网地址失败 %s", err)
		return ""
	}
	v, ok := addr["eth0"]
	if !ok {
		Logger.Errorf("未发现内网网卡 eth0")
		Logger.Errorf("网卡信息 %+v", addr)
		return ""
	}
	var ip net.IP
	for _, item := range v {
		if ip = item.To4(); ip != nil {
			break
		}
	}
	if ip.String() == "" {
		Logger.Errorf("未发现 IPv4 地址")
		return ""
	}
	address := ip.String() + ":0"
	address = "127.0.0.1:0"
	return address
}
