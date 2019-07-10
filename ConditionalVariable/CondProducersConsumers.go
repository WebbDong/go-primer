package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	条件变量实现生产者消费者模型
*/
func main() {
	// 条件变量实现生产者消费者模型
	condProducersConsumers2()
}

var cond2 sync.Cond

func producers2(w chan<- int, i int) {
	for {
		cond2.L.Lock()
		// 当缓冲区满了，将生产者Go程Wait
		for len(w) == 3 {
			fmt.Printf("写Thread%dWait\n", i)
			cond2.Wait()
		}
		num := rand.Intn(1000)
		fmt.Printf("Thread%d写入数据：%d\n", i, num)
		w <- num
		cond2.L.Unlock()
		cond2.Signal() // 唤醒任意等待的一个Go程
	}
}

func consumers2(r <-chan int, i int) {
	for {
		cond2.L.Lock()
		// 消费者Go程不Wait，使用select防止channel死锁。当读取到数据是打印
		// Signal只唤醒生产者Go程
		select {
		case num := <-r:
			fmt.Printf("---Thread%d读取数据：%d\n", i, num)
		default:
		}
		cond2.L.Unlock()
		cond2.Signal() // 唤醒生产者
	}
}

// 条件变量实现生产者消费者模型
func condProducersConsumers2() {
	ch := make(chan int, 3)
	// 设置锁为互斥锁
	cond2.L = new(sync.Mutex)

	// 当生产者与消费者Go程数量不一致且相差比较大时，还有一定概率出现死锁情况
	for i := 0; i < 3; i++ {
		go producers2(ch, i)
	}

	for i := 0; i < 10; i++ {
		go consumers2(ch, i)
	}

	time.Sleep(time.Hour * 2)
}
