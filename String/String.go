package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串处理
func main() {
	//contains()
	//join()
	//index()
	//repeat()
	//replace()
	//split()
	stringConvert()
}

// 1、Contains方法，判断第一个参数的字符串是否包含第二个参数的字符串
func contains() {
	str1 := "Ferrari"
	fmt.Println(strings.Contains(str1, "ari"))
	fmt.Println(strings.Contains(str1, "abc"))
}

// 2、Join方法
func join() {
	s1 := []string{"str1", "str2", "str3", "str4", "str5"}
	str1 := strings.Join(s1, "")
	fmt.Println(str1)
	fmt.Println(strings.Join(s1, ","))
}

// 3、Index方法，查找某个子字符串的位置，找到返回所在位置，未找到返回-1
func index() {
	pos := strings.Index("Lamborghini", "Lam")
	fmt.Println(pos)
	fmt.Println(strings.Index("Ferrari", "Lam"))
}

// 4、Repeat方法，返回指定的字符串重复指定的次数
func repeat() {
	str1 := "Lamborghini"
	fmt.Println(strings.Repeat(str1, 3))
}

// 5、Replace方法，字符串替换
func replace() {
	str1 := "Lamborghini Lamborghini Lamborghini"
	// 第四个参数，是指定要替换多少个字符串，如果指定的是负数，表示替换所有匹配的字符串
	fmt.Println(strings.Replace(str1, "Lam", "Fer", -1))
	fmt.Println(strings.Replace(str1, "Lam", "Fer", 2))
}

// 6、Split方法，分割字符串
func split() {
	str1 := "Lamborghiin:Ferrari:Pagani"
	strs := strings.Split(str1, ":")
	fmt.Println(strs)
	for i, v := range strs {
		fmt.Printf("strs[%d] = %s\n", i, v)
	}
}

// 7、字符串转换
func stringConvert() {
	// 将其他类型转换成字符串
	// 第二个参数指定进制数
	strInt := strconv.FormatInt(400, 10)
	fmt.Println(strInt)

	// int转string
	strI := strconv.Itoa(400)
	fmt.Println(strI)

	// bool转string
	fmt.Println(strconv.FormatBool(true))

	// 浮点数转string
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// 		'b' (-ddddp±ddd，二进制指数)
	// 		'e' (-d.dddde±dd，十进制指数)
	// 		'E' (-d.ddddE±dd，十进制指数)
	// 		'f' (-ddd.dddd，没有指数)
	// 		'g' ('e':大指数，'f':其它情况)
	// 		'G' ('E':大指数，'f':其它情况)
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	fmt.Println(strconv.FormatFloat(29.547, 'f', 2, 64))
	fmt.Println("-------------------------------------------------")

	// 把字符串转换成其他类型
	str1 := "true"
	b, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	fmt.Println(b)

	str2 := "900"
	i, err := strconv.ParseInt(str2, 10, 32)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	fmt.Println(i)

	str3 := "800.876"
	f, err := strconv.ParseFloat(str3, 64)
	if err != nil {
		fmt.Println("转换失败", err)
	}
	fmt.Println(f)
}
