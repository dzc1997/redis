package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx  = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	RPop()
}

func RPop() {
	str := conn.RPop(ctx,"list").Val()
	fmt.Println("str = ", str)
}
