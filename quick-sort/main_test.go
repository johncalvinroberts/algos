package main

import "testing"

func TestMain(t *testing.T) {
	nums := []int{22, 0, 4, 100, 5, 6, 9, 10, 8, 101, 110, 14, 3}
	res := quickSort(nums)
	if len(nums) != len(res) {
		t.Error("wrong length")
	}
}
