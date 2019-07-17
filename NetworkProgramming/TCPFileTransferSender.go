package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 文件发送端
func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("输入文件路径")
		return
	}

	filePath := args[1]
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// 发送文件名
	conn.Write([]byte(info.Name()))
	resData, err := readServer(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	if "success" == resData {
		sendFile(conn, filePath)
	}
}

// 发送文件
func sendFile(conn net.Conn, filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for {
		buf := make([]byte, 4096)
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送成功")
			} else {
				fmt.Println(err)
			}
			return
		}
		conn.Write(buf[:n])
	}
}

// 读取服务端的返回
func readServer(conn net.Conn) (string, error) {
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	var data string
	if err == nil {
		data = string(buf[:n])
	} else {
		data = ""
	}
	return data, err
}
