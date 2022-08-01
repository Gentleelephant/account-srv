package server

import (
	"fmt"
	"github.com/Gentleelephant/account-srv/biz"
	"github.com/Gentleelephant/account-srv/config"
	"github.com/Gentleelephant/account-srv/internal"
	"github.com/Gentleelephant/common/utils"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start() {

	port, err := utils.GetTCPPort()
	if err != nil {
		panic(err)
	}
	ip, err := utils.GetOutBoundIP()
	if err != nil {
		addr, err := utils.GetInterfaceIpv4Addr("eth0")
		if err != nil {
			panic(err)
		}
		ip = addr
	}
	host := ip

	// 初始化配置
	config.GetRemoteConfig()
	// 初始化数据库
	config.GetDB()
	// 初始化日志
	internal.InitLogger()

	// 启动grpc server
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	zap.L().Info("grpc server listen on", zap.String("host", host), zap.Int("port", port))
	if err != nil {
		zap.L().Error("Start grpc server", zap.Error(err))
	}
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	go func() {
		err = server.Serve(listen)
		if err != nil {
			zap.L().Error("Server Start Error", zap.Error(err))
		}
		zap.L().Info("Server Start", zap.String("host", host), zap.Int("port", port))
	}()

	// 服务注册
	param := vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ClusterName: config.LocalConfig.GetString("service.cluster"),
		ServiceName: config.LocalConfig.GetString("service.name"),
		GroupName:   config.LocalConfig.GetString("service.group"),
		Ephemeral:   true,
	}
	_, err = utils.RegisterInstance(*config.NacosConfig, param)
	if err != nil {
		zap.L().Error("RegisterInstance Error", zap.Error(err))
	}
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	s := <-sig
	// 注销服务
	deparams := vo.DeregisterInstanceParam{
		Ip:          host,
		Port:        uint64(port),
		Cluster:     config.LocalConfig.GetString("service.cluster"),
		ServiceName: config.LocalConfig.GetString("service.name"),
		GroupName:   config.LocalConfig.GetString("service.group"),
		Ephemeral:   true,
	}
	_, err = utils.DeregisterInstance(*config.NacosConfig, deparams)
	if err != nil {
		zap.L().Error("DeregisterInstance Error", zap.Error(err))
	}
	zap.L().Info("Server deregister", zap.String("host", host),
		zap.Int("port", port),
		zap.String("cluster", "DEFAULT"),
		zap.String("service name", "account-srv"),
		zap.String("group name", "DEFAULT_GROUP"),
		zap.Bool("ephemeral", true))
	zap.L().Info("Receive Signal", zap.String("signal", s.String()))
}
