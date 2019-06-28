package main

import (
	"fmt"
	"time"
)

/*
	select是一个关键字，通过select可以监听Channel上的数据流动，多路监听。select中的每个case语句必须是一个IO操作
	select语句中，会按顺序从头到尾判断每一个case，如果其中的任意一个case可以继续执行(没有被阻塞)，那么就从可执行的
	case语句中任意选择一个来执行，如果没有一个case可以执行，当有default时，就执行default分支中的语句，并且结束select
    语句继续往下执行，当没有default时，select语句被阻塞，直到至少有一个case可以被执行。

	注意：for select语句不能直接使用break跳出循环，必须使用break标签或者goto语句
*/
func main() {
	//selectCase1()
	//forSelectBreak()
	//forSelectGoto()
	//fibonacci()
	forSelectTimeout()
}

// select案例1
func selectCase1() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		i := 0
		for {
			select {
			case v := <-ch1:
				fmt.Println("<- ch1 =", v)
			case ch2 <- "test":
				fmt.Println("ch2 <- \"test\"")
			}
			i++
			if i == 10 {
				break
			}
			fmt.Println("i =", i)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			ch1 <- i
			<-ch2
		}
	}()

	time.Sleep(5 * time.Second)
}

// for select break标签退出循环
func forSelectBreak() {
	ch := make(chan string)
	quit := make(chan bool)

	go func() {
		i := 0
		for {
			v := <-ch
			fmt.Println("子Go程读取：", v)
			if i == 50 {
				close(ch)
				quit <- true
			}
			i++
		}
	}()

	i := 0
Loop:
	for {
		select {
		case ch <- fmt.Sprintf("test%d", i):
			fmt.Println("ch <- test", i)
		case <-quit:
			// 不能使用break直接跳出循环
			//break
			break Loop
		}
		i++
	}
}

// for select goto退出循环
func forSelectGoto() {
	ch := make(chan string)
	quit := make(chan bool)

	go func() {
		i := 0
		for {
			fmt.Println("写入数据：test", i)
			ch <- fmt.Sprintf("test%d", i)
			if i == 50 {
				close(ch)
				quit <- true
			}
			i++
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println("读取到数据：", v)
		case <-quit:
			goto EndLoop
		}
	}
EndLoop:
}

func printFibonacci(ch <-chan int, quit chan<- bool, count int) {
	i := 0
	for {
		fmt.Printf("%d ", <-ch)
		if i == count-1 {
			quit <- true
		}
		i++
	}
}

// 斐波那契额数列
func fibonacci() {
	ch := make(chan int)
	quit := make(chan bool)

	go printFibonacci(ch, quit, 20)

	for i, x, y := 0, 0, 1; i < 20; i++ {
		x, y = x+y, x
		ch <- x
	}

Loop:
	for {
		select {
		case <-quit:
			break Loop
		}
	}
}

// for select超时机制
func forSelectTimeout() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-ch:
				fmt.Println("读取到的数据：", v)
			case <-time.After(5 * time.Second):
				// 5秒钟后超时，执行此case，每次for循环进入select时，都会以5秒钟计算超时
				fmt.Println("超时")
				quit <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i
	}

	<-quit
}
