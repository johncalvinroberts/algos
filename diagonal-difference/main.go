package main

import (
	"fmt"
	"math"
)

// https://www.hackerrank.com/challenges/diagonal-difference/problem

func diagonalDifference(arr [][]int32) int32 {
	var a int32
	var b int32
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
		a += arr[i][i]
		b += arr[i][len(arr)-1-i]
	}
	fmt.Println(a, b)
	return int32(math.Abs(float64(a) - float64(b)))
}

func main() {
	test := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	result := diagonalDifference(test)
	if result != 15 {
		fmt.Printf("failed")
	} else {
		fmt.Printf("success")
	}
}
