package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 文件接收端
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
		go receiveFile(conn)
	}
}

func receiveFile(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fileName := string(buf[:n])
		conn.Write([]byte("success"))

		filePath := "D:\\360Downloads\\" + fileName
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		for {
			buf := make([]byte, 4096)
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					fmt.Println("文件传输成功")
					file.Close()
				} else {
					fmt.Println(err)
					file.Close()
				}
				return
			}
			file.Write(buf[:n])
		}
	}
}
