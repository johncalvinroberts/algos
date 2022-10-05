package main

import "fmt"

func partition(arr *[]int, lo int, hi int) int {
	cloned := *arr
	pivotVal := cloned[hi]
	fmt.Printf("Pivotval %d\n", pivotVal)
	idx := lo
	for i := lo; i < hi; i++ {
		if pivotVal <= cloned[i] {
			idx++
			tmp := cloned[i]
			fmt.Printf("Swapping %d, %d\n", tmp, cloned[idx])
			cloned[i] = cloned[idx]
			cloned[idx] = tmp
		}
	}
	fmt.Println(cloned)
	*arr = cloned
	return idx
}

func qs(arr *[]int, lo int, hi int) *[]int {
	if lo >= hi {
		return arr
	}
	pivotIdx := partition(arr, lo, hi)
	fmt.Println(pivotIdx)
	// fmt.Println(arr)
	// *arr = *qs(arr, lo, pivotIdx-1)
	// *arr = *qs(arr, pivotIdx+1, hi)
	return arr
}

func QuickSort(arr []int) *[]int {
	lo := 0
	hi := len(arr) - 1
	ret := qs(&arr, lo, hi)
	return ret
}
