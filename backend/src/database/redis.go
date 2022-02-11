package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	log.Println("Connected with redis...")
}

// SetupCacheChanel creates cache channel and delete cache from redis
func SetupCacheChanel() {
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			key := <-ch

			Cache.Del(context.Background(), key)

			fmt.Println("Cache cleared " + key)
		}
	}(CacheChannel)
}

// ClearCache passes the the key to the SetupCacheChannel
func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
