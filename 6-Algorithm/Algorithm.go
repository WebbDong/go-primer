package main

import "fmt"

func main() {
	s := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	//bubbleSort(s)
	// selectionSort(s)
	mergeSort(s)
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

var temp []int

// 归并排序
func mergeSort(s []int) {
	temp = make([]int, len(s)>>1)
	divide(s, 0, len(s))
}

func divide(s []int, b int, e int) {
	if e-b < 2 {
		return
	}

	m := (b + e) >> 1
	divide(s, b, m)
	divide(s, m, e)
	merge(s, b, m, e)
}

func merge(s []int, b int, m int, e int) {
	le := m - b
	re := e
	li, ri, ai := 0, m, b

	for i := 0; i < le; i++ {
		temp[i] = s[b+i]
	}

	for li < le {
		if ri < re && s[ri] < temp[li] {
			s[ai] = s[ri]
			ri++
		} else {
			s[ai] = temp[li]
			li++
		}
		ai++
	}
}
