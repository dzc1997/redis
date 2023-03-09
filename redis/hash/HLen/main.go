package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HLen()
}

func HLen() {
	res := conn.HLen(ctx,"hash-key").Val()
	fmt.Println("The len of the hash = ", res)
}
