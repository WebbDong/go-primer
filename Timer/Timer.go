package main

import (
	"fmt"
	"time"
)

/*
	定时器
*/
func main() {
	//createTimer()
	//stopTimer()
	resetTimer()
}

// 1、创建定时器
func createTimer() {
	fmt.Println("系统当前时间：", time.Now())
	// 指定定时时长，定时到后，系统会自动向管道C写入系统当前时间
	timer := time.NewTimer(10 * time.Second)
	t := <-timer.C
	fmt.Println("<- timer.C：", t)
	fmt.Println("-----------------------")

	// time.Sleep方式定时
	time.Sleep(2 * time.Second)

	// time.After方式定时
	fmt.Println("系统当前时间：", time.Now())
	timeCh := time.After(5 * time.Second)
	fmt.Println("<- timeCh：", <-timeCh)
}

// 2、停止定时器
func stopTimer() {
	fmt.Println("当前系统时间：", time.Now())
	timer := time.NewTimer(5 * time.Second)
	go func() {
		// 停止定时器后，系统不会向C中写入时间数据，从管道C中也不会读取到数据，后续也不会执行
		t := <-timer.C
		fmt.Println("子Go程读取到的时间：", t)
	}()
	// 停止定时器
	isStop := timer.Stop()
	if isStop {
		fmt.Println("定时器关闭")
	}
	// 读取关闭后的定时器C管道会导致死锁
	//<- timer.C
	time.Sleep(10 * time.Second)
}

// 3、重置定时器
func resetTimer() {
	fmt.Println("当前系统时间：", time.Now())
	timer := time.NewTimer(20 * time.Second)
	// 重置定时器定时时间
	timer.Reset(2 * time.Second)
	go func() {
		t := <-timer.C
		fmt.Println("子Go程读取到的时间：", t)
	}()
	time.Sleep(8 * time.Second)
}
