package main

import "testing"

func TestStringObtain(t *testing.T) {
	type test struct {
		dataS string
		dataT string

		answer string
	}

	tests := []test{
		{"nice", "nicer", "ADD r"},
		{"nicse", "nicers", "IMPOSSIBLE"},
		{"nicer", "nicers", "ADD s"},

		{"test", "tent", "CHANGE s n"},
		{"ghtraadsf", "ghjraadsf", "CHANGE t j"},
		{"beans", "banes", "MOVE e"},
		{"asdf1234", "asdf12345", "ADD 5"},
		{"0", "odd", "IMPOSSIBLE"},
		{"0sd", "odd", "IMPOSSIBLE"},

		{"sdd", "odd", "CHANGE s o"},

		{"abc", "abc", "NOTHING"},
		{"abc123", "abc123", "NOTHING"},

		{"bseans", "beanss", "MOVE s"},
		{"mmasdfm", "masdfmm", "MOVE m"},
		{"mmmasdfm2", "mmasdfmm2", "MOVE m"},
		{"mmmasdfm", "mmasdfmm2", "IMPOSSIBLE"},
	}

	for _, v := range tests {
		x := Solution_String_Obtain(v.dataS, v.dataT)
		if x != v.answer {
			t.Error(
				"Expected", v.answer, "Got", x,
			)
		}
	}

}
