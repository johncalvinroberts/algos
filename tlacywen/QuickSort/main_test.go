package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	nums := []int{22, 0, 4, 100, 5, 6, 9, 10, 8, 101, 110, 14, 3}
	expected := []int{0, 3, 4, 5, 6, 8, 9, 10, 14, 22, 100, 101, 110}
	res := QuickSort(nums)
	if len(nums) != len(*res) {
		t.Error("wrong length")
	}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Got %v, expected %v", res, expected)
	}
}
