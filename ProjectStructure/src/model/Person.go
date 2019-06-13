package model

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

func (p *Person) Run() {
	fmt.Println("跑步")
}

func (p *Person) Walk() {
	fmt.Println("走路")
}

func (p *Person) Speak() {
	fmt.Printf("My name is %s, age is %d\n", p.name, p.age)
}
