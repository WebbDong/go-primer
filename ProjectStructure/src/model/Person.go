package model

import "fmt"

type Person struct {
	// 字段也是同样的规则，开头小写的是私有字段。不能跨包，只能在同一个包中被访问
	// 开头大写的是公有字段，能跨包访问
	id   int
	name string
	age  int
}

// 类的成员方法也遵循开头小写为私有方法，大写为公有方法
func (p *Person) Run() {
	func1()
	fmt.Println("跑步")
}

func (p *Person) Walk() {
	fmt.Println("走路")
}

func (p *Person) Speak() {
	fmt.Printf("My name is %s, age is %d\n", p.name, p.age)
}

func (p *Person) GetName() string {
	return p.name
}

func CreatePerson(id int, name string, age int) *Person {
	return &Person{id, name, age}
}
