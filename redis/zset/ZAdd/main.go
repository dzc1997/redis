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
	ZAdd()
}

func ZAdd() {
	res := conn.ZAdd(ctx,"zset-1", redis.Z{Member: "a", Score: 1}, redis.Z{Member: "b", Score: 2}, redis.Z{Member: "c", Score: 3}).Val()
	fmt.Println("res = ", res)

	res2 := conn.ZAdd(ctx,"zset-2", redis.Z{Member: "b", Score: 4}, redis.Z{Member: "d", Score: 0}, redis.Z{Member: "c", Score: 1}).Val()
	fmt.Println("res = ", res2)
}
