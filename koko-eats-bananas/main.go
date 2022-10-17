package main

import (
	"fmt"
	"math"
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	min := 1
	sort.Ints(piles)
	max := piles[len(piles)-1]
	fmt.Printf("max %d\n", max)
	ans := max
	for min <= max {
		mid := min + int(math.Floor((float64(max)-float64(min))/2))
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(mid)))
		}
		if hours <= h {
			ans = int(math.Min(float64(mid), float64(ans)))
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return ans
}

func main() {
	piles := []int{3, 6, 7, 11}
	hours := 8
	expected := 4
	res := minEatingSpeed(piles, hours)
	fmt.Printf("Got %d, expected %d", res, expected)
}
