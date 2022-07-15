package internal

import (
	"account-srv/config"
	"account-srv/util"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"strings"
)

func InitConfig() {
	initRemoteConfig()
	initRemoteConfig()
}

// 初始化本地配置
func initLocalConfig() {
	vl := config.LocalConfig
	vl.SetConfigFile(config.ConfigFilePath)
	vl.SetConfigType("yaml")
	err := vl.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 初始化端口
	port := util.GetTCPPort()
	vl.Set(config.ServicePort, port)
}

// 初始化远程配置
func initRemoteConfig() {
	initLocalConfig()
	vl := config.LocalConfig
	vr := config.RemoteConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         vl.GetString("nacos.namespace"), //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              vl.GetString("nacos.logDir"),
		CacheDir:            vl.GetString("nacos.cacheDir"),
		LogLevel:            vl.GetString("logger.level"),
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      vl.GetString("nacos.host"),
			ContextPath: "/nacos",
			Port:        uint64(vl.GetInt("nacos.port")),
			Scheme:      "http",
		},
	}
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		panic(err)
	}
	// 获取配置
	remoteConfig, err := client.GetConfig(vo.ConfigParam{
		DataId: vl.GetString("nacos.dataId"),
		Group:  vl.GetString("nacos.group"),
		OnChange: func(namespace, group, dataId, data string) {
			// TODO 刷新配置
			// db
			// redis
			zap.L().Info("监听到配置改变")
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(remoteConfig)
	// 解析配置
	reader := strings.NewReader(remoteConfig)
	vr.SetConfigType("yaml")
	err = vr.ReadConfig(reader)
	if err != nil {
		return
	}
}
