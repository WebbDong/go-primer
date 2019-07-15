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
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	// 返回值：1、读到的字节数，2：客户端的地址，3：error
	count, udpAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("count =", count)
	fmt.Println("udpAddr =", udpAddr)
	fmt.Println("读取到的数据：", string(buf))

	conn.WriteToUDP([]byte("Hello World"), udpAddr)
}
