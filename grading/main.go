package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var nextGrades []int32
	for i := 0; i < len(grades); i++ {
		x := grades[i]
		if (x%5) > 2 && x > 37 {
			fmt.Println(x) // should see 73 + 38
			x += 5 - (x % 5)
		}
		nextGrades = append(nextGrades, x)
	}
	return nextGrades
}

func main() {
	fmt.Println(
		gradingStudents([]int32{
			73,
			67,
			38,
			33,
		}),
	)
}
