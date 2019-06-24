package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 文本文件操作
// https://www.jianshu.com/p/9fcfad300e46
func main() {
	//createFile()
	//writeStringFunc()
	//writeFunc()
	//writeAtFunc()
	//seekFunc()
	//appendDataToFile()
	//readFile()
	//readLine()
	//readLine2()
	//readDir()
	readDirRecursion("D:\\logo")
}

// 1、文件创建
func createFile() {
	// os.Create函数返回文件指针和错误信息
	file, err := os.Create("D:\\code\\a.txt")
	// 关闭文件
	defer file.Close()
	if err != nil {
		fmt.Println("文件创建失败，", err)
	} else {
		fmt.Println(file)
	}
}

// 2、写入数据，WriteString函数
func writeStringFunc() {
	file, err := os.Create("D:\\code\\a.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("文件创建失败，", err)
	}
	// 返回数据长度与错误信息
	n, err := file.WriteString("Ferrari")
	fmt.Println("n =", n)
	if err != nil {
		fmt.Println("数据写入错误，", err)
	}
}

// 3、写入数据，Write函数
func writeFunc() {
	file, err := os.Create("D:\\code\\b.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("文件创建失败，", err)
	}
	// 返回数据长度与错误信息
	str := "Pagani"
	n, err := file.Write([]byte(str))
	fmt.Println("n =", n)
	if err != nil {
		fmt.Println("数据写入错误，", err)
	}
}

// 4、写入数据，WriteAt函数
func writeAtFunc() {
	file, err := os.Create("D:\\code\\c.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("数据写入错误，", err)
	}
	n, err2 := file.WriteString("Lamborghini")
	fmt.Println(n)
	if err2 != nil {
		fmt.Println("数据写入错误，", err2)
	}
	str1 := "Ferrari"
	// 在指定位置写入数据
	file.WriteAt([]byte(str1), int64(n))
}

// 5、Seek函数，指定文件指针位置
func seekFunc() {
	file, err := os.Create("D:\\code\\d.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("数据写入错误，", err)
	}
	file.WriteString("Lamborghini")
	file.WriteString("Ferrari")

	// 第一个参数为偏移量，第二个参数为起始位置
	// io.SeekStart：文件起始位置
	// io.SeekCurrent：文件指针当前位置
	// io.SeekEnd：文件末尾位置
	file.Seek(6, io.SeekStart)
	file.WriteString("Ritch Man")

	file.Seek(10, io.SeekEnd)
	file.WriteString("Pagani Zonda")
}

// 6、追加数据到已有文件
func appendDataToFile() {
	/*
		打开文件
		1、第一个参数为文件路径
		2、第二个参数为模式
			os.O_RDONLY     只读
			os.O_WRONLY     只写
			os.O_RDWR       读写文件
			os.O_APPEND     追加文件
			os.O_CREATE     不存在时创建文件
			os.O_TRUNC      打开时截断文件

			如果需要指定多个，使用|
		3、第三个参数为权限
			const (
				// The single letters are the abbreviations
				// used by the String method's formatting.
				ModeDir        FileMode    = 1 << (32 - 1 - iota)    // d: is a directory 文件夹模式
				ModeAppend                        // a: append-only 追加模式
				ModeExclusive                        // l: exclusive use 单独使用
				ModeTemporary                        // T: temporary file (not backed up) 临时文件
				ModeSymlink                        // L: symbolic link 象征性的关联
				ModeDevice                        // D: device file 设备文件
				ModeNamedPipe                        // p: named pipe (FIFO) 命名管道
				ModeSocket                        // S: Unix domain socket  Unix 主机 socket
				ModeSetuid                        // u: setuid 设置uid
				ModeSetgid                        // g: setgid 设置gid
				ModeCharDevice                        // c: Unix character device, when ModeDevice is set Unix 字符设备,当设备模式是设置Unix
				ModeSticky                        // t: sticky 粘性的
				// Mask for the type bits. For regular files, none will be set. bit位遮盖.不变的文件设置为none
				ModeType    = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
				ModePerm    FileMode    = 0777    // Unix permission bits 权限位.
			)
			或者：
				0：没有任何权限
				1：执行权限(如果是可执行文件)
				2：写权限
				3：写权限与执行权限
				4：读权限
				5：读权限与执行权限
				6：读权限与写权限
				7：读权限，写权限，执行权限
	*/
	file, err := os.OpenFile("D:\\code\\d.txt", os.O_APPEND, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败，", err)
		return
	}
	_, err2 := file.WriteString("追加内容")
	if err2 != nil {
		fmt.Println("数据追加失败，", err)
	}
}

// 7、读取文件
func readFile() {
	// Open与OpenFile的区别是，Open内部调用的是OpenFile，并且模式设置为O_RDONLY，权限设置为0
	file, err := os.Open("D:\\code\\d.txt")
	if err != nil {
		fmt.Println("文件打开失败，", err)
		return
	}
	defer file.Close()
	b := make([]byte, 1024*2)
	n, err := file.Read(b)
	fmt.Println(n)
	if err != nil {
		fmt.Println("读取失败，", err)
	}
	fmt.Println(string(b[:n]))
	fmt.Println("---------------------------")

	file.Seek(0, io.SeekStart)
	// 循环读取
	b = make([]byte, 5)
	for {
		_, err2 := file.Read(b)
		if err2 == io.EOF {
			break
		}
		fmt.Print(string(b))
	}
}

// 8、读取行
func readLine() {
	file, err := os.OpenFile("D:\\a.txt", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 创建一个带有缓冲区的Reader
	reader := bufio.NewReader(file)
	for {
		bytes, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(string(bytes))
	}
}

func readLine2() {
	file, err := os.OpenFile("D:\\a.txt", os.O_RDONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		line, isPrefix, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(line), isPrefix)
	}
}

// 9、读取目录
func readDir() {
	path, err := os.OpenFile("D:\\", os.O_RDONLY, os.ModePerm)
	defer path.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	// -1：获取所有目录项
	infos, err := path.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, info := range infos {
		fmt.Println(info.Name(), info.IsDir())
	}
}

// 10、递归读取目录
func readDirRecursion(pathStr string) {
	path, err := os.OpenFile(pathStr, os.O_RDONLY, os.ModePerm)
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
	for _, info := range infos {
		if info.IsDir() {
			readDirRecursion(pathStr + "\\" + info.Name())
		}
		fmt.Println(info.Name())
	}
}
