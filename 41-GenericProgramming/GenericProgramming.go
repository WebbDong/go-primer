package main

import (
	"./model"
	"fmt"
)

// any 关键字就是 interface{}，是 interface{} 的一个别名，定义语句：type any = interface{}

// 1.泛型自定义类型
type MyMap[K comparable, V interface{ int | float32 }] map[K]V

// 2.泛型函数
func genericFunc1(x any, y any) any {
	fmt.Println(x)
	fmt.Println(y)
	return "Ferrari"
}

// any 泛型 T 不能用户函数返回值
func genericFunc2[T any](x T) any {
	fmt.Println(x)
	return model.Car{Brand: "Lamborghini", HousePower: 800}
}

func genericFunc3[T1 int, T2 float32](x T1, y T2) {
	fmt.Println(x)
	fmt.Println(y)
}

func genericFunc4[K comparable, V int | float32](map1 map[K]V, key K) V {
	return map1[key]
}

func genericFunc5(map1 MyMap[string, int], key string) int {
	return map1[key]
}

// 3.自定义泛型结构体
type Node[T any] struct {
	Next *Node[T]
	Data T
}

func main() {
	fmt.Println(genericFunc1(40, 500))
	fmt.Println("...........................")
	fmt.Println(genericFunc2(400))
	fmt.Println("...........................")
	genericFunc3(400, 50.55)
	fmt.Println("...........................")
	map1 := make(MyMap[string, int])
	map1["Kobe"] = 40
	map1["Curry"] = 30
	fmt.Println(map1)
	fmt.Println("...........................")
	map2 := make(MyMap[string, float32])
	map2["key1"] = 50.55
	map2["key2"] = 60.88
	fmt.Println(genericFunc4(map2, "key1"))
	fmt.Println("...........................")
	map3 := make(MyMap[string, int])
	map3["Ferrari"] = 800
	map3["Lamborghini"] = 780
	fmt.Println(genericFunc5(map3, "Ferrari"))
	fmt.Println("...........................")
	//car1 := model.Car{Brand: "Ferrari", HousePower: 800}
	node1 := Node[int]{Data: 500, Next: &Node[int]{Data: 300, Next: nil}}
	fmt.Println(node1)
	fmt.Println("...........................")
	node2 := Node[model.Car]{Data: model.Car{Brand: "Lamborghini", HousePower: 900}}
	fmt.Println(node2)
}
