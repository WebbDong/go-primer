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
	go handleClientConnWrite(conn, 1)
	go handleClientConnRead(conn)
	time.Sleep(60 * time.Second)
}

func handleClientConnRead(conn net.Conn) {
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

func handleClientConnWrite(conn net.Conn, i int) {
	for {
		conn.Write([]byte(fmt.Sprintf("Hello%d", i)))
		time.Sleep(5 * time.Second)
	}
}
