package main

import "testing"

func TestLovelyNumber(t *testing.T) {
	type test struct {
		a      int
		b      int
		answer int
	}

	tests := []test{
		{0, 0, 1},
		{1, 111, 110},
		{1, 333, 330},
		{1, 999, 990},
		{100000, 100000, 0},
	}

	for _, v := range tests {
		x := Solution_Lovely_Number(v.a, v.b)
		if x != v.answer {
			t.Error(
				"Expected", v.answer, "Got", x,
			)
		}
	}

}
