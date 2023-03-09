package main

import (
	"fmt"

	"context"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	LRange()
}

func LRange() {
	res := conn.LRange(ctx, "list", 0, -1).Val()
	for k, v := range res {
		fmt.Printf("res %d = %s\n", k, v)
	}

	res2 := conn.LRange(ctx, "list2", 0, -1).Val()
	for k, v := range res2 {
		fmt.Printf("res2 %d = %s\n", k, v)
	}
}
