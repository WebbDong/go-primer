package main

import (
	"fmt"
	"math/big"
)

type ClassA struct {
	id       int
	name     string
	isCan    bool
	p        uintptr
	money    big.Float
	score    float64
	array    [10]int
	classB   ClassB
	objArray [10]ClassB
	slice    []int
}

type ClassB struct {
	a string
}

// 类成员初始化总结
func main() {
	var c1 ClassA
	fmt.Println(c1.objArray)

	c2 := ClassA{}
	fmt.Println(c2)
}
