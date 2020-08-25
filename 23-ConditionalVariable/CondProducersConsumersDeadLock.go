package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	条件变量，经常与锁结合使用

	条件变量结构体：
		type Cond struct {
			noCopy noCopy

			// L is held while observing or changing the condition
			L Locker

			notify  notifyList
			checker copyChecker
		}

	条件变量的操作函数：
		Wait方法：
			1、阻塞
			2、释放自己拥有的互斥锁(cond.L.Unlock())，这2步为一个原子操作
			3、当被唤醒后，Wait方法返回时，解除阻塞并重新获取互斥锁(cond.L.Lock())
		Signal方法：
			给一个阻塞在该条件变量上的Go程发送唤醒通知
		Broadcast方法：
			给阻塞在该条件变量上的所有Go程发送唤醒通知

	此实例是一个错误的实例
*/
func main() {
	// 条件变量实现生产者消费者模型
	condProducersConsumers()
}

var cond sync.Cond

func producers(w chan<- int, i int) {
	for {
		cond.L.Lock()
		// 当缓冲区满了，将生产者Go程Wait
		for len(w) == 3 {
			fmt.Printf("写Thread%dWait\n", i)
			cond.Wait()
		}
		num := rand.Intn(1000)
		fmt.Printf("Thread%d写入数据：%d\n", i, num)
		w <- num
		cond.L.Unlock()
		cond.Signal() // 唤醒任意等待的一个Go程
	}
}

func consumers(r <-chan int, i int) {
	for {
		cond.L.Lock()
		// 当缓冲区没有数据可以获取时，将消费者Go程Wait
		for len(r) == 0 {
			fmt.Printf("读Thread%dWait\n", i)
			cond.Wait()
		}
		num := <-r
		fmt.Printf("---Thread%d读取数据：%d\n", i, num)
		cond.L.Unlock()
		cond.Signal() // 唤醒任意等待的一个Go程
	}
}

// 条件变量实现生产者消费者模型
func condProducersConsumers() {
	ch := make(chan int, 3)
	// 设置锁为互斥锁
	cond.L = new(sync.Mutex)

	// 当生产者与消费者Go程数量不一致且相差比较大时，还有一定概率出现死锁情况
	for i := 0; i < 3; i++ {
		go producers(ch, i)
	}

	for i := 0; i < 10; i++ {
		go consumers(ch, i)
	}

	time.Sleep(time.Hour * 2)
}
