package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	Goroutine是Go语言中的协程(coroutine)，简称Go程，协程是一种轻量级的线程，一个线程可以拥有多个协程，
	共享一个线程资源。协程不是被操作系统内核所管理，完全由程序控制，不会有线程的系统级别上下文切换，
    性能有很大的提高。

	主Go程结束之后，子Go程也会结束
*/
func main() {
	//useGoroutine()
	//gosched()
	//goexit()
	//gomaxprocs()
	otherRuntimeFunc()
	time.Sleep(10 * time.Second)
}

func eat() {
	for i := 0; i < 50; i++ {
		fmt.Println("吃火锅")
	}
}

func walk() {
	for i := 0; i < 50; i++ {
		fmt.Println("走路")
	}
}

// 1、使用Go程
func useGoroutine() {
	// 使用go关键字进行开启协程
	go eat()
	go walk()
}

func speak() {
	for {
		// 出让CPU时间片，当再次获取时间片时继续向下执行
		runtime.Gosched()
		fmt.Println("Speak")
	}
}

// 2、runtime.Gosched()的作用是出让CPU的时间片，给其他协程使用
func gosched() {
	go speak()
	for {
		fmt.Println("Run")
	}
}

func test() {
	defer fmt.Println("cccccc")
	// 终止当前Go程
	runtime.Goexit()
	defer fmt.Println("dddddd")
}

// 3、runtime.Goexit()结束当前的go程，在执行它之前的defer函数生效
func goexit() {
	go func() {
		defer fmt.Println("aaaaaa")
		// test函数没有使用go关键字，所以也在主Go程中，与匿名函数在同一个Go程
		// 所以在test函数中Goexit，也一起Goexit了
		test()
		//go test()
		defer fmt.Println("bbbbbb")
	}()
}

// 4、runtime.GOMAXPROCS(n)用来设置可以并行计算的CPU最大核数，返回之前的值
func gomaxprocs() {
	// 在英特尔的CPU中，有多线程技术，四核八线程，设置时返回的是8，返回的是逻辑CPU个数
	i := runtime.GOMAXPROCS(1) // 设置为单核
	fmt.Println("原核心数：", i)
	for {
		go fmt.Print(0) // 子Go程
		fmt.Print(1)    // 主Go程
	}
}

// 5、其他runtime函数
func otherRuntimeFunc() {
	// 返回GOROOT目录
	fmt.Println(runtime.GOROOT())
	// 返回go的版本
	fmt.Println(runtime.Version())
	// 返回逻辑CPU个数
	fmt.Println(runtime.NumCPU())
	// 执行一次垃圾回收
	runtime.GC()
}
