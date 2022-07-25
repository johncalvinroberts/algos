package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	amPm, hrs := s[len(s)-2:], strings.Split(s[:len(s)-2], ":")
	offset := 0
	if amPm == "PM" {
		offset = 1
	}
	ogHrs, err := strconv.Atoi(hrs[0])
	if err != nil {
		panic(err)
	}
	mHrs := ogHrs%12 + offset*12
	return fmt.Sprintf("%02d:%v", mHrs, strings.Join(hrs[1:], ":"))
}

func main() {
	testCase := "07:05:45PM"

	fmt.Println(timeConversion(testCase))
}
