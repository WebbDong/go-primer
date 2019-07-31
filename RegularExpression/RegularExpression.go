package main

import (
	"fmt"
	"regexp"
)

// 正则表达式
func main() {
	str1 := "abc aaa bbb adc cba"
	compile1 := regexp.MustCompile(`a.c`)
	// 第二个参数是寻找多少个匹配的，-1代表全部
	submatch1 := compile1.FindAllStringSubmatch(str1, -1)
	for i, len := 0, len(submatch1); i < len; i++ {
		fmt.Printf("%s\n", submatch1[i][0])
	}
	fmt.Println("--------------------------------")

	str2 := "<div>Hello World</div>"
	compile2 := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	submatch2 := compile2.FindAllStringSubmatch(str2, -1)
	for _, match := range submatch2 {
		fmt.Println(match[0], match[1])
	}
}
