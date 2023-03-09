package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	SUnion()
}

func SUnion() {
	set := conn.SUnion(ctx, "skey1", "skey2").Val()
	for k, v := range set {
		fmt.Printf("v%d = %s\n", k, v)
	}
}
