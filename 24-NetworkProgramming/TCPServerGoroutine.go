package main

import (
	"fmt"
	"io"
	"net"
)

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
			continue
		}
		fmt.Println("客户端连接成功。。。")

		go handleTCPConnect(conn)
	}
}

func handleTCPConnect(conn net.Conn) {
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		data := string(buf[:n])
		if err == io.EOF || "quit" == data {
			fmt.Println("客户端断开连接。。。")
			conn.Write([]byte("Closed"))
			conn.Close()
			break
		}
		fmt.Println("读取到数据：", data)

		conn.Write([]byte("Success"))
	}
}
