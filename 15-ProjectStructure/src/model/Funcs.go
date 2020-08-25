package model

import "fmt"

// 首字母为小写的函数，只在当前包下可见。属于包私有函数
func func1() {
	fmt.Println("func1")
}

// 首字母为大写的函数，为公有函数，可以在其他包中调用
func Func2() {
	fmt.Println("Func2")
}
