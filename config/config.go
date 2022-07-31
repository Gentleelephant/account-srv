package config

import (
	"fmt"
	"github.com/Gentleelephant/common/consts"
	"github.com/Gentleelephant/common/utils"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 本地配置
var (
	LocalConfig  *viper.Viper
	RemoteConfig *viper.Viper
	DB           *gorm.DB
	FilePath     string // 配置文件路径，命令启动时指定
	NacosConfig  *utils.NacosConfigparams
)

func GetRemoteConfig() {
	localConfig, err := utils.GetConfig(FilePath, consts.ConfigFileType)
	if err != nil {
		panic(err)
	}
	LocalConfig = localConfig
	LocalConfig.SetConfigFile(FilePath)
	LocalConfig.SetConfigType("yaml")
	err = LocalConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("read local config:", LocalConfig.AllSettings())
	configParams := utils.NacosConfigparams{
		DataId: LocalConfig.GetString(consts.NacosDataId),
		Group:  LocalConfig.GetString(consts.NacosGroup),
		ClientConfig: constant.ClientConfig{
			TimeoutMs:    LocalConfig.GetUint64(consts.NacosTimeoutMs),
			BeatInterval: LocalConfig.GetInt64(consts.NacosBeatInterval),
			NamespaceId:  LocalConfig.GetString(consts.NacosNamespaceId),
			CacheDir:     LocalConfig.GetString(consts.NacosCacheDir),
			LogDir:       LocalConfig.GetString(consts.NacosLogDir),
			LogLevel:     LocalConfig.GetString(consts.NacosLogLevel),
			ContextPath:  LocalConfig.GetString(consts.NacosContextPath),
		},
		ServerConfig: constant.ServerConfig{
			Scheme:      LocalConfig.GetString(consts.NacosScheme),
			ContextPath: LocalConfig.GetString(consts.NacosContextPath),
			IpAddr:      LocalConfig.GetString(consts.NacosIpAddr),
			Port:        LocalConfig.GetUint64(consts.NacosPort),
		},
	}
	NacosConfig = &configParams
	fmt.Println("read nacos config:", configParams)
	RemoteConfig, err = utils.InitRemoteConfig(configParams)
	if err != nil {
		panic(err)
	}
}

func GetDB() {
	sqlConfig := utils.SQLConfig{
		Host:     RemoteConfig.GetString(consts.SqlHost),
		Port:     RemoteConfig.GetInt(consts.SqlPort),
		Username: RemoteConfig.GetString(consts.SqlUsername),
		Password: RemoteConfig.GetString(consts.SqlPassword),
		Database: RemoteConfig.GetString(consts.SqlDatabase),
	}
	db, err := utils.InitDB(sqlConfig)
	if err != nil {
		panic(err)
	}
	DB = db
}
