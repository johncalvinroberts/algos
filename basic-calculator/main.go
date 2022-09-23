package main

// https://leetcode.com/problems/basic-calculator-ii
// https://leetcode.com/problems/basic-calculator-ii/discuss/1971812/Javascript-Easy-to-Understand

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func isDigit(s string) bool {
	digitRegEx := regexp.MustCompile(`\d`)
	return digitRegEx.MatchString(s)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Failed to convert string to int")
	}
	return i
}

func calculate(s string) int {
	vals := strings.Split(s, "")
	stack := []int{}
	currentNum := 0
	currentOperator := "+"
	for i := 0; i <= len(vals); i++ {
		var v string
		if i != len(s) {
			v = vals[i]
		}
		if v == " " {
			continue
		}
		fmt.Printf("stack %v, currentNum %d, currentOperator %s\n", stack, currentNum, currentOperator)
		if isDigit(v) {
			currentNum = currentNum*10 + strToInt(v)
		} else {
			switch currentOperator {
			case "+":
				stack = append(stack, currentNum)
			case "-":
				stack = append(stack, -currentNum)
			case "*":
				num := stack[len(stack)-1]
				stack = append(stack[:len(stack)-1], num*currentNum)
			case "/":
				num := stack[len(stack)-1]
				val := math.Trunc(float64(num) / float64(currentNum))
				stack = append(stack[:len(stack)-1], int(val))
			}
			currentNum = 0
			currentOperator = v
		}
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}

func main() {
	cases := []map[string]int{
		{"3+2*2": 7},
		{" 3/2 ": 1},
		{" 3+5 / 2 ": 5},
		{"3+3*10-2": 31},
		{"4*10+2": 42},
	}

	for _, item := range cases {
		for key, value := range item {
			// fmt.Printf("key: %s, val: %d\n", key, value)
			result := calculate(key)
			if result != value {
				fmt.Printf("FAILED. Expected %d, got %d, expr: %s\n", value, result, key)
			} else {
				fmt.Printf("PASSED. Expected %d, got %d, expr: %s\n", value, result, key)
			}
		}
	}

}
