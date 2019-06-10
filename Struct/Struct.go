package main

import (
	"fmt"
)

// 结构体
// Go语言结构体中不允许直接定义函数
func main() {
	//structInit()
	//structArray()
	//structSlice()
	//structMap()
	//structParam()
	structScan()
}

// 结构体创建与初始化
func structInit() {
	// 结构体成员变量不能再前面添加var
	type Student struct {
		id   int
		name string
		age  int
		addr string
	}

	var s1 Student
	s1.id = 5
	s1.age = 28
	s1.name = "Ferrari"
	s1.addr = "ShangHai"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	fmt.Println("-----------------------------------")

	s2 := Student{10, "Lamborghini", 30, "Beijing"}
	fmt.Println(s2)
	fmt.Println("name =", s2.name)
	fmt.Println("age =", s2.age)
	fmt.Println("-----------------------------------")

	// 指定成员变量名初始化，可以部分成员初始化
	s3 := Student{name: "Pagani", age: 40}
	fmt.Println(s3)
}

type Car struct {
	horsepower int
	brand      string
	topSpeed   int
}

// 结构体数组
func structArray() {
	var sa1 = [5]Car{
		{800, "Ferrari", 350},
		Car{horsepower: 1000, brand: "Lamborghini", topSpeed: 350},
	}
	sa1[4] = Car{1200, "Bugatti", 420}
	fmt.Println(sa1)
	fmt.Println(sa1[0].brand)
	sa1[0].topSpeed = 320
	fmt.Println(sa1)
	fmt.Println("len(sa1) =", len(sa1))
	fmt.Println("---------------------------------")

	for i := 0; i < len(sa1); i++ {
		c := sa1[i]
		fmt.Println(c)
	}
	fmt.Println("---------------------------------")

	for _, v := range sa1 {
		fmt.Println(v)
	}
}

// 结构体切片
func structSlice() {
	ss1 := []Car{
		{800, "Ferrari", 350},
		{1500, "Nissan", 400},
	}
	ss1 = append(ss1, Car{1000, "Bugatti", 450})
	fmt.Println(ss1)
	fmt.Println(ss1[len(ss1)-1].brand)
}

// 结构体Map
func structMap() {
	m1 := map[int]Car{
		0: {800, "Ferrari", 350},
		1: {horsepower: 1000, brand: "Pagani", topSpeed: 360},
	}
	m1[2] = Car{1500, "Nissan", 400}
	fmt.Println(m1)

	for _, v := range m1 {
		fmt.Println(v)
	}
}

func structParam() {
	c1 := Car{1500, "Nissan", 380}
	funcTest(c1)
	fmt.Println(c1)
}

func funcTest(c Car) {
	fmt.Println(c)
	// 对结构体的成员变量修改不会影响原结构体
	c.horsepower = 190
}

// Scan
func structScan() {
	var c1 Car
	fmt.Println("输入数据：")
	fmt.Scan(&c1.horsepower, &c1.brand, &c1.topSpeed)
	fmt.Println(c1)
}
