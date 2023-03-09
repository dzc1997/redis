package main

import (
	"context"
	"fmt"
	"gitee.com/wedone/redis_course/redis/redisConn"
)

var conn = redisConn.ConnectRedis()
var ctx = context.Background()

func main() {
	Append()
	Append2()
	GetRange()
	SetRange()
	SetRange2()
	SetRange3()
	Get()
	Del()
}

func Append() {
	res := conn.Append(ctx,"new-string", "hello,").Val() // 返回当前string长度
	fmt.Println("res = ", res)
}

func Append2() {
	res := conn.Append(ctx,"new-string", "world!").Val()
	fmt.Println("res = ", res)
}

func GetRange() { // 获取[4,6]字符
	str := conn.GetRange(ctx,"new-string", 4, 6).Val()
	fmt.Println("str = ", str)
}

func SetRange() { // setRange修改字符h->H
	res := conn.SetRange(ctx,"new-string", 0, "H").Val() // 返回当前字符串总长度
	fmt.Println("res = ", res)
}

func SetRange2() { // setRange修改字符
	res := conn.SetRange(ctx,"new-string", 6, "W").Val() // 返回当前字符串总长度,32offset偏移量
	fmt.Println("res = ", res)
}

func SetRange3() { // 追加string
	res := conn.SetRange(ctx,"new-string", 11, ",how are you?")
	fmt.Println("res = ", res)
}

func Get() {
	res := conn.Get(ctx,"new-string").Val()
	fmt.Println("res = ", res)
}

func Del() {
	res := conn.Del(ctx,"new-string")
	fmt.Println("res = ", res)
}
