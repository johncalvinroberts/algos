package main

import (
	"fmt"
	"math"
)

func nonDivisibleSubset(k int32, s []int32) int32 {
	// build a map
	counts := make(map[int32]int32)
	for _, v := range s {
		counts[v%k] += 1
	}
	count := math.Max(float64(counts[0]), 1)

	for i := int32(1); i < k; i++ {
		if i != k-i {
			toAdd := math.Max(float64(counts[i]), float64(counts[k-i]))
			count += float64(toAdd)
		}
	}
	if k%2 == 0 {
		count += 1
	}
	return int32(count)
}

func main() {

	testCase := []int32{19, 10, 12, 10, 24, 25, 22}
	testDivisor := int32(4)

	/*
		possible solutions are:
		[10,12,25]
		[19,22,24]

		expect 3
	*/
	fmt.Println(nonDivisibleSubset(testDivisor, testCase))
}
