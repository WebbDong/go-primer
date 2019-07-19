package main

import (
	"fmt"
	"net"
)

type Client struct {
	name       string
	addr       string
	msgChannel chan string
}

// 在线用户map
var onlineClientMap map[string]Client

// 消息channel
var commonMsgChannel chan string

// 聊天室服务器
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	go manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// 初始化
func init() {
	onlineClientMap = make(map[string]Client)
	commonMsgChannel = make(chan string)
}

// 处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	client := Client{addr, addr, make(chan string)}
	onlineClientMap[addr] = client
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf[:n]))
	}
}

// 处理公共消息channel和广播消息
func manager() {
	for {
		msg := <-commonMsgChannel
		for _, v := range onlineClientMap {
			v.msgChannel <- msg
		}
	}
}
