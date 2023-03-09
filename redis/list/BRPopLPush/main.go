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
	BRPopLPush()
}

func BRPopLPush() { // 将一个元素从一个列表移动到另一个列表，并返回被移动的元素
	str := conn.BRPopLPush(ctx,"list2", "list", 1*time.Second).Val()
	fmt.Println("str = ", str)
}
