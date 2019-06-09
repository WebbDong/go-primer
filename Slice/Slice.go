package main

import (
	"fmt"
)

// 切片，相当于动态数组的功能
// 切片的底层是一个结构体
/*
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/
func main() {
	//slice()
	//sliceInit()
	//sliceLoop()
	//sliceCutOut()
	//sliceValueEdit()
	//appendFunc()
	//copyFunc()
	//sliceParam()
	sliceParamAppend()
}

// 切片创建
func slice() {
	// 空切片
	var s1 []int
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	fmt.Println(len(s1))

	s2 := []int{}
	fmt.Println(s2)

	var s3 = []int{1, 2, 3, 4, 5}
	fmt.Println(s3)

	// 使用make函数创建切片，make(切片类型, 长度, 容量)
	// 长度是已经初始化的空间，容量是已经开辟的空间，包含已经初始化的空间和空闲的空间
	// 长度不能大于容量，容量参数可以省略
	var s4 = make([]int, 10, 20)
	fmt.Println(s4)
	s5 := make([]int, 20, 30)
	fmt.Println(s5)
	// 省略容量时，长度和容量相同
	s6 := make([]int, 5)
	fmt.Println(s6)
	fmt.Println(len(s6))
	// 获取容量
	fmt.Println(cap(s6))
}

// 切片初始化
func sliceInit() {
	var s1 []int
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	// 不允许用角标追加元素，必须使用append函数，角标只能修改或获取
	//s1[0] = 0
	// 追加数据
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Println(s1)
	s1 = append(s1, 100, 200)
	fmt.Println(s1)
	fmt.Println("s1[5]=", s1[5])
	s1[5] = 1000
	fmt.Println("s1[5]=", s1[5])

	fmt.Println("-----------------------------------")

	s2 := []string{"test1", "test2", "test3"}
	fmt.Println(s2)
	fmt.Println("len(s2) =", len(s2))
	fmt.Println("cap(s2) =", cap(s2))
	s2 = append(s2, "test4", "test5")
	fmt.Println(s2)
	fmt.Println("len(s2) =", len(s2))
	fmt.Println("cap(s2) =", cap(s2))
	fmt.Println(s2)

	s3 := make([]int, 5, 10)
	s3[0] = 50
	s3[1] = 60
	s3[2] = 70
	fmt.Println(s3)
	s3 = append(s3, 80, 90)
	fmt.Println(s3)

	// 循环必须根据长度，而不是容量
	for i := 0; i < len(s3); i++ {
		s3[i] = i
	}
	fmt.Println(s3)
}

// 切片循环
func sliceLoop() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(s1); i++ {
		fmt.Print(s1[i], " ")
	}
	fmt.Println()

	for i, v := range s1 {
		fmt.Printf("s1[%d] = %d\n", i, v)
	}

	for _, v := range s1 {
		fmt.Print(v, " ")
	}
}

// 切片截取
func sliceCutOut() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("cap(s1) =", cap(s1))
	// s[n]， 获取指定索引的元素，类似数组
	fmt.Println("s1[1] =", s1[1])
	fmt.Println("------------------------------------")

	// 对切片进行截取，会返回一个新的切片，对返回的新切片进行修改，会影响到原来的切片
	// 非s[low:high:max]形式的截取方式，默认以原切片长度减去low，如果low没有指定默认为0，
	// s[:]， 从切片的索引位置0到len - 1处所获得的切片
	s2 := s1[:]
	fmt.Println("s1[:] =", s2)
	fmt.Println("len(s2) =", len(s2))
	fmt.Println("cap(s2) =", cap(s2))
	fmt.Println("------------------------------------")

	// s[low:]，从切片的索引位置low到len - 1处所获得的切片
	s3 := s1[1:]
	fmt.Println("s1[1:] =", s3)
	fmt.Println("len(s3) =", len(s3))
	fmt.Println("cap(s3) =", cap(s3))
	fmt.Println("------------------------------------")

	// s[:high]，从切片的索引位置0到high处所获得的切片，不包含high索引的元素，相当于长度len = high
	s4 := s1[:8]
	fmt.Println("s1[:8] =", s4)
	fmt.Println("len(s4) =", len(s4))
	fmt.Println("cap(s4) =", cap(s4))
	fmt.Println("------------------------------------")

	// s[low:high]，从切片的索引位置low到high处所获得的切片，不包含high索引的元素，len = high - low
	s5 := s1[1:8]
	fmt.Println("s1[1:8] =", s5)
	fmt.Println("len(s5) =", len(s5))
	fmt.Println("cap(s5) =", cap(s5))
	fmt.Println("------------------------------------")

	// s[low:high:max]，从切片的索引位置low到high处所获得的切片，len = high - low，cap = max - low
	// max为切片保留的原切片的最大下标（不含max）max不能超过原切片的长度，必须low <= high <= max
	s6 := s1[1:8:10]
	fmt.Println("s1[1:8:10] =", s6)
	fmt.Println("len(s6) =", len(s6))
	fmt.Println("cap(s6) =", cap(s6))
	fmt.Println("------------------------------------")

	// 长度
	fmt.Println("len(s1) =", len(s1))

	// 容量
	fmt.Println("cap(s1) =", cap(s1))
}

// 切片值的修改
func sliceValueEdit() {
	// 对切片进行截取，会返回一个新的切片，对返回的新切片进行修改，会影响到原来的切片
	// 截取并没有开辟新的内存
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("s1 =", s1)
	s2 := s1[3:]
	fmt.Println("s2 =", s2)

	s2[0] = 1000
	fmt.Println("s2 =", s2)
	fmt.Println("s1 =", s1)
}

// append函数
func appendFunc() {
	s1 := make([]int, 5, 10)
	fmt.Println(s1)
	fmt.Println("len(s1) =", len(s1))
	fmt.Println("cap(s1) =", cap(s1))
	fmt.Println("-----------------------------------------")
	// append函数从切片的末尾追加数据，容量不够时，会进行扩容
	// 一般扩容方式为上一次：容量*2，如果超过1024字节，每次扩容上一次的1/4
	s1 = append(s1, 6)
	s1 = append(s1, 7)
	s1 = append(s1, 8)
	s1 = append(s1, 9)
	s1 = append(s1, 10)
	s1 = append(s1, 11)
	fmt.Println(s1)
	fmt.Println("len(s1) =", len(s1))
	fmt.Println("cap(s1) =", cap(s1))
}

// copy函数
func copyFunc() {
	s1 := make([]string, 5, 10)
	s1[0] = "test1"
	s1[1] = "test2"
	s1[2] = "test3"
	s1[3] = "test4"
	s1[4] = "test5"
	s1 = append(s1, "Ferrari")
	s1 = append(s1, "Lamborghini")
	fmt.Println(s1)

	s2 := make([]string, 5, 10)
	// 将s1中的内容拷贝到s2中，如果2个切片长度不同。只会拷贝较小长度的内容。
	copy(s2, s1)
	fmt.Println(s2)

	s3 := []int{1, 2}
	s4 := []int{3, 4, 5, 6, 7}
	copy(s4, s3)
	fmt.Println(s4)

	s5 := []int{1, 2}
	s6 := []int{3, 4, 5, 6, 7}
	copy(s5, s6)
	fmt.Println(s5)
}

// 切片当作函数参数
// 修改切片参数的元素值时，会影响原切片
func sliceParam() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	sliceParamTest(s)
	fmt.Println(s)
}

func sliceParamTest(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = 2
	}
}

// 切片当作函数参数，再新函数在中使用append不会影响原切片
// append如果内存空间不够会开辟一个新的内存空间将原数据复制，然后追加数据
func sliceParamAppend() {
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println("sliceParamAppend-s1:", s1)
	sliceParamTest2(s1)
	fmt.Println("sliceParamAppend-s1:", s1)
	fmt.Println("-----------------------------------")

	s2 := make([]int, 5, 20)
	s2[0] = 0
	s2[1] = 1
	s2[2] = 2
	s2[3] = 3
	s2[4] = 4
	fmt.Printf("sliceParamAppend-s2:%p\n", &s2)
	fmt.Println("sliceParamAppend-s2:", s2)
	sliceParamTest2(s2)
	fmt.Println("sliceParamAppend-s2:", s2)
}

func sliceParamTest2(s []int) {
	newSlice := append(s, 6, 7, 8, 9, 10)
	fmt.Printf("s:%p\n", &s)
	fmt.Printf("newSlice:%p\n", &newSlice)
	fmt.Println("sliceParamTest2:", newSlice)
	s = newSlice
}
