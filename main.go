package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	// Client
	client := GetClient()

	ctx := context.Background()

	f := []any{
		"name", "Fulano",
		"age", 30,
		"document", "123.456.789-00",
		"created_at", time.Now(),
	}

	// HSET
	hset, err := client.HSet(ctx, "person:1", f).Result()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("HSET result:\t%#v\n", hset)

	// HGET
	hget, err := client.HGet(ctx, "person:1", "name").Result()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("HGET result:\t%#v\n", hget)

	// HGETALL
	hgetall, err := client.HGetAll(ctx, "person:1").Result()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("HGETALL result:\t%#v\n", hgetall)

	// HMGET
	hmget, err := client.HMGet(ctx, "person:1", "name", "document", "age", "created_at").Result()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("HMGET result:\t%#v\n", hmget)

	// HDEL
	hdel, err := client.HDel(ctx, "person:1", "name").Result()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("HDEL result:\t%#v\n", hdel)

	// HEXISTS
	hexists, err := client.HExists(ctx, "person:1", "age").Result()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("HEXISTS result:\t%#v\n", hexists)
}
