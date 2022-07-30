package internal

import (
	"fmt"
	"github.com/Gentleelephant/account-srv/config"
	"github.com/Gentleelephant/account-srv/util"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strings"
)

func InitConfig() {
	initLocalConfig()
	initRemoteConfig()
}

// 初始化本地配置
func initLocalConfig() {
	vl := config.LocalConfig
	vl.SetConfigFile(config.ConfigFilePath)
	vl.SetConfigType(config.ConfigFileType)
	err := vl.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// 初始化端口
	port := util.GetTCPPort()
	vl.Set(config.ServiceDynamicPort, port)
}

// 初始化远程配置
func initRemoteConfig() {
	initLocalConfig()
	vl := config.LocalConfig
	vr := config.RemoteConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         vl.GetString(config.NacosNamespace), //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              vl.GetString(config.NacosLogDir),
		CacheDir:            vl.GetString(config.NacosCacheDir),
		LogLevel:            vl.GetString(config.NacosLogLevel),
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      vl.GetString(config.NacosHost),
			ContextPath: "/nacos",
			Port:        uint64(vl.GetInt(config.NacosPort)),
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
		DataId: vl.GetString(config.NacosDataId),
		Group:  vl.GetString(config.NacosGroup),
	})
	if err != nil {
		panic(err)
	}
	// 解析配置
	parseConfig(vr, remoteConfig)
	err = client.ListenConfig(vo.ConfigParam{
		DataId: vl.GetString(config.NacosDataId),
		Group:  vl.GetString(config.NacosGroup),
		OnChange: func(namespace, group, dataId, data string) {
			// 配置变更
			zap.L().Info("config changed", zap.String("namespace", namespace), zap.String("group", group), zap.String("dataId", dataId), zap.String("data", data))
			// 刷新配置
			parseConfig(vr, data)
		},
	})
	if err != nil {
		zap.L().Error("listen config failed", zap.Error(err))
	}
	fmt.Println(remoteConfig)
}

func parseConfig(viper *viper.Viper, data string) {
	// 解析配置
	reader := strings.NewReader(data)
	viper.SetConfigType(config.ConfigFileType)
	err := viper.ReadConfig(reader)
	if err != nil {
		zap.L().Error("parse config failed", zap.Error(err))
		return
	}
}
