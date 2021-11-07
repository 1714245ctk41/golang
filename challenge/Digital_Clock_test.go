package main

import "testing"

func TestDigitalClock(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{1, 8, 3, 2}, 6},
		{[]int{2, 3, 3, 2}, 3},
		{[]int{6, 2, 4, 7}, 0},
		{[]int{0, 0, 0, 0, 0}, 1},
		{[]int{0, 0, 0, 0, 1}, 5},
	}

	for _, v := range tests {
		x := Solution_Digital_clock(v.data)
		if x != v.answer {
			t.Error(
				"Expected", v.answer, "Got", x,
			)
		}
	}

}
