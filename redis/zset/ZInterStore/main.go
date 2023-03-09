package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
	"github.com/go-redis/redis/v9"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZInterStore()
}

func ZInterStore() {
	res := conn.ZInterStore(ctx,"zset-i", &redis.ZStore{Keys: []string{"ZSet-1", "ZSet-2"}}).Val()
	fmt.Println("res = ", res)
}
