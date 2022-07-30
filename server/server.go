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
	"sync"
)

func Start() error {

	w := sync.WaitGroup{}

	port, err := utils.GetTCPPort()
	if err != nil {
		panic(err)
	}
	host := "127.0.0.1"

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
		return err
	}
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	w.Add(1)
	go func() {
		defer w.Done()
		err = server.Serve(listen)
		if err != nil {
			zap.L().Error("Server Start Error", zap.Error(err))
		}
		zap.L().Info("Server Start", zap.String("host", host), zap.Int("port", port))
	}()

	// 服务注册
	param := vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        uint64(port),
		Weight:      0,
		Enable:      true,
		Healthy:     true,
		Metadata:    nil,
		ClusterName: "cluster-account",
		ServiceName: "account-srv",
		GroupName:   "account",
		Ephemeral:   true,
	}
	//param := vo.DeregisterInstanceParam{
	//	Ip:          "localhost",
	//	Port:        51558,
	//	Cluster:     "cluster-account",
	//	ServiceName: "account-srv",
	//	GroupName:   "account",
	//	Ephemeral:   false,
	//}
	_, err = utils.RegisterInstance(*config.NacosConfig, param)
	if err != nil {
		zap.L().Error("RegisterInstance Error", zap.Error(err))
		return err
	}
	w.Wait()
	return nil
}
