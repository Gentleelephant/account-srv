package internal

import (
	"fmt"
	"github.com/Gentleelephant/account-srv/config"
	"github.com/Gentleelephant/account-srv/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	vr := config.RemoteConfig
	host := vr.GetString(config.MysqlHost)
	port := vr.GetInt(config.MysqlPort)
	user := vr.GetString(config.MysqlUsername)
	password := vr.GetString(config.MysqlPassword)
	database := vr.GetString(config.MysqlDatabase)
	// 从配置中读取数据库配置
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	zap.L().Info("Connecting to database", zap.String("host", host), zap.Int("port", port), zap.String("database", database))
	open, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 是否使用单数表名
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = open.AutoMigrate(&model.Account{})
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	DB = open
}
