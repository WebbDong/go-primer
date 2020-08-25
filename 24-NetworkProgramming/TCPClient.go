package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for i := 0; i < 10; i++ {
		if i == 9 {
			conn.Write([]byte("quit"))
		} else {
			conn.Write([]byte("你好吗？"))
		}

		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("读取到的数据：", string(buf[:n]))
		time.Sleep(1 * time.Second)
	}
}
