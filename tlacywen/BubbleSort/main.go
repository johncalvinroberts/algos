package main

import "fmt"

func bubbleSort(arr []int) []int {
	offset := len(arr) - 1
	for offset > 0 {
		for i := 0; i < offset; i++ {
			point := arr[i]
			right := arr[i+1]
			if right < point {
				arr[i+1] = point
				arr[i] = right
			}
		}
		offset--
	}
	return arr
}

func main() {
	test := []int{1, 4, 10, 2, 6, 7, 3, 5}
	fmt.Println(bubbleSort(test))
}
