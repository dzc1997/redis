package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HMSet()
}

func HMSet() {
	b := conn.HMSet(ctx,"hash-key2", map[string]interface{}{
		"short": "hello",
		"long":  "1000",
	}).Val()
	fmt.Println("b = ", b)
}
