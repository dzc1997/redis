package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	ZRank()
}

func ZRank() {
	res := conn.ZRank(ctx,"zset-key", "a").Val() // 排名
	fmt.Println("res-a = ", res)

	res2 := conn.ZRank(ctx,"zset-key", "b").Val() // 排名
	fmt.Println("res-b = ", res2)

	res3 := conn.ZRank(ctx,"zset-key", "c").Val() // 排名
	fmt.Println("res-c = ", res3)
}
