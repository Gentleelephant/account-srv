package internal

import (
	"account-srv/config"
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"os"
	"time"
)

const (
	ServiceName = "account-srv"
)

var (
	// 服务入口在 etcd 存储的 key 前缀
	serviceEndpointKeyPrefix = "/service_register"
	hostname                 string
	endpoint                 string
)

var etcdConfig = clientv3.Config{
	Endpoints:            []string{"127.0.0.1:2379"},
	DialTimeout:          30 * time.Second,
	DialKeepAliveTimeout: 30 * time.Second,
}

func ServiceRegister(ctx context.Context) {
	hostname, _ = os.Hostname()
	cli, err := clientv3.New(etcdConfig)
	if err != nil {
		panic(err)
	}
	address := fmt.Sprintf("%s:%d", hostname, config.LocalConfig.GetInt(config.ServicePort))
	key := fmt.Sprintf("%s/%s/%s", serviceEndpointKeyPrefix, ServiceName, address)
	zap.L().Debug("register service", zap.String("service", ServiceName), zap.String("address", address))
	// 过期时间: 3秒钟
	ttl := 3
	// 创建租约
	lease, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		Logger.Error(err.Error())
	}
	b, _ := json.Marshal(lease)
	zap.L().Info("grant lease success: ", zap.String("", string(b)))
	// put kv
	res, err := cli.Put(ctx, key, endpoint, clientv3.WithLease(lease.ID))
	if err != nil {
		zap.L().Error(err.Error())
	}
	b, _ = json.Marshal(res)
	zap.L().Info("put kv with lease success: ", zap.String("", string(b)))
	// 保持租约不过期
	klRes, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		panic(err)
	}
	// 监听续约情况
	for v := range klRes {
		b, _ = json.Marshal(v)
		zap.L().Info("keep lease alive success: ", zap.String("", string(b)))
	}
	zap.L().Info("keep lease alive end")
}
