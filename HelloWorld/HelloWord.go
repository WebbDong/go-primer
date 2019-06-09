package main

import "fmt"

// 全局变量, 不能使用:=定义全局变量
var global = 1000

func main() {
	fmt.Println("Hello World")
	//variable()
	//outputFormat()
	//input()
	//dataType()
	//typeCast()
	//constant()
	operator()
}

// 变量相关
func variable() {
	// 整型变量默认值为0
	var age int
	fmt.Println(age)

	// 多变量声明
	var num1, num2, num3 int
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)

	// 变量初始化，赋值时自动推断类型
	var num = 200
	fmt.Println(num)

	// 变量赋值
	var num4 int
	num4 = 300
	fmt.Println(num4)

	// 同时打印多个变量
	fmt.Println(num1, num2, num3)

	var num5 int = num
	fmt.Println(num5)

	// 这种写法是错误的，必须指定数据类型
	/*	var num6
		num6 = 3000
		fmt.Println(num6)*/

	// 交换变量
	var a = 100
	var b = 200
	var temp = a
	a = b
	b = temp
	fmt.Println(a, b)

	// 自动推导类型
	n := 2.05
	fmt.Println(n)

	// 多重变量赋值
	n1, n2, n3 := 500, 600, 700
	fmt.Println(n1, n2, n3)
	var nn1, nn2, nn3 = 1000, 2000, 3000
	fmt.Println(nn1, nn2, nn3)

	// 多重赋值方式交换变量
	aa1 := 10
	aa2 := 20
	aa1, aa2 = aa2, aa1
	fmt.Println(aa1, aa2)
}

// 输出格式
func outputFormat() {
	num1 := 100
	num2 := 200
	num3 := 300
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)
	fmt.Println("num1 =", num1, "num2 =", num2, "num3 =", num3)
	fmt.Printf("num1 = %d num2 = %d\n", num1, num2)
	fmt.Printf("num1变量内存地址: %p", &num1)
}

// 输入
func input() {
	var age int
	fmt.Println("输入年龄")
	//fmt.Scanf("%d", &age)
	// Scan，不需要指定数据格式符
	fmt.Scan(&age)
	fmt.Println("age =", age)
	fmt.Println("age变量内存地址:", &age)
	fmt.Printf("age变量内存地址: %p", &age)
}

/*
数据类型
	1、uint8，无符号 8 位整型 (0 到 255)
	2、uint16，无符号 16 位整型 (0 到 65535)
	3、uint32，无符号 32 位整型 (0 到 4294967295)
	4、uint64，无符号 64 位整型 (0 到 18446744073709551615)
	5、int8，有符号 8 位整型 (-128 到 127)
	6、int16，有符号 16 位整型 (-32768 到 32767)
	7、int32，有符号 32 位整型 (-2147483648 到 2147483647)
	8、int64，有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	9、float32，IEEE-754 32位浮点型数,精确到小数点后7位
	10、float64，IEEE-754 64位浮点型数，精确到小数点后15位，浮点数建议使用float64
	11、complex64，32 位实数和虚数
	12、complex128，64 位实数和虚数
	13、byte，类似 uint8，当作字符类型
	14、rune，类似 int32
	15、uint，32 或 64 位
	16、int，与 uint 一样大小
	17、uintptr，无符号整型，用于存放一个指针
	18、bool，布尔类型，true false，默认值为false
*/
func dataType() {
	var num1 int32
	num1 = 2147483647
	fmt.Println("num1 =", num1)

	var num2 float32
	num2 = 900.14567
	fmt.Printf("num2 = %.2f\n", num2)

	// 注意：自动推导类型的浮点数，默认为float64
	// Printf中%T为打印数据类型
	num3 := 110.56
	fmt.Printf("num3 = %.2f, num3 type = %T\n", num3, num3)

	// 布尔类型，默认值为false
	var isResult bool
	isResult = true
	fmt.Println("isResult =", isResult)
	fmt.Printf("isResult = %t\n", isResult)

	// 字符类型
	var ch1 byte = 'a'
	fmt.Printf("ch1 = %c\n", ch1)

	// 字符串类型，字符串类型也有一个'\0'的结束字符
	var str1 string
	str1 = "Ferrari"
	fmt.Println("str1 =", str1)
	fmt.Printf("str1 = %s\n", str1)
	fmt.Printf("str1 length = %d\n", len(str1))
	str2 := "a"
	fmt.Printf("str2 = %s, str2 type = %T\n", str2, str2)
	// 一个汉字占3个字符
	str3 := "法拉利"
	fmt.Printf("str3 = %s, str3 length = %d", str3, len(str3))
}

// 强制类型转换
func typeCast() {
	var num1 float64 = 3.15
	fmt.Println("num1 =", int(num1))
	var num2 int = 20
	fmt.Printf("num2 = %.2f\n", float64(num2))

	var num3 float32 = 2.15
	var num4 float64 = 3.44
	fmt.Printf("num3 + num4 = %.2f\n", float64(num3)+num4)

	var b = true
	// 整型不允许强制转换成布尔类型
	//b = bool(0)
	fmt.Println("b =", b)
}

// 常量
func constant() {
	const PI float64 = 3.14
	const APP_VERSION = 4.56
	fmt.Printf("PI = %.2f", PI)
	fmt.Printf("CONST_VALUE = %.2f\n", APP_VERSION)
	// 常量不能打印地址
	//fmt.Printf("PI = %p", &PI)
}

// 运算符
func operator() {
	var num = 10
	// GO没有前++或--
	//++num
	//--num
	// GO也不允许自增运算符的变量作为赋值的右值
	//var num2 = num++
	fmt.Println("num =", num)

	var global = 3000
	fmt.Println("global =", global)
}
