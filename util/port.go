package util

import (
	"go.uber.org/zap"
	"net"
)

func GetTCPPort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		zap.L().Error("Func GetTCPPort", zap.Error(err))
	}
	l, err := net.ListenTCP("tcp", addr)
	defer l.Close()
	if err != nil {
		zap.L().Error("Func GetTCPPort", zap.Error(err))
	}
	port := l.Addr().(*net.TCPAddr).Port
	return port
}
