package main

import "fmt"

type Point struct {
	x int
	y int
}

const (
	WALL  = "x"
	END   = "E"
	START = "S"
)

func Walk(puzzle [][]string, curr *Point, end *Point, seen *map[Point]bool, path *[]*Point) bool {
	seenRef := *seen
	fmt.Printf("curr %v\n", curr)
	// BASE CASE: how do we know if we stop recursing?

	// 1. we've seen curr already
	if seenRef[*curr] {
		// WE HAVE BEEN HERE BEFORE
		return false
	}
	// mark as seen
	seenRef[*curr] = true
	*seen = seenRef

	// 2. curr is off the map
	if curr.x > len(puzzle[0]) || curr.y > len(puzzle) {
		return false
	}
	// 3. curr is a wall
	if puzzle[curr.y][curr.x] == WALL {
		return false
	}
	// 4. curr is end
	if curr.x == end.x && curr.y == end.y {
		// WE HIT THE END, add to the path and stop recursing
		*path = append(*path, curr)
		return false
	}

	// if we get here, we've reached our recurse case!
	// add the current point to the path
	*path = append(*path, curr)

	// use this for looping through all sides of the current point
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	// loop through curr's neighbors
	for _, dir := range directions {
		x, y := dir[0], dir[1]
		p := &Point{x: curr.x + x, y: curr.y + y}
		if Walk(puzzle, p, end, seen, path) {
			return true
		}
	}
	return true
}

func Solve(puzzle [][]string, start *Point, end *Point) []*Point {
	path := []*Point{}
	seen := map[Point]bool{}
	Walk(puzzle, start, end, &seen, &path)
	return path
}
