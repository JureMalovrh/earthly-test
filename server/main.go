package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()
	c := rdb.Set(ctx, "test", "hello", -1)
	_ = c

	x := rdb.Get(ctx, "test")
	fmt.Println(x)
}