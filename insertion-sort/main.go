package main

import "fmt"

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{19, 2, 0, 3, 4, 5, 90, 20, 21, 25, 24, 33, 6}
	res := insertionSort(nums)
	fmt.Println(res)
}
