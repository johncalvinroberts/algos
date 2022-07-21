package main

import "fmt"

func bubble_sort(nums []int) []int {
	for i, _ := range nums {
		for j := 0; j < len(nums)-i-1; j++ {

			if nums[j] > nums[j+1] {
				v := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = v
			}
		}
	}
	return nums
}

func main() {
	nums := []int{1, 5, 0, -1, 9, 10, 200, 300, 2, 30, 40}
	res := bubble_sort(nums)
	fmt.Println(res)
}
