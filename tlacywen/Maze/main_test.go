package main

import (
	"fmt"
	"reflect"
	"testing"
)

func drawPath(data [][]string, path []*Point) []string {
	for _, p := range path {
		if data[p.y] != nil && data[p.y][p.x] != "" {
			data[p.y][p.x] = "*"
		}
	}
	res := []string{}
	for _, d := range data {
		s := "\n"
		for _, x := range d {
			s += x
		}
		res = append(res, s)
	}
	return res
}

func TestSolve(t *testing.T) {
	fixture := [][]string{
		{"x", "x", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", " ", " ", " ", "x", " ", "x"},
		{"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", " ", "x"},
		{"x", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", "x"},
		{"x", " ", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x"},
	}
	expected := []*Point{
		{x: 10, y: 0},
		{x: 10, y: 1},
		{x: 10, y: 2},
		{x: 10, y: 3},
		{x: 10, y: 4},
		{x: 9, y: 4},
		{x: 8, y: 4},
		{x: 7, y: 4},
		{x: 6, y: 4},
		{x: 5, y: 4},
		{x: 4, y: 4},
		{x: 3, y: 4},
		{x: 2, y: 4},
		{x: 1, y: 4},
	}
	end := &Point{x: 10, y: 0}
	start := &Point{x: 1, y: 5}
	result := Solve(fixture, start, end)
	resultDrawn := drawPath(fixture, result)
	expectedDrawn := drawPath(fixture, expected)
	fmt.Printf("Results: \n algo: \n %v \n\n expected:\n %v\n", resultDrawn, expectedDrawn)
	if !reflect.DeepEqual(resultDrawn, expectedDrawn) {
		t.Fail()
	}
}
