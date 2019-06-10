package main

import "fmt"

// For Range分析
// 1、作用：用于遍历容器类型的数据，例如数组、切片、map等
// 2、range的本质是一个函数方法，使用时候可以加括号
// 3、修改range得到的value，不影响原切片或者数组
func main() {
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range arr1 {
		// 修改得到后的value，不影响原数据或切片
		v = 900
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(arr1)
	fmt.Println("--------------------------------------")

	for i, v := range arr1 {
		arr1[i] *= 10
		fmt.Println("value =", v)
	}
	fmt.Println(arr1)
}
