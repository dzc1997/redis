package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	SMembers()
}

func SMembers() {
	res := conn.SMembers(ctx,"skey1").Val()
	for k, v := range res {
		fmt.Printf("setVal%d = %s\n", k, v)
	}
}
