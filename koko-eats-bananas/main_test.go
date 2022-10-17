package main

import "testing"

type Case struct {
	piles    []int
	h        int
	expected int
}

func TestMinEatingSpeed(t *testing.T) {

	cases := []*Case{

		{
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},

		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},

		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}
	for _, x := range cases {
		res := minEatingSpeed(x.piles, x.h)
		if res != x.expected {
			t.Errorf("Failed. Got %d, expected %d", res, x.expected)
		}
	}
}
