package main

import (
	"fmt"
	"net"
)

// UDP服务端
func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := net.ListenUDP("udp", addr)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 返回值：1、读到的字节数，2：客户端的地址，3：error
		for {
			buf := make([]byte, 4096)
			count, udpAddr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			data := string(buf)
			fmt.Println("count =", count)
			fmt.Println("udpAddr =", udpAddr)
			fmt.Println("读取到的数据：", data)
			if "exit" == data {
				conn.Close()
			}
			if "quit" == data {
				conn.Close()
				return
			}
			conn.WriteToUDP([]byte("Hello World"), udpAddr)
		}
	}
}
