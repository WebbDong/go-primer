package main

import "fmt"

type People struct {
}

func (p *People) ShowA() {
	fmt.Println("showA")
	// 调用的始终是People的showA，不管创建的对象是否是子类，与java等其他语言不同
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Professor struct {
	People
}

func (p *Professor) ShowB() {
	fmt.Println("Professor showB")
}

// 多态的一种注意情况，与Java等其他编程语言不同
func main() {
	t := Professor{}
	t.ShowA()
}
