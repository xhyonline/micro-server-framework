package internal

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/xhyonline/micro-server-framework/component"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/xhyonline/micro-server-framework/configs"
	"github.com/xhyonline/micro-server-framework/gen/golang"
	"github.com/xhyonline/micro-server-framework/rpc"
	"github.com/xhyonline/xutil/helper"
	"github.com/xhyonline/xutil/logger"
	"github.com/xhyonline/xutil/micro"
	"github.com/xhyonline/xutil/sig"
	"google.golang.org/grpc"
)

type grpcInstance struct {
	*grpc.Server
	listener net.Listener
}

// GracefulClose 优雅停止
func (s *grpcInstance) GracefulClose() {
	logger.Info("服务" + configs.Name + "接收到关闭通知")
	s.GracefulStop()
	logger.Info("服务" + configs.Name + "已优雅停止")
}

// Run 启动
func (s *grpcInstance) Run() {
	go func() {
		if err := s.Serve(s.listener); err != nil {
			logger.Errorf("服务 %s 启动失败 %s", configs.Name, err)
			os.Exit(1)
		}
	}()
}

func Run() <-chan struct{} {
	ip := internalIP()
	port := strconv.Itoa(configs.Instance.Base.Port)
	l, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		logger.Errorf("%s 监听失败 %s", ip+":"+port, err)
		return nil
	}
	g := &grpcInstance{Server: grpc.NewServer(registerMiddleware()...), listener: l}
	golang.RegisterRunnerServer(g.Server, &rpc.Service{})
	g.Run()
	ctx := sig.Get().RegisterClose(g)

	pprofMonitor()
	// 服务注册
	if err := micro.NewMicroServiceRegister(component.Instance.ETCD, configs.Instance.ETCD.Prefix, 10).
		Register(configs.Name, &micro.Node{
			Host: ip,
			Port: port,
		}); err != nil {
		logger.Errorf("服务注册失败 %s", err)
		return nil
	}

	logger.Info("服务"+configs.Name, "已启动,启动地址:"+fmt.Sprintf("%s:%s", ip, port))
	return ctx.Done()
}

// internalIP 获取内网 IP
func internalIP() string {
	var address = "127.0.0.1"
	addr, err := helper.IntranetAddress()
	if err != nil {
		logger.Errorf("获取内网地址失败,服务停止 %s", err)
		os.Exit(1)
	}
	v, _ := addr["eth0"]
	var ip net.IP
	for _, item := range v {
		if ip = item.To4(); ip != nil {
			break
		}
	}
	if ip != nil {
		address = ip.String()
	} else {
		logger.Errorf("未发现 IPv4 地址,将使用 %s 替代", address)
	}

	return address
}

// registerMiddleware 注册中间键
func registerMiddleware() []grpc.ServerOption {
	return []grpc.ServerOption{
		// 处理 panic
		grpcmiddleware.WithUnaryServerChain(
			grpcrecovery.UnaryServerInterceptor(RecoveryInterceptor()),
		),
		grpc.WriteBufferSize(0), grpc.ReadBufferSize(0),
	}
}
