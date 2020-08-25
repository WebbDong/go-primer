package main

import (
	"fmt"
	"net"
)

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

	for {
		buf := make([]byte, 4096)
		n, udpAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := string(buf[:n])
		fmt.Println("读取到的数据：", data)
		if "quit" == data {
			fmt.Println("关闭。。。")
			conn.Close()
			return
		}
		go handleUDPWriteConnect(conn, udpAddr)
	}
}

func handleUDPWriteConnect(conn *net.UDPConn, updAddr *net.UDPAddr) {
	conn.WriteToUDP([]byte("Success"), updAddr)
}
