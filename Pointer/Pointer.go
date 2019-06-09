package main

import (
	"fmt"
	"reflect"
)

// 指针
func main() {
	pointerDef()
	//pointerParam()
	//arrayPointer()
	//arrayPointerParam()
	//pointerArray()
	//slicePointer()
	//structPointer()
	//multilevelPointer()
}

// 指针定义
func pointerDef() {
	num1 := 1000
	num2 := 2000
	num3 := 3000
	var p1 *int = &num1
	p2 := &num2
	var p3 = &num3
	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(*p3)
	fmt.Println(p1)
	fmt.Printf("%p\n", p2)
	fmt.Printf("%p\n", &p2)
	*p3 = 9000
	fmt.Println(num3)
	fmt.Println("---------------------")

	// 空指针
	var p4 *int
	fmt.Println(p4)
	fmt.Println("---------------------")

	// new函数，根据指定的数据类型开辟内存空间
	var p5 = new(int)
	*p5 = 200
	fmt.Println(*p5)
	fmt.Println(p5)

	p6 := new(string)
	*p6 = "Lamborghini"
	fmt.Println(*p6)

	str1 := "Ferrari"
	p7 := &str1
	fmt.Println(*p7)

	var p8 *string
	p8 = &str1
	fmt.Println(*p8)
}

// 指针作为函数参数
func pointerParam() {
	num := 10
	fmt.Println(num)
	funcTest(&num)
	fmt.Println(num)
}

func funcTest(p *int) {
	*p = 8888
}

// 数组指针
func arrayPointer() {
	arr1 := [7]int{0, 1, 2, 3, 4, 5, 6}
	var pa1 *[7]int
	pa1 = &arr1
	fmt.Println(*pa1)
	funcTest2(&arr1)
	fmt.Println(arr1)
	fmt.Println("-------------------------")

	pa2 := new([5]int)
	(*pa2)[0] = 90
	(*pa2)[1] = 80
	(*pa2)[2] = 70
	fmt.Println(*pa2)
	fmt.Println("-------------------------")

	// 第二种访问数组指针元素的方式
	pa3 := new([5]string)
	pa3[0] = "aaaaa"
	pa3[1] = "bbbbb"
	pa3[2] = "ccccc"
	pa3[3] = "ddddd"
	pa3[4] = "ccccc"
	fmt.Println(*pa3)
	fmt.Println("-------------------------")

	for i := 0; i < len(pa3); i++ {
		fmt.Print(pa3[i], ",")
	}
	fmt.Println("-------------------------")
}

func funcTest2(pa *[7]int) {
	(*pa)[1] = 500
}

// 数组指针作为函数参数
func arrayPointerParam() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	funcTest3(&arr1)
	fmt.Println(arr1)
}

func funcTest3(pa *[5]int) {
	pa[3] *= 100
}

// 指针数组
func pointerArray() {
	num := 200
	var parr1 [10]*int
	parr1[0] = new(int)
	*parr1[0] = 100
	parr1[1] = &num
	parr1[2] = new(int)
	*parr1[2] = 300
	fmt.Println("len(parr1) =", len(parr1))
	for i := 0; i < len(parr1); i++ {
		p := parr1[i]
		if !reflect.ValueOf(p).IsNil() {
			fmt.Println(*p)
		}
	}
}

// 指针与切片
func slicePointer() {
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1, 6, 7)

	var p1 *[]int
	p1 = &s1
	*p1 = append(*p1, 8, 9)
	*p1 = append(s1, 10)
	fmt.Println(*p1)
	// 切片的指针，不能像数组指针那样直接使用p1[0]这种形式
	fmt.Println((*p1)[0])
	(*p1)[1] = 100
	fmt.Println(s1)

	p2 := new([]int)
	*p2 = append(*p2, 100, 200, 300)
	fmt.Println(*p2)
	(*p2)[0] = 5000
	fmt.Println(*p2)
}

type Student struct {
	name string
	age  int
	addr string
}

// 指针与结构体
func structPointer() {
	s1 := Student{"Ferrari", 30, "Shanghai"}
	var ps1 = &s1
	fmt.Printf("%p\n", ps1)
	fmt.Println(ps1)
	fmt.Println("name =", ps1.name)
	fmt.Println("age =", ps1.age)
	fmt.Println("addr =", (*ps1).addr)

	*ps1 = Student{name: "Lamborghini"}
	fmt.Println(s1)
	ps1.age = 35
	fmt.Println(*ps1)

	ps2 := new(Student)
	ps2.name = "Test"
	ps2.age = 35
	ps2.addr = "Beijing"
	fmt.Println(*ps2)
}

// 多级指针
func multilevelPointer() {
	var pp1 **int
	num := 1000
	p1 := &num
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Println(*pp1)
	fmt.Println(**pp1)
	fmt.Println("--------------------------")

	pp2 := new(*int)
	*pp2 = new(int)
	**pp2 = 5000
	fmt.Println(**pp2)
}
