package redisConn

import (
	"context"
	"gitee.com/wedone/redis_course/redis/config"
	"github.com/go-redis/redis/v9"
	"log"
)

func ConnectRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	if _, err := conn.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Connect to redis client failed,err:%v\n", err)
	}
	return conn
}
