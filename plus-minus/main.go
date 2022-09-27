package main

import "fmt"

func plusMinus(arr []int32) {
	counts := map[string]int{"positive": 0, "negative": 0, "zero": 0}
	for _, v := range arr {
		if v == 0 {
			counts["zero"]++
		}
		if v > 0 {
			counts["positive"]++
		}
		if v < 0 {
			counts["negative"]++
		}
	}
	positive := float64(counts["positive"]) / float64(len(arr))
	negative := float64(counts["negative"]) / float64(len(arr))
	zero := float64(counts["zero"]) / float64(len(arr))
	fmt.Println(fmt.Sprintf("%f", positive))
	fmt.Println(fmt.Sprintf("%f", negative))
	fmt.Println(fmt.Sprintf("%f", zero))
}

func main() {
	tests := []int32{1, 1, 0, -1, -1}
	plusMinus(tests)
}
