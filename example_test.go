package main

import (
	"github.com/go-redis/redis/v9"
	"testing"
)

func BenchmarkRedisPing(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6333",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	for i := 0; i < b.N; i++ {
		RedisPing(rdb)
	}
}
