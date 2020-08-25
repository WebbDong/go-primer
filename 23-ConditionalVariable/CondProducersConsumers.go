package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	条件变量实现生产者消费者模型

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
