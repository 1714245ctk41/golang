package main

import "testing"

func TestCutTrees(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{3, 4, 5, 8, 7, 7, 6, 5, 3, 7, 8, 9}, -1},
		{[]int{4, 3, 6, 3, 6}, 0},
		{[]int{3, 4, 5, 3, 7}, 3},
		{[]int{1, 2, 3, 4}, -1},
		{[]int{1, 3, 1, 2}, 0},
	}

	for _, v := range tests {
		x := Solution_Trees(v.data)
		if x != v.answer {
			t.Error(
				"Expected", v.answer, "Got", x,
			)
		}
	}

}
