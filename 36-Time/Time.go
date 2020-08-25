package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 单位秒
	fmt.Println(now.Unix())
	// 单位纳秒
	fmt.Println(now.UnixNano())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.Format("2006-01-02 15:04:05"))

	duration, err := time.ParseDuration("2160h")
	if err != nil {
		fmt.Println(err)
		return
	}
	newTime := now.Add(duration)
	fmt.Println(newTime.Format("2006-01-02 15:04:05"))

	t := time.Unix(newTime.Unix(), 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
