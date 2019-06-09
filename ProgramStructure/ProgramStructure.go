package main

import "fmt"

func main() {
	//ifAndSwitch()
	for_()
}

func ifAndSwitch() {
	var age = 30
	if age == 30 {

	} else {

	}

	switch age {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	case 40:
		fmt.Println("40")
	default:
		fmt.Println("default")
	}

	switch age - 10 {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	case 40:
		fmt.Println("40")
	default:
		fmt.Println("default")
	}

	fmt.Println("输入分数：")
	var score int
	fmt.Scan(&score)
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}
}

func for_() {
	var index int
	for index = 0; index < 10; index++ {

	}

	for i := 1; i < 20; i++ {
		if i == 5 {
			continue
		}

		if i == 10 {
			break
		}
	}

	var j = 10
	for j >= 0 {
		fmt.Println(j)
		j--
	}

	/*
		for {
			fmt.Println("死循环1")
		}

		for true {
			fmt.Println("死循环2")
		}
	*/

	for i := 0; true; i++ {
		if i == 100 {
			fmt.Println("i =", i)
			goto BREAK_LOOP
		}
	}
BREAK_LOOP:
	fmt.Println("BREAK_LOOP")

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range arr {
		fmt.Printf("arr[%d] = %d\n", i, v)
	}

	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}
}
