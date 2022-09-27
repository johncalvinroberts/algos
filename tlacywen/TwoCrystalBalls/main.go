package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*

This algorithm is cool, because it uses square root of the input length
to determine how to walk through the input.

It's like binary search, however instead of halving the input everytime, it uses square root of the input
to determine where to look/how to chunk up the input

The runtime complexity is: O√N -- O Square root of n. The square root of the input size.
*/

func twoCrystalBalls(breaks []bool) int {
	chunkSize := int(math.Sqrt(float64(len(breaks))))
	i := chunkSize
	for ; i < len(breaks); i += chunkSize {
		if breaks[i] {
			break
		}
	}
	i -= chunkSize

	for j := 0; j < chunkSize; {
		if breaks[i] {
			return i
		}
		j++
		i++
	}

	return -1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(10000)
	fmt.Println(idx)
	data := make([]bool, 10000)
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	result := twoCrystalBalls(data)
	if result != idx {
		fmt.Println("Failed")
	} else {
		fmt.Println("Succeeded")
	}

}
