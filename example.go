package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/wfan99/log"
)

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6333",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set(ctx, "key", 123, 0).Err()
	if err != nil {
		log.Error("err != nil", log.ErrObj(err))
		return
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Error("err != nil", log.ErrObj(err))
		return
	}
	log.Debug("key", log.String("value", val))

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		log.Debug("key2 does not exist")
	} else if err != nil {
		log.Error("err != nil", log.ErrObj(err))
	} else {
		log.Debug("key2", log.String("value", val2))
	}
	// Output: key value
	// key2 does not exist
}

func RedisPing(rdb *redis.Client) {
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Error("err != nil", log.ErrObj(err))
		return
	}
}

func Tutorial() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6333",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Set无论key是否存在设置值
	val, err := rdb.Set(ctx, "key", "value", 0).Result()
	if err != nil {
		log.Error(" err != nil", log.ErrObj(err))
		return
	}
	log.Debug("Set", log.String("val", val))
	// SetNX
	bVal, err := rdb.SetNX(ctx, "key", "nxValue", 0).Result()
	if err != nil {
		log.Error(" err != nil", log.ErrObj(err))
		return
	}
	log.Debug("SetNX", log.Bool("val", bVal))
	// SetXX
	bVal, err = rdb.SetXX(ctx, "key", "xxValue", 0).Result()
	if err != nil {
		log.Error(" err != nil", log.ErrObj(err))
		return
	}
	log.Debug("SetXX", log.Bool("val", bVal))
}
