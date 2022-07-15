package internal

import (
	"account-srv/config"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"testing"
)

func TestRegisterService(t *testing.T) {
	config.ConfigFilePath = "../config/config.yaml"
	initLocalConfig()
	InitLogger()
	ServiceRegister(context.Background())
}

func TestIp(t *testing.T) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		//if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		//	if ipnet.IP.To4() != nil {
		//		fmt.Println(ipnet.IP.String())
		//	}
		//
		//}
		if ipNet, isIpNet := address.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 := ipNet.IP.String()
				fmt.Println(ipv4)
				return
			}
		}
	}
}

func TestPort(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", addr)
	p := l.Addr().(*net.TCPAddr).Port
	fmt.Println(p)
	defer l.Close()

}

func TestExternalIp(t *testing.T) {

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func TestViper(t *testing.T) {
	config.ConfigFilePath = "../config/config.yaml"
	initLocalConfig()
}

func TestViperEtcd(t *testing.T) {
	config.ConfigFilePath = "../config/config.yaml"
	initRemoteConfig()
}
