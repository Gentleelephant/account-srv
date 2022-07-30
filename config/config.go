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
	LocalConfig.SetConfigFile(FilePath)
	LocalConfig.SetConfigType("yaml")
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
