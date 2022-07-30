package config

import (
	"fmt"
	"github.com/Gentleelephant/common/consts"
	"github.com/Gentleelephant/common/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 本地配置
var (
	LocalConfig  = viper.New()
	RemoteConfig *viper.Viper
	DB           *gorm.DB
	FilePath     string // 配置文件路径，命令启动时指定
)

func GetRemoteConfig() {
	LocalConfig.SetConfigName(FilePath)
	err := LocalConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("read local config:", LocalConfig.AllSettings())
	configParams := utils.NacosConfigparams{
		NacosNamespace: LocalConfig.GetString(consts.NacosNamespace),
		NacosHost:      LocalConfig.GetString(consts.NacosHost),
		NacosLogDir:    LocalConfig.GetString(consts.NacosLogDir),
		NacosCacheDir:  LocalConfig.GetString(consts.NacosCacheDir),
		NacosLogLevel:  LocalConfig.GetString(consts.NacosLogLevel),
		NacosPort:      LocalConfig.GetUint64(consts.NacosPort),
		NacosDataId:    LocalConfig.GetString(consts.NacosDataId),
		NacosGroup:     LocalConfig.GetString(consts.NacosGroup),
	}
	RemoteConfig, err = utils.InitRemoteConfig(configParams)
	if err != nil {
		panic(err)
	}
}

func GetDB() {

}
