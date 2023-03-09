package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

type Client struct {
	Conn *redis.Client
}

func main() {
	var c *Client
	for i := 0; i < 3; i++ {
		go c.NotRans()
	}
	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 3; i++ {
		go c.Trans()
	}
	time.Sleep(500 * time.Millisecond)
	defer c.Reset()
}

func (r *Client) Reset() {
	r.Conn.FlushDB(ctx)
}

func (r *Client) NotRans() {
	fmt.Println(r.Conn.Incr(ctx, "notrans:").Val())
	time.Sleep(100 * time.Millisecond)
	fmt.Println(r.Conn.Decr(ctx, "notrans:").Val())
}

func (r *Client) Trans() {
	pipeline := r.Conn.Pipeline()
	pipeline.Incr(ctx, "trans:")
	time.Sleep(100 * time.Millisecond)
	pipeline.Decr(ctx, "trans:")
	_, err := pipeline.Exec(ctx)
	if err != nil {
		log.Println("pipeline failed,the err is:", err)
	}
}
