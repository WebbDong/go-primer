package main

import (
	"./model"
	"fmt"
)

// ide默认运行时，只会编辑当前的go文件，如果要编译整个src目录下的go文件，需要进行配置
func main() {
	fmt.Println(Add(10, 20))
	p1 := model.Person{}
	fmt.Println(p1)
}
