package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HExists()
}

func HExists() {
	b := conn.HExists(ctx,"hash-key2", "short").Val()
	fmt.Println("b = ", b)
}
