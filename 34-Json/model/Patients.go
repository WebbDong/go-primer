package model

import "fmt"

type Patients struct {
	Name     string   `json:"name"`
	Num      int      `json:"num"`
	Symptoms []string `json:"symptoms"`
	Doctor   Doctor   `json:"doctor"`
}

func (p *Patients) PrintAllFields() {
	fmt.Println("name = ", p.Name, ", num = ", p.Num, ", symptoms = ", p.Symptoms, ", doctor = ", p.Doctor)
}
