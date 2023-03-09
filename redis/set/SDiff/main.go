package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	SDiff()
}

func SDiff() {
	set := conn.SDiff(ctx,"skey1", "skey2").Val() // key1=[a,b,c] key2=[b] => [a,c]
	for k, v := range set {
		fmt.Printf("set%d = %s\n", k, v)
	}
}
