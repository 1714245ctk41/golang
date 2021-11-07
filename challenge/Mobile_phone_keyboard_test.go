package main

import "testing"

func TestMMobilePhone(t *testing.T) {
	type test struct {
		data   string
		answer string
	}

	tests := []test{
		{"CrCellBax", "Relax"},
		{"CgCoodlBClCuck", "GoodLuck"},
		{"aCaBBCCab", "AB"},
		{"aBB", ""},
		{"adfCfsdafCBDSBBdsfB", "adfFSDAds"},
	}

	for _, v := range tests {
		x := Solution_Mobile(v.data)
		if x != v.answer {
			t.Error(
				"Expected", v.answer, "Got", x,
			)
		}
	}

}
