package config

const (
	ServiceDynamicPort = "service.dynamic.port"
	ConfigFileType     = "yaml"
)

// Nacos
const (
	NacosHost      = "nacos.host"
	NacosPort      = "nacos.port"
	NacosNamespace = "nacos.namespace"
	NacosLogDir    = "nacos.logDir"
	NacosCacheDir  = "nacos.cacheDir"
	NacosGroup     = "nacos.group"
	NacosDataId    = "nacos.dataId"
	NacosLogLevel  = "nacos.logLevel"
)

// Mysql
const (
	MysqlHost     = "mysql.host"
	MysqlPort     = "mysql.port"
	MysqlUsername = "mysql.username"
	MysqlPassword = "mysql.password"
	MysqlDatabase = "mysql.database"
)

// Redis
const (
	RedisHost     = "redis.host"
	RedisPort     = "redis.port"
	RedisPassword = "redis.password"
	RedisDB       = "redis.db"
)
