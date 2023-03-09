package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	SAdd()
}

func SAdd() {
	res := conn.SAdd(ctx,"skey1", "a", "b", "c").Val()
	fmt.Println("res= ", res)
}
