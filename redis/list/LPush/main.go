package main

import (
	"context"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()

var ctx  = context.Background()

func main() {
	LPush()
}

func LPush() {
	conn.LPush(ctx,"list-key", "first")
}
