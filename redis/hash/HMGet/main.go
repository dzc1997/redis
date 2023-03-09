package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	HMGet()
}

func HMGet() {
	res := conn.HMGet(ctx,"hash-key", "k2", "k3").Val()
	for _, v := range res {
		fmt.Println("v = ", v)
	}
}
