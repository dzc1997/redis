package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	HKeys()
}

func HKeys() {
	res := conn.HKeys(ctx,"hash-key2").Val()
	fmt.Println("res = ", res)
}
