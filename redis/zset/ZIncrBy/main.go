package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZIncrBy()
}

func ZIncrBy() {
	res := conn.ZIncrBy(ctx,"zset-key", 3, "c").Val()
	fmt.Println("res = ", res)
}
