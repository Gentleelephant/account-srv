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

	port := "9901"
	host := "127.0.0.1"

	// 初始化数据库
	internal.InitDB()
	// 初始化日志
	internal.InitLogger()

	// 启动grpc server
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
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
		fmt.Printf("listen on %s:%s", host, port)
	}()
	w.Wait()
	return nil
}
