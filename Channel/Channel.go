package main

import (
	"fmt"
	"time"
)

type Mobile struct {
	brand string
	cpu   int
	price float64
}

/*
	Channel是一种数据类型，用来传递数据的一种数据结构，Channel可用于两个goroutine之间通过传递一个指定类型的值来
	同步运行和通讯。goroutine奉行通过通信来共享内存，而不是通过共享内存来通信
*/
func main() {
	//declareChannel()
	//channelCommunication()
	//noCacheChannel()
	//cacheChannel()
	//closeChannel()
	forRangeChannel()
}

// 1、定义Channel
func declareChannel() {
	/*
		make(chan Type)  // 等价于make(chan Type, 0)
		make(chan Type, capacity)
		当参数capacity == 0时，Channel是无缓冲阻塞读写的，当capacity > 0时，Channel有缓冲，非阻塞的。直到写满
		capacity个元素才阻塞写入。

		Channel类型变量是引用传递。
	*/

	// Channel变量的默认值是nil
	var c1 chan int
	fmt.Println("c1 =", c1)

	c2 := make(chan string, 5)
	fmt.Println(c2)
	// len：Channel中剩余未读取数据的数量，cap：Channel容量的大小
	fmt.Println("len(c2) =", len(c2), "cap(c2) =", cap(c2))

	c3 := make(chan Mobile)
	fmt.Println(c3)
	fmt.Println("len(c3) =", len(c3), "cap(c3) =", cap(c3))

	/*
		当Channel是无缓冲时，往Channel写数据时，如果没有读出数据，写数据端会阻塞。当Channel是有缓冲时，
		往Channel写数据时，不会阻塞，直到缓冲被写满没有被读取出来写数据端才阻塞

		读数据端只要没有读取到数据就会阻塞
	*/
	// 往Channel里写数据
	c2 <- "channel"
	// 往Channel里读数据
	str, ok := <-c2
	fmt.Print(str, ok)
}

var channel = make(chan bool)

func printString(str string) {
	for _, c := range str {
		fmt.Printf("%c", c)
		time.Sleep(200 * time.Millisecond)
	}
}

func printHello() {
	printString("Hello")
	channel <- true
}

func printWorld() {
	value, ok := <-channel
	printString("World")
	fmt.Println("\nvalue =", value, ", ok =", ok)
}

// 2、Channel用于Go程通信
func channelCommunication() {
	go printHello()
	go printWorld()

	time.Sleep(10 * time.Second)
}

// 3、无缓冲Channel
func noCacheChannel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子Go程写：", i)
			ch <- i
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		v := <-ch
		fmt.Println("主Go程读：", v)
	}

	time.Sleep(10 * time.Second)
}

// 4、有缓冲Channel
func cacheChannel() {
	ch := make(chan int, 5)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("子Go程写：", i)
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		v := <-ch
		fmt.Println("主Go程读：", v)
	}
	time.Sleep(10 * time.Second)
}

// 5、关闭Channel，已经关闭的Channel不能再向它写数据，但是能读取已经关闭的Channel
func closeChannel() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭Channel
		close(ch)
	}()
	time.Sleep(3 * time.Second)
	for {
		if v, ok := <-ch; ok {
			fmt.Println("读到的值：", v)
		} else {
			v = <-ch
			fmt.Println("Channel关闭")
			fmt.Println("关闭后读取的值：", v)
			break
		}
	}

	time.Sleep(10 * time.Second)
}

// 6、for range遍历Channel
func forRangeChannel() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	// for range能自动判断Channel是否已经关闭，如果已经关闭自动跳出循环
	// 如果使用for range，没有进行Channel的关闭操作，那么将导致死锁异常
	for v := range ch {
		fmt.Println("读取到的值：", v)
	}

	time.Sleep(5 * time.Second)
}
