package main

import "fmt"

// GO语法加强
func main() {
	pointerStrengthen()
}

// 1、指针加强
func pointerStrengthen() {
	// 使用new创建的是在堆内存中的
	p := new(int)
	*p = 50
	fmt.Println(*p)

	pstr1 := test1()
	fmt.Println(*pstr1)

	pint := test2()
	fmt.Println(*pint)
}

func test1() *string {
	// 在堆内存中创建的，所以test1函数调用结束后，数据依然存在
	p := new(string)
	*p = "Lamborghini"
	return p
}

func test2() *int {
	var num = 500
	// go语言编译器会自动决定把一个变量放在栈还是放在堆，编译器会做逃逸分析(escape analysis)，
	// 当发现变量的作用域没有跑出函数范围，就可以在栈上，反之则必须分配在堆。
	return &num
}
