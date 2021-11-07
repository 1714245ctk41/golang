package main

import (
	// "fmt"
	// "regexp"
	"strings"
)

func Solution_Mobile(S string) string {
	sAr := strings.Split(S, "")
	capl := false
	for i, v := range sAr {
		if capl{
				sAr[i] = strings.ToUpper(v)
			}
		if v == "C" {
			sAr[i] = "_"
			capl = !capl
		}
		if v == "B"{
			sAr[i] = "_"
			for j:= i - 1; j > -1; j--{
				if sAr[j] != "_" {
					sAr[j] = "_"
					break
				}
			}
		}
	}
	S = strings.Join(sAr, "")
	S = strings.ReplaceAll(S, "_", "")

	return S
}
