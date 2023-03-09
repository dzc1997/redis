package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	SRem()
}

func SRem() {
	res := conn.SRem(ctx,"skey2", "a", "b", "c", "d").Val()
	fmt.Println("The num of remove set = ", res)
}
