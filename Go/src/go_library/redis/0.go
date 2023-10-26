package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func initRedis() {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	result, err := redisDB.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func main() {
	initRedis()
}
