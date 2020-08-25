package main

import (
	"fmt"
	"os"
)

// 命令行参数
func main() {
	// args[0]是当前应用exe的路径
	args := os.Args
	for i := 0; i < len(args); i++ {
		fmt.Println(os.Args[i])
	}
}
