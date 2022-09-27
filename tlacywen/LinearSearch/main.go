package main

import (
	"fmt"
	"os"
)

func linearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func main() {
	haystack := []int{0, 3, 4, 6, 2, 1, 24, 55, 8, 9, 13, 5}
	needle := 9
	res := linearSearch(haystack, needle)

	if !res {
		fmt.Println("FAILED")
		os.Exit(1)
	}
	fmt.Println("Passed")
	os.Exit(0)
}
