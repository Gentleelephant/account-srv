package config

import "github.com/spf13/viper"

// 配置文件路径，命令启动时指定
var ConfigFilePath string

// 本地配置
var LocalConfig = viper.New()

// 远程配置
var RemoteConfig = viper.New()
