package main

import "fmt"

// Map的Key不能是引用类型
func main() {
	//mapInit()
	//mapKeyAndValue()
	//mapParam()
	statisticsLetter()
}

// map创建与初始化
func mapInit() {
	var m1 map[int]string = map[int]string{0: "test1", 1: "test2", 2: "test3", 3: "test4"}
	m1[4] = "test5"
	m1[5] = "test6"
	fmt.Println(m1)
	fmt.Println("len(m1) =", len(m1))
	fmt.Println("---------------------------------------")

	m2 := map[int]string{0: "ferrari", 1: "lamborghini", 2: "pagani"}
	m2[3] = "bugatti"
	fmt.Println(m2)
	fmt.Println("len(m2) =", len(m2))
	fmt.Println("---------------------------------------")

	m3 := make(map[int]string)
	m3[0] = "aaaa"
	m3[1] = "bbbb"
	m3[2] = "cccc"
	fmt.Println(m3)
	fmt.Println("len(m3) =", len(m3))
	fmt.Println("---------------------------------------")

	// 指定长度
	m4 := make(map[int]string, 10)
	m4[0] = "a1"
	m4[1] = "a2"
	m4[2] = "a3"
	m4[3] = "a4"
	// 修改
	m4[3] = "dddd"
	fmt.Println(m4)
	fmt.Println("len(m4) =", len(m4))
	fmt.Println("---------------------------------------")

	m5 := map[string]int{"test1": 1000, "test2": 2000, "test3": 3000}
	fmt.Println(m5)
}

// map键与值
func mapKeyAndValue() {
	m1 := map[string]int{"key1": 1000, "key2": 2000, "key3": 3000, "key4": 4000}
	// value为值，existsKey为boolean，指定的key是否存在
	value, existsKey := m1["key1"]
	if existsKey {
		fmt.Println("value =", value, ",existsKey =", existsKey)
	} else {
		fmt.Println("指定的key不存在")
	}

	// 循环遍历map
	for key, value := range m1 {
		fmt.Println(key, "=", value)
	}

	// 根据指定的key删除值
	delete(m1, "key1")
	fmt.Println(m1)
}

// map作为函数参数
func mapParam() {
	m := map[int]string{0: "test1", 1: "test2", 2: "test3", 3: "test4"}
	funcTest(m)
	fmt.Println(m)
}

func funcTest(m map[int]string) {
	for key, value := range m {
		fmt.Println(key, "=", value)
	}
	// 对map进行修改会影响到原map
	delete(m, 0)
}

// 统计字母
func statisticsLetter() {
	str := "Lamborghini"
	m := map[int32]int{}
	for _, v := range str {
		_, exsitsKey := m[v]
		if exsitsKey {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		fmt.Printf("%c:%d\n", k, v)
	}
}
