package main

import "fmt"

type Numerator interface {
	calc() int
}

type Compute struct {
	num1 int
	num2 int
}

type Summator struct {
	Compute
}

func (s *Summator) calc() int {
	return s.num1 + s.num2
}

type Subtracter struct {
	Compute
}

func (s *Subtracter) calc() int {
	return s.num1 - s.num2
}

type Calculator struct {
	summator   Summator
	subtracter Subtracter
}

func (c *Calculator) calc(num1 int, num2 int, operate string) int {
	var numberator Numerator
	switch operate {
	case "+":
		c.summator.num1 = num1
		c.summator.num2 = num2
		numberator = &c.summator
	case "-":
		c.subtracter.num1 = num1
		c.subtracter.num2 = num2
		numberator = &c.subtracter
	default:
		return 0
	}
	return numberator.calc()
}

// 案例2，计算器
func main() {
	c := Calculator{}
	fmt.Println(c.calc(20, 30, "+"))
	fmt.Println(c.calc(20, 30, "-"))
}
