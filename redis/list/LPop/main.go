package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
	"sync"
)

var (
	conn = redisConn.ConnectRedis()
	ctx  = context.Background()
	wg   sync.WaitGroup
)

func main() {
	wg.Add(10)
	go LPop()
	wg.Wait()
}

func LPop() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		str := conn.LPop(ctx, "list").Val()
		fmt.Printf("str %d = %s\n", i, str)
		str2 := conn.LPop(ctx, "list2").Val()
		fmt.Printf("str2 %d = %s\n", i, str2)
	}
}
