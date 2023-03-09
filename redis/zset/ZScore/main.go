package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZScore()
}

func ZScore() {
	res := conn.ZScore(ctx,"zset-key", "c").Val()
	fmt.Println("res = ", res)
}
