package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HIncry()
}

func HIncry() {
	res := conn.HIncrBy(ctx,"hash-key2", "num", 1).Val()
	fmt.Println("res = ", res)
}
