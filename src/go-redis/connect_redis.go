package main

import (
	"fmt"
	"github.com/go-redis/redis" //引入redis的包
	"time"
)

var Client *redis.Client

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		PoolSize:     1000,
		ReadTimeout:  time.Millisecond * time.Duration(100),
		WriteTimeout: time.Millisecond * time.Duration(100),
		IdleTimeout:  time.Second * time.Duration(60),
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic("init redis error")
	} else {
		fmt.Println("init redis OK")
	}
}

func get(key string) (string, bool) {
	r, err := Client.Get(key).Result()
	if err != nil {
		return "", false
	}
	return r, true
}

func set(key string, val string, expTime int32) {
	Client.Set(key, val, time.Duration(expTime)*time.Second)
}

func main() {
	//通过Go向redis写入数据和读取数据
	set("name", "x", 100)
	s, b := get("name")
	fmt.Println(s, b)
}
