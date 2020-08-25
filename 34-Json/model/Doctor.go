package model

import "fmt"

type Doctor struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (d *Doctor) PrintAllFields() {
	fmt.Println("name = ", d.Name, ", age = ", d.Age, ", address = ", d.Address)
}
