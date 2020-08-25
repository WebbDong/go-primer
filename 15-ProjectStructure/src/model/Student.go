package model

import "fmt"

type Student struct {
	Person
	score float64
}

func (s *Student) Study() {
	fmt.Println("Study")
}
