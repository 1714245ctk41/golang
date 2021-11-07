package main

import "testing"

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{2, 4, 7}, 13},
		{[]int{1, 1}, 2},
		{[]int{3, -1, 0}, 2},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error(
				"Expected", v.answer, "Got", x,
			)
		}
	}

}
