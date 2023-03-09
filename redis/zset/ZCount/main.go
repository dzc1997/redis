package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZCount()
}

func ZCount() {
	res := conn.ZCount(ctx,"zset-key", "0", "3").Val() // (min,max)
	fmt.Println("res = ", res)
}
