package main

import (
	"strconv"
	"strings"
)

func Solution_Lovely_Number(A, B int) int {
	result := []int{}
	isLovely := true
	for i := A; i <= B; i++ {
		count := make([]int, len([]rune(strconv.Itoa(i))))
		for k, v := range strconv.Itoa(i) {
			count[k] += strings.Count(strconv.Itoa(i), string(v))
		}
		for _, v := range count {
			if v >= 3 {
				isLovely = false
				break
			}
		}
		if isLovely {
			result = append(result, i)
		}
		isLovely = true
	}
	return len(result)
}
