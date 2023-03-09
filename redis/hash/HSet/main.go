package main

import (
	"context"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	HSet()
}

func HSet() {
	conn.HSet(ctx, "d-7", "field", 5)
	conn.HSet(ctx, "d-15", "field", 1)
	conn.HSet(ctx, "d-23", "field", 9)
	conn.HSet(ctx, "d-110", "field", 3)
}
