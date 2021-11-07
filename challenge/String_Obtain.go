package main

import (
	"strconv"
	"strings"
)

func Solution_String_Obtain(S string, T string) string {
	cases := []string{"ADD", "CHANGE", "MOVE", "NOTHING", "IMPOSSIBLE"}
	compare := strings.Compare(S, T)
	result := ""
	if compare == 0 {
		return cases[3]
	}
	indeff := []string{"0", "", ""}
	if len(S) == len(T) {
		for i, v := range S {
			if v != rune(T[i]) {
				indeffNum, _ := strconv.Atoi(indeff[0])
				indeff[0] = strconv.Itoa(indeffNum + 1)
				indeff[1] = string(rune(v))
				indeff[2] = string(rune(T[i]))
			}
		}
		if indeff[0] == "1" {
			result = cases[1] + " " + indeff[1] + " " + indeff[2]
			return result
		}
		checkcontain := true
		for _, v := range T {
			if !strings.Contains(S, string(v)) {
				checkcontain = false
				break
			}
		}
		if checkcontain {
			sRune := []rune(S)
			for i := 0; i < len(sRune)-1; i++ {
				for j := i + 1; j < len(sRune); j++ {
					com := make([]string, len(sRune))
					for k, v := range sRune {
						if k == i {
							continue
						}
						com = append(com, string(v))
						if k == j {
							com = append(com, string(sRune[i]))
						}
					}
					comS := strings.Join(com, "")
					if strings.Compare(comS, T) == 0 {
						result = cases[2] + " " + string(sRune[i])
						return result
					}

				}
			}
		}
		return cases[4]
	}
	if len(T)-len(S) == 1 {
		for i, v := range S {
			if strings.Compare(string(T[i]), string(v)) != 0 {
				return cases[4]
			}
		}
		return cases[0] + " " + string(rune(T[len(T)-1]))
	}

	return cases[4]
}
