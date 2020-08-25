package main

import (
	"fmt"
	"io"
	"net"
)

// TCP服务端
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("客户端连接成功。。。")

		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err == io.EOF {
				fmt.Println("客户端断开连接。。。")
				conn.Close()
				break
			}
			fmt.Println("读取到数据：", string(buf[:n]))

			conn.Write([]byte("Success"))
		}
	}
}
