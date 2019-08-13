package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// 第二阶段测试
func main() {
	//test1()
	test2()
}

// ------------------------------------------------------ start ----------------------------------------------------
const BASE_PATH = "C:\\a\\"
const OUTPUT_PATH = "C:\\b\\"

func test1() {
	path, err := os.OpenFile(BASE_PATH, os.O_RDONLY, os.ModePerm)
	defer path.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	infos, err := path.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	createDir(OUTPUT_PATH)
	for i, len := 0, len(infos); i < len; i++ {
		info := infos[i]
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".mp3") {
			filePath := BASE_PATH + info.Name()
			totalData, err := readFile(filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			newFilePath := OUTPUT_PATH + fmt.Sprintf("%d.mp3", i+1)
			writeFile(newFilePath, totalData)
		}
	}
	fmt.Println("done...")
}

func createDir(dir string) {
	exists, err := isDirExists(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !exists {
		os.Mkdir(dir, os.ModePerm)
	}
}

func isDirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func writeFile(filePath string, data []byte) {
	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Write(data)
}

func readFile(filePath string) (totalData []byte, err error) {
	file, e := os.Open(filePath)
	defer file.Close()
	if e != nil {
		err = e
		return
	}
	buf := make([]byte, 4096)
	for {
		n, e := file.Read(buf)
		if e != nil {
			if e == io.EOF {
				err = nil
				break
			} else {
				err = e
				return
			}
		}
		totalData = append(totalData, buf[:n]...)
	}
	return
}

// ------------------------------------------------------ end ----------------------------------------------------

// ------------------------------------------------------ start ----------------------------------------------------
func test2() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	defer listener.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
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
		datas, err := readStringFromConn(conn)
		if err != nil {
			fmt.Println(err)
			return
		}
		str := string(datas)
		if str == "quit" {
			conn.Close()
			return
		}
		fmt.Println("读取到：" + str)
		conn.Write([]byte(reverse(str)))
	}
}

func readStringFromConn(conn net.Conn) ([]byte, error) {
	const BUF_SIZE = 4096
	buf := make([]byte, BUF_SIZE)
	n, e := conn.Read(buf)
	if e != nil && e != io.EOF {
		return nil, e
	}
	return buf[:n], nil
}

func reverse(str string) string {
	splits := strings.Split(str, " ")
	var newStr string
	for i := len(splits) - 1; i >= 0; i-- {
		newStr += splits[i]
		if i != 0 {
			newStr += " "
		}
	}
	return newStr
}

func reverseString(str string) string {
	runes := []rune(str)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

// ------------------------------------------------------ end ----------------------------------------------------
