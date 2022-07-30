package server

import (
	"fmt"
	"github.com/Gentleelephant/account-srv/biz"
	"github.com/Gentleelephant/account-srv/internal"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func Start() error {

	w := sync.WaitGroup{}

	port := "9901"
	host := "127.0.0.1"

	// 初始化配置
	internal.InitConfig()
	// 初始化数据库
	internal.InitDB()
	// 初始化日志
	internal.InitLogger()

	// 启动grpc server
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	zap.L().Info("grpc server listen on", zap.String("host", host), zap.String("port", port))
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
		zap.L().Info("Server Start", zap.String("host", host), zap.String("port", port))
	}()
	w.Wait()
	return nil
}
