package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 第二阶段测试
func main() {
	test1()
}

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
