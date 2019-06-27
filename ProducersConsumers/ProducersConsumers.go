package main

import (
	"fmt"
	"time"
)

// 生产者
func producers(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产者写入：", i)
		ch <- i
	}
	close(ch)
}

// 消费者
func consumers(ch <-chan int) {
	for v := range ch {
		fmt.Println("消费者读取：", v)
	}
}

/*
	生产者消费者
	使用缓冲区的好处：解耦，并发，缓存
*/
func main() {
	ch := make(chan int, 10)
	go producers(ch)
	consumers(ch)
	time.Sleep(5 * time.Second)
}
