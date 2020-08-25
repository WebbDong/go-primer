package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	go handleClientConnWrite2(conn, 2)
	go handleClientConnRead2(conn)
	time.Sleep(60 * time.Second)
}

func handleClientConnRead2(conn net.Conn) {
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func handleClientConnWrite2(conn net.Conn, i int) {
	for {
		conn.Write([]byte(fmt.Sprintf("Hello%d", i)))
		time.Sleep(2 * time.Second)
	}
}
