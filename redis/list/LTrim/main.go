package main

import (
	"context"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var ctx = context.Background()
var conn = redisConn.ConnectRedis()

func main() {
	LTrim()
}

func LTrim() {
	conn.LTrim(ctx, "list-key", 2, -1)
}
