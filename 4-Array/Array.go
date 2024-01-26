package main

import "fmt"

// 复合类型
func main() {
	array()
}

// 数组
func array() {
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	var arr2 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr2)

	arr3 := [5]int{100}
	fmt.Println(arr3)

	// 指定数组下标初始化
	arr4 := [5]int{1: 100, 3: 200, 4: 900}
	fmt.Println(arr4)

	// 根据初始化确定数组长度
	arr5 := [...]int{1, 3, 5, 6, 90, 20}
	fmt.Println(arr5)

	var arr7 [5]int
	arr7[3] = 1000
	arr7[4] = 3000
	fmt.Println(arr7)

	var arr8 = []int{10, 20, 30, 40}
	fmt.Println(arr8)
	fmt.Printf("len(arr8) = %d\n", len(arr8))

	for i := 0; i < len(arr5); i++ {
		fmt.Print(arr5[i], ",")
	}
	fmt.Println()

	for i, v := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", i, v)
	}

	for _, v := range arr1 {
		fmt.Print(v, ",")
	}
	fmt.Println()

	param()
	arrCompare()
	twoDimensionalArray()
	arrayCutOut()
}

// 数组作为函数参数
func param() {
	var arr1 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	test1(arr1)
	// GO中数组是值传递，在其他函数中修改传入的数组，不会影响原来的数组
	fmt.Println("arr1[0] =", arr1[0])
	fmt.Println("-----------------------------")
}

func test1(arr [10]int) {
	fmt.Println(arr)
	arr[0] = 100000
}

// 数组比较
func arrCompare() {
	var arr1 = [...]int{0, 1, 2, 3, 4}
	var arr2 = [...]int{0, 1, 2, 3, 4}
	arr3 := [...]int{1, 1, 2, 3, 4}
	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Println("arr1 != num3:", arr1 != arr3)
	fmt.Println("-----------------------------")
}

// 二维数组
func twoDimensionalArray() {
	arr1 := [5][5]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24},
	}

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr1[0]); j++ {
			fmt.Print(arr1[i][j], ",")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
	for _, v1 := range arr1 {
		for _, v := range v1 {
			fmt.Print(v, ",")
		}
		fmt.Println()
	}

	// 列的下标必须指定元素个数
	arr2 := [...][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(arr2)
	fmt.Println("-----------------------------")
}

// 数组截取，同切片
func arrayCutOut() {
	arr5 := [6]int{1, 3, 5, 6, 90, 20}
	// 数组截取出来的是切片
	s1 := arr5[1:2:3]
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s1 = append(s1, 100, 1000)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
