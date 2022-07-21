package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	n, i, m := len(nums)-1, 0, 0

	for i <= n {
		m = i + (n-i)/2

		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			i = m + 1
		}
		if nums[m] > target {
			i = m - 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 54, 66}
	res := search(nums, 54)
	fmt.Println(res)
}
