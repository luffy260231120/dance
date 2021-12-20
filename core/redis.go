package core

import (
	"dance/conf"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func InitRedis() {
	redis_opt := redis.Options{
		Addr:     conf.Config.Redis,
		Password: conf.Config.RedisPassword,
	}
	// 创建连接池
	redisDB = redis.NewClient(&redis_opt)
	// 判断是否能够链接到数据库
	err := redisDB.Ping().Err()
	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.Client {
	return redisDB
}
