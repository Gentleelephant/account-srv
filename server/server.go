package server

import (
	"account-srv/biz"
	"account-srv/internal"
	"fmt"
	pb "github.com/Gentleelephant/proto-center/pb/model"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func Start() error {

	w := sync.WaitGroup{}

	// 初始化数据库
	internal.InitDB()

	// 启动grpc server
	listen, err := net.Listen("tcp", ":9901")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	w.Add(1)
	go func() {
		defer w.Done()
		err := server.Serve(listen)
		if err != nil {
			fmt.Println("server error")
		}
	}()
	w.Wait()
	return nil
}
