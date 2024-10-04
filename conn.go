package main

import "github.com/redis/go-redis/v9"

func GetClient() *redis.Client {
	// Client
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		DB:       0,
	})
}
