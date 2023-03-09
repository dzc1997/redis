package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HDel()
}

func HDel() {
	res := conn.HDel(ctx,"hash-key", "k2", "k3").Val()
	fmt.Println("The len of the hash = ", res)
}
