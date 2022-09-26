package main

import (
	"fmt"
)

func binarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)
	for lo < hi {
		fmt.Println(lo, hi)
		mid := lo + (hi-lo)/2
		v := haystack[mid]

		if v == needle {
			// if midpoint IS the value, return true
			return true
		} else if v > needle {
			// if midpoint is greater than the value we're searching for, take the left half
			// set the hi to the middlepoint
			hi = mid
		} else {
			// if midpoint is LESS THAN the value we're searching for, take the right half
			// set the lo index to mid + 1
			lo = mid + 1
		}
	}
	return false
}

func main() {
	test := []int{2, 4, 6, 77, 89, 909, 1011, 1012, 1013, 1045, 1066, 1099}
	result := binarySearch(test, 89)
	if !result {
		fmt.Println("Failed")
	} else {
		fmt.Println("Succeeded")
	}
}
