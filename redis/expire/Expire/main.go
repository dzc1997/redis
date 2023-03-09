package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
	"time"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	Expire()
}

func Expire() {
	conn.Set(ctx, "key", "value", 0)
	res := conn.Get(ctx, "key").Val()
	fmt.Println("res = ", res)
	conn.Expire(ctx, "key", 1*time.Second)
	time.Sleep(2 * time.Second)
	conn.Get(ctx, "key").Val()
	conn.Set(ctx, "key", "value2", 0)
	conn.Expire(ctx, "key", 100*time.Second)
}
