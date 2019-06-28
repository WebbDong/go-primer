package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	同步与锁

*/
func main() {
	//deadLock()
	mutexLock()
}

// 死锁案例一，单Go程自己死锁
func deadLockCase1() {
	ch := make(chan int)
	ch <- 100
	num := <-ch
	fmt.Println("num =", num)
}

// 死锁案例二，Go程操作Channel顺序不对导致死锁
func deadLockCase2() {
	ch := make(chan int)
	ch <- 100

	go func() {
		fmt.Println("value =", <-ch)
	}()
}

// 死锁案例三，多Go程，多Channel交叉死锁
func deadLockCase3() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch2 <- v
			}
		}
	}()

	select {
	case v := <-ch2:
		ch1 <- v
	}
}

// 1、死锁
func deadLock() {
	//deadLockCase1()
	//deadLockCase2()
	deadLockCase3()
}

// 互斥锁
var mutex sync.Mutex

func print(str string) {
	mutex.Lock()
	for _, c := range str {
		fmt.Printf("%c", c)
		time.Sleep(200 * time.Millisecond)
	}
	mutex.Unlock()
}

func person1() {
	print("Hello ")
}

func person2() {
	print("World")
}

// 2、互斥锁
func mutexLock() {
	go person1()
	go person2()
	time.Sleep(5 * time.Second)
}
