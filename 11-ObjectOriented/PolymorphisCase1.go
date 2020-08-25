package main

import "fmt"

type Storager interface {
	Write()
	Read()
}

// 移动硬盘
type MobileHDD struct {
}

func (m *MobileHDD) Write() {
	fmt.Println("移动硬盘写入...")
}

func (m *MobileHDD) Read() {
	fmt.Println("移动硬盘读取...")
}

// U盘
type USBFlashDisk struct {
}

func (u *USBFlashDisk) Write() {
	fmt.Println("U盘写入...")
}

func (u *USBFlashDisk) Read() {
	fmt.Println("U盘读取...")
}

// 多态案例1，模拟实现移动硬盘，U盘的读写操作
func main() {
	m1 := MobileHDD{}
	u1 := USBFlashDisk{}
	var storager Storager = &m1
	storager.Read()
	storager.Write()
	fmt.Println("-----------------------")
	storager = &u1
	storager.Read()
	storager.Write()
}
