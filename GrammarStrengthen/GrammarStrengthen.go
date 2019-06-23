package main

import "fmt"

// GO语法加强
func main() {
	//pointerStrengthen()
	//sliceStrengthen()
	mapStrengthen()
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

// 2、切片加强
func sliceStrengthen() {
	/*
		为什么用切片：
			1、数组的容量固定，不能自动扩展
			2、数组值传递。数组作为函数参数时，是将整个数组拷贝一份给形参，切片是引用传递
		切片的本质：
			1、不是一个数组指针，是一种数据结构体。
	*/

	// 练习，去除空字符串
	strs := []string{"red", "", "black", "", "", "pink", "blue"}
	i := 0
	for _, v := range strs {
		if v != "" {
			strs[i] = v
			i++
		}
	}
	strs = strs[:i]
	fmt.Println(strs)
	fmt.Println("--------------------------------")

	// 联系，去重
	str2 := []string{"red", "black", "red", "pink", "blue", "red", "pink", "blue"}
	newStr := make([]string, 0, 10)
	var isExists bool
	for _, v := range str2 {
		isExists = false
		for i, len := 0, len(newStr); i < len; i++ {
			if v == newStr[i] {
				isExists = true
				break
			}
		}
		if !isExists {
			newStr = append(newStr, v)
		}
	}
	fmt.Println(newStr)
	fmt.Println(len(newStr))
}

// 3、Map加强
func mapStrengthen() {
	// 只是声明了变量，map1为nil，没有初始化
	var map1 map[int]string
	fmt.Println(map1 == nil)

	map2 := map[int]string{}
	fmt.Println(map2 == nil)

	map3 := make(map[int]string)
	fmt.Println(map3 == nil)

	map4 := make(map[int]string, 5)
	map4[100] = "Lamborghini"
	map4[50] = "Ferrari"
	map4[10] = "Pagani"
	map4[5] = "BMW"
	map4[2] = "BENZ"
	map4[0] = "TOYOTA"
	// 判断key是否存在
	if v, exists := map4[0]; exists {
		fmt.Println("v =", v)
	} else {
		fmt.Println("不存在")
	}
}
