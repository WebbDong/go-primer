package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Gun struct {
	Caliber         float32
	InitialVelocity int
}

// go get github.com/gomodule/redigo/redis
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
	//basic(conn)
	//structBytesRedis(conn)
	structJsonRedis(conn)
}

func basic(conn redis.Conn) {
	conn.Do("set", "key1", "helloworld")
	conn.Do("set", "key2", "Ferrari")
	conn.Do("set", "key3", "测试")

	reply, err := conn.Do("get", "key1")
	value, err := redis.String(reply, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key1 =", value)

	reply, err = conn.Do("mget", "key1", "key2")
	values, err := redis.Strings(reply, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(values)

	reply, err = conn.Do("get", "key3")
	value, err = redis.String(reply, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("key3 =", value)
}

// 将struct存入redis，从redis读取struct
func structBytesRedis(conn redis.Conn) {
	gun := Gun{Caliber: 0.5, InitialVelocity: 800}
	guns := []Gun{{Caliber: 1.0, InitialVelocity: 1000}, gun}
	// 序列化存入redis
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(&guns)
	conn.Do("set", "gun1", buffer.Bytes())

	// 反序列化从redis取出
	var gunsRes []Gun
	reply, err := conn.Do("get", "gun1")
	datas, err := redis.Bytes(reply, err)
	decoder := gob.NewDecoder(bytes.NewReader(datas))
	decoder.Decode(&gunsRes)
	fmt.Println("res:", gunsRes)
}

func structJsonRedis(conn redis.Conn) {
	guns := []Gun{{Caliber: 1.0, InitialVelocity: 1000}, {Caliber: 0.5, InitialVelocity: 800}}
	bytes, err := json.Marshal(guns)
	conn.Do("set", "gun2", bytes)

	var gunsRes []Gun
	reply, err := conn.Do("get", "gun2")
	bytes, err = redis.Bytes(reply, err)
	err = json.Unmarshal(bytes, &gunsRes)
	fmt.Println("res:", gunsRes)
}
