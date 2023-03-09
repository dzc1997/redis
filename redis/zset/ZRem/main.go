package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZRem()
}

func ZRem() {
	res := conn.ZRem(ctx,"zset-key", "c").Val() // 移除给定的成员,返回被移除成员的数量
	fmt.Println("res = ", res)
}
