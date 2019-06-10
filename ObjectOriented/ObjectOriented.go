package main

import "fmt"

// GO语言使用结构体定义类
type Student struct {
	id   int
	name string
	age  int
	addr string
}

// 面向对象
func main() {
	var s1 = Student{1, "dongwenbin", 28, "Shanghai"}
	s2 := Student{id: 2, name: "Ferrari", age: 90, addr: "Italy"}
	fmt.Println(s1)
	fmt.Println(s2)
}
