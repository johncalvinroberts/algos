package main

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivotIndex := len(nums) / 2
	pivot := nums[pivotIndex]

	left := make([]int, 0, len(nums))
	right := make([]int, 0, len(nums))

	for i, v := range nums {
		if i != pivotIndex {
			if v < pivot {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}
	}
	left, right = quickSort(left), quickSort(right)
	left = append(left, pivot)
	left = append(left, right...)
	return left
}

func main() {
	nums := []int{22, 0, 4, 100, 5, 6, 9, 10, 8, 101, 110, 14, 3}
	res := quickSort(nums)
	fmt.Println(res)
}
