package main

import "fmt"

func merge(left []int, right []int) []int {
	sorted := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			sorted = append(sorted, left[0])
			left = left[1:]
		} else {
			sorted = append(sorted, right[0])
			right = right[1:]
		}
	}
	sorted = append(sorted, left...)
	sorted = append(sorted, right...)
	return sorted
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	middle := len(nums) / 2
	left, right := nums[:middle], nums[middle:]
	return merge(mergeSort(left), mergeSort((right)))
}

func main() {
	nums := []int{9, 0, 4, 5, 10, 2, 44, 23, 3, 21, 6, 99, 7, 11, 10, 999, 998, -1}
	res := mergeSort(nums)
	fmt.Println(res)

}
