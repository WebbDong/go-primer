package main

import "fmt"

// 面向对象，接口，多态
func main() {
	//interfaces()
	polymorphis()
}

// 接口
type Personer interface {
	Walk()
	Run()
	Speak()
}

type Human struct {
	height float64
	weight float64
	age    int
	name   string
}

func (h *Human) Walk() {
	fmt.Println("走路")
}

func (h *Human) Run() {
	fmt.Printf("跑步")
}

func (h *Human) Speak() {
	fmt.Printf("My name is %s, age is %d, height is %f, weight is %f\n", h.name, h.age, h.height, h.weight)
}

type RacingDriver struct {
	Human
	speed int
}

type SoftwareEngineer struct {
	Human
	progLang string
}

// 1、接口
func interfaces() {
	var se1 SoftwareEngineer
	se1.name = "Kobe"
	se1.Speak()

	rd1 := RacingDriver{Human{name: "舒马赫", age: 38}, 350}

	// 实现了接口中的某一个方法时，必须同时实现全部的接口方法，才能赋值给对应的接口变量
	// 接口变量相当于指针变量，赋值时需要使用取地址符
	var per1 Personer = &se1
	per1.Speak()
	per1 = &rd1
	per1.Speak()
}

func polymorphisTest(p Personer) {
	p.Speak()
}

// 2、多态
func polymorphis() {
	rd1 := RacingDriver{Human{name: "舒马赫", age: 38}, 380}
	var per Personer = &SoftwareEngineer{Human{name: "欧文", age: 28}, "Go"}
	per.Speak()
	per = &rd1
	per.Speak()
	fmt.Println("--------------------------------------------------------")

	polymorphisTest(per)
	rd2 := RacingDriver{Human{name: "科比", age: 30}, 400}
	polymorphisTest(&rd2)
}
