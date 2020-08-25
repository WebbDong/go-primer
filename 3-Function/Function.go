package main

import "fmt"

// 不能使用:=定义全局变量
//global := 30
var global = 30

func add(num1 int, num2 int) int {
	return num1 + num2
}

// 可变参数，可变参数必须是函数的最后一个参数
func sum(num ...int) int {
	var sum = 0
	/*	for i := 0; i < len(num); i++ {
		sum += num[i]
	}*/

	// for range 遍历集合
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	/*
		for i, v := range num {
			// i为数组的角标，v为数组的值
			fmt.Printf("num[%d] = %d\n", i, v)
			sum += v
		}
	*/

	// _为匿名变量，匿名变量不会保存数据。当不想要使用角标时，使用匿名变量
	for _, v := range num {
		sum += v
	}
	return sum
}

// 等价于创建了sum变量
func subtraction(num1 int, num2 int) (sum int) {
	sum = num1 - num2
	return sum
}

// 等价于创建了sum变量，如果指定了返回变量名，return可以不指定变量名
func subtraction2(num1 int, num2 int) (sum int) {
	sum = num1 - num2
	return
}

// 返回多个值
func funcReturnMore() (int, int) {
	num1 := 100
	num2 := 300
	return num1, num2
}

// 返回多个值，指定返回变量名
func funcReturnMore2() (n1 int, n2 int) {
	n1 = 1000
	n2 = 5000
	return
}

// 递归，阶乘
func factorial(num int64) int64 {
	if num == 1 {
		return 1
	}
	res := num * factorial(num-1)
	return res
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sum(3, 6, 7))
	fmt.Println(subtraction(9, 8))
	fmt.Println(subtraction2(10, 8))

	num1, num2 := funcReturnMore()
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)

	n1, n2 := funcReturnMore2()
	fmt.Println("n1 =", n1)
	fmt.Println("n2 =", n2)

	// 延迟执行函数defer关键字，defer会等它后面的非defer修饰的函数执行完后，再执行defer函数
	// defer执行顺序：如果一个函数有多个defer，它们会以LIFO(后进先出)的顺序执行。也就是后面的defer函数优先执行
	defer fmt.Println("第一个")
	fmt.Println("第二个")
	fmt.Println("第三个")
	fmt.Println("第四个")

	defer fmt.Println("aaaa")
	defer fmt.Println("bbbb")
	defer fmt.Println("cccc")

	fmt.Printf("factorial(%d) = %d\n", 10, factorial(10))
}
