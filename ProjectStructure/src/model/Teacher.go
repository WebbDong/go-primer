package model

import "fmt"

type Teacher struct {
	Person
}

func (t *Teacher) Teach() {
	fmt.Println("Teach")
}

func CreateTeacher(id int, name string, age int) *Teacher {
	return &Teacher{Person{id, name, age}}
}
