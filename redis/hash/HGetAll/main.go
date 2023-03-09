package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HGetAll()
}

func HGetAll() {
	res := conn.HGetAll(ctx,"hash-key").Val()
	fmt.Println("res = ", res)
}
