package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	RPush()
}

func RPush() {
	res := conn.RPush(ctx, "sort-input", 1, 9, 3, 7, 5).Val()
	fmt.Println("res = ", res)
}
