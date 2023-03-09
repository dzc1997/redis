package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	SMove()
}

func SMove() {
	b := conn.SMove(ctx,"set-key", "set-key2", "b").Val() // 从key->key2
	fmt.Println("b = ", b)
}
