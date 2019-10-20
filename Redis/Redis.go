package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	if _, err := conn.Do("AUTH", "LocalRedis"); err != nil {
		fmt.Println(err)
		return
	}
	conn.Do("set", "key1", "helloworld")
}
