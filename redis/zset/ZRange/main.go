package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZRangeWithScores()
}

func ZRangeWithScores() {
	res := conn.ZRangeWithScores(ctx,"zset-u", 0, -1).Val() // 移除给定的成员,返回被移除成员的数量
	for _, v := range res {
		fmt.Println("v=", v)
	}
}
