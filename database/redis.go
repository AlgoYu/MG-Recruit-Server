package database

import "github.com/go-redis/redis"

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}
