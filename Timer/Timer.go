package main

import (
	"fmt"
	"time"
)

/*
	定时器
*/
func main() {
	createTimer()
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
