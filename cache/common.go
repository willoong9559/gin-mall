package cache

import (
	"strconv"

	"github.com/go-redis/redis"
	"github.com/willoong9559/gin-mall/global"
	"gopkg.in/ini.v1"
)

// RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		global.Logger.Info("配置文件读取错误，请检查文件路径:", err)
	}
	LoadRedisData(file)
	Redis()
}

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		// Password: conf.RedisPw,
		DB: int(db),
	})
	// redis状态心跳检测
	_, err := client.Ping().Result()
	if err != nil {
		global.Logger.Panic(err)
	}
	RedisClient = client
}

func LoadRedisData(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
