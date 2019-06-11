package main

import (
	"fmt"
)

// 面向对象，接口，多态
func main() {
	//interfaces()
	//polymorphis()
	//interfaceExtends()
	//emptyInterface()
	typeAssert()
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

type Interface1 interface {
	Walk()
}

type Interface2 interface {
	Interface1
	Eat()
}

type Class1 struct {
}

func (c *Class1) Walk() {
	fmt.Println("Walk...")
}

func (c *Class1) Eat() {
	fmt.Println("Eat...")
}

// 3、接口继承
func interfaceExtends() {
	c1 := Class1{}

	var i1 Interface1 = &c1
	var i2 Interface2 = &c1

	i1.Walk()
	i2.Walk()
	i2.Eat()

	i1 = i2
	i1.Walk()
}

type Bike struct {
	speed int
}

func (b *Bike) ride() {
	fmt.Println("ride...")
}

// 4、空接口(interface{})，不包含任何方法的接口，所有类型都实现了空接口，所以空接口可以存储任意类型的值
func emptyInterface() {
	var ei1 interface{}
	ei1 = "ferrari"
	fmt.Println(ei1)
	ei1 = 100
	fmt.Println(ei1)

	ei1 = Bike{30}
	fmt.Println(ei1)

	ei1 = &Bike{40}
	fmt.Println(ei1)
	fmt.Println("---------------")

	// 空接口类型切片，可以添加任意类型的数据
	var s1 []interface{}
	s1 = append(s1, 1, 2, 3)
	s1 = append(s1, 5, "ferrari", 38.5, &Bike{})
	fmt.Println(s1)
}

// 5、类型断言
func typeAssert() {
	var i interface{}
	i = 100

	// 判断空接口变量是否是指定的数据类型，如果判断成功，ok为true，并且value为数据值，否则ok为false，
	value, ok := i.(int)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("类型错误")
		fmt.Println(value)
	}
	fmt.Println("-------------------")

	i = Bike{}
	value2, ok2 := i.(Bike)
	if ok2 {
		fmt.Println(value2)
	} else {
		fmt.Println("类型错误")
		fmt.Println(value2)
	}

	i = &Bike{}
	value3, ok3 := i.(*Bike)
	if ok3 {
		fmt.Println(value3)
	} else {
		fmt.Println("类型错误")
		fmt.Println(value3)
	}
}
