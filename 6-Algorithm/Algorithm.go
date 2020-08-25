package main

import "fmt"

func main() {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//bubbleSort(s)
	selectionSort(s)
	fmt.Println(s)
}

// 冒泡排序
func bubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				//temp := s[j]
				//s[j] = s[j + 1]
				//s[j + 1] = temp
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// 选择排序
func selectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		minIndex := i
		for j := i; j < len(s); j++ {
			if s[minIndex] > s[j] {
				minIndex = j
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}
