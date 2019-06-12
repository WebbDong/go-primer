package main

import (
	"errors"
	"fmt"
)

// 异常处理
func main() {
	//panicFunc()
	//errorInterface()
	recoverFunc()
}

// 1、panic函数
func panicFunc() {
	i := 1000
	if i == 100 {
		// 引发异常，终止后续的程序运行
		panic("Error")
	}

	arr := [5]int{0, 1, 2, 3, 4}
	n := 0
	// 如果数组越界时，系统内部会调用panic函数引发异常，终止程序运行
	fmt.Println(arr[n])

	num1 := 10
	num2 := 0
	// 如果除数为0，系统内部会调用panic函数引发异常，终止程序运行
	fmt.Println(num1 / num2)

	fmt.Println("done...")
}

func TestError(num1 int, num2 int) (res int, err error) {
	if num2 == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = num1 / num2
	return
}

// 2、error接口
func errorInterface() {
	res, err := TestError(10, 0)
	if err == nil {
		fmt.Println("res =", res)
	} else {
		fmt.Println(err)
	}
}

// 3、recover函数
func recoverFunc() {

}
