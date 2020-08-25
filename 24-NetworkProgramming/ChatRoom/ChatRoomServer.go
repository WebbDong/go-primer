package main

import (
	"fmt"
	"net"
)

// 客户
type Client struct {
	name       string
	addr       string
	msgChannel chan Message
}

// 消息
type Message struct {
	senderAddr string
	msgContent string
}

// 在线用户map
var onlineClientMap map[string]Client

// 消息channel
var commonMsgChannel chan Message

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
	commonMsgChannel = make(chan Message)
}

// 处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	client := Client{addr, addr, make(chan Message)}
	onlineClientMap[addr] = client

	go handleClientMsgChannel(conn, client.msgChannel)

	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := string(buf[:n])
		msg := Message{client.addr, data}
		fmt.Println(msg)
		commonMsgChannel <- msg
	}
}

// 处理客户消息channel
func handleClientMsgChannel(conn net.Conn, msgChannel chan Message) {
	for {
		msg := <-msgChannel
		content := msg.senderAddr + ":" + msg.msgContent
		conn.Write([]byte(content))
	}
}

// 处理公共消息channel和广播消息
func manager() {
	for {
		msg := <-commonMsgChannel
		for _, v := range onlineClientMap {
			if msg.senderAddr != v.addr {
				v.msgChannel <- msg
			}
		}
	}
}
