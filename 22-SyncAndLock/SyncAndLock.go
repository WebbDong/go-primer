package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	同步与锁

	互斥锁：多个线程访问共享数据的时候是串行的，效率低
	读写锁：
		读写锁是一把锁。同时拥有读锁和写锁，读锁用于对内存做读操作，写锁用于对内存做写操作。
		特性：
			1、线程A加读锁成功，又来了三个线程，做读操作，可以加锁成功，不会阻塞，读共享并行处理
			2、线程A加写锁成功，又来了三个线程，做读操作，三个线程阻塞。写独占
			3、线程A加读锁成功，又来了B线程加写锁阻塞，又来了C线程加读锁也阻塞，读写不能同时，写的优先级高于读
*/
func main() {
	//deadLock()
	//mutexLock()
	//rwLock()
	rwLockChannelSimulate()
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

// 读写锁
var rwMutex sync.RWMutex

func read(in <-chan int) {
	for {
		rwMutex.RLock()
		v := <-in
		fmt.Println("读取到的数据：", v)
		rwMutex.RUnlock()
	}
}

func write(out chan<- int) {
	for {
		rwMutex.Lock()
		num := rand.Intn(1000)
		out <- num
		fmt.Println("写入的数据：", num)
		rwMutex.Unlock()
	}
}

// 死锁案例四，读写锁与Channel共用时交叉死锁
func deadLockCase4() {
	/*
		死锁原因分析：当读线程运行时，加了读锁，写线程运行时，由于已经加了读锁，所以说有写线程全部阻塞，Channel是无缓冲的，在读线程中
					读取Channel的时候由于没有数据写入，所以此时读取Channel也阻塞，所有写线程也阻塞，造成死锁。此种类型的死锁不会
					报任何异常。
		建议：读写锁与互斥锁尽量不要与Channel混用
	*/

	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go read(ch)
	}

	for i := 0; i < 5; i++ {
		go write(ch)
	}

	time.Sleep(15 * time.Second)
}

// 1、死锁
func deadLock() {
	//deadLockCase1()
	//deadLockCase2()
	//deadLockCase3()
	deadLockCase4()
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

// 共享变量
var sharedValue int

var rwMutex2 sync.RWMutex

func read2(i int) {
	for {
		// 读锁
		rwMutex2.RLock()
		num := sharedValue
		fmt.Printf("Thread%d读取到数据：%d\n", i, num)
		rwMutex2.RUnlock()
	}
}

func write2(i int) {
	for {
		// 写锁
		rwMutex2.Lock()
		sharedValue = rand.Intn(1000)
		fmt.Printf("Thread%d写数据：%d\n", i, sharedValue)
		rwMutex2.Unlock()
	}
}

// 3、读写锁
func rwLock() {
	for i := 0; i < 5; i++ {
		go read2(i)
	}

	for i := 0; i < 5; i++ {
		go write2(i)
	}
	time.Sleep(20 * time.Second)
}

func read3(r <-chan int, i int) {
	for {
		num := <-r
		fmt.Printf("Thread%d读取数据：%d\n", i, num)
	}
}

func write3(w chan<- int, i int) {
	for {
		num := rand.Intn(1000)
		w <- num
		fmt.Printf("Thread%d写数据：%d\n", i, num)
	}
}

// 4、Channel模拟读写锁的情况
func rwLockChannelSimulate() {
	// 使用Channel无法实现读写锁的功能效果。
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go read3(ch, i)
	}
	for i := 0; i < 5; i++ {
		go write3(ch, i)
	}
	time.Sleep(20 * time.Second)
}
