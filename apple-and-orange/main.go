package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	var (
		applesOnHouse  int32
		orangesOnHouse int32
	)
	for _, apple := range apples {
		landing := a + apple
		if landing >= s && landing <= t {
			applesOnHouse++
		}
	}
	for _, orange := range oranges {
		landing := b + orange
		if landing >= s && landing <= t {
			orangesOnHouse++
		}
	}
	fmt.Println(applesOnHouse)
	fmt.Println(orangesOnHouse)
}

func main() {

	countApplesAndOranges(7, 11,
		5, 15,
		[]int32{-2, 2, 1},
		[]int32{5, -6},
	)
}
