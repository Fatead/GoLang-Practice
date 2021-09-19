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
	_, err := Client.Do("set", "catName", "tom").Result()
	if err != nil {
		fmt.Println("set出错")
	}
	result, err := Client.Do("Get", "catName").Result()
	if err != nil {
		fmt.Println("Get error", err)
		return
	}
	fmt.Println(result)

	//通过GoLang对hash数据类型进行操作
	_, err = Client.Do("HSet", "user01", "name", "john").Result()
	if err != nil {
		fmt.Println("HSet error")
	}
	_, err = Client.Do("HSet", "user01", "age", "18").Result()
	if err != nil {
		fmt.Println("Hset error 02")
	}
	name, err := Client.Do("HGet", "user01", "name").Result()
	fmt.Println(name)
	age, err := Client.Do("HGet", "user01", "age").Result()
	fmt.Println(age)

	//通过GoLang对Redis进行操作，一次操作多个key-val 数据
	_, err = Client.Do("MSet", "name", "尚硅谷", "adress", "北京昌平").Result()
	res, err := Client.Do("MGet", "name", "adress").Result()
	fmt.Println(res)

	//对list进行操作
	Client.RPush("list", "val1", "val2")
	Client.LPush("list", "val0")
	r, err := Client.RPop("list").Result()
	fmt.Println(r)

}
