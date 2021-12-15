package core

import (
	"dance/conf"
	"time"

	//"commercial/conf"
	//"commercial/cons"
	//"strings"
	//"time"
	"github.com/go-redis/redis"
	"strings"
)

var redisDB *redis.ClusterClient

func InitRedis() {
	redisDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        strings.Split(conf.Config.Redis, ","),
		DialTimeout:  time.Millisecond * 100,
		ReadTimeout:  time.Millisecond * 300,
		WriteTimeout: time.Millisecond * 300,
		Password:     conf.Config.RedisPassword,
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func GetRedis() *redis.ClusterClient {
	return redisDB
}
