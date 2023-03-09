package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZCard()
}

func ZCard() {
	res := conn.ZCard(ctx,"zset-key").Val() // 有序集合成员数量
	fmt.Println("res = ", res)
}
