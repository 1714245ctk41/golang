package main

import (
	"strconv"
	"strings"
)

func Solution_Digital_clock(digits []int) int {
	hours := []string{}
	minutes := []string{}
	houMin := []string{}
	result := []string{}
	for i := 0; i < len(digits)-1; i++ {
		for j := i + 1; j < len(digits); j++ {
			digitsStr1 := strconv.Itoa(digits[i])
			digitsStr2 := strconv.Itoa(digits[j])
			num := digitsStr1 + digitsStr2
			hours, minutes = checkNumber(hours, minutes, num)
			num = digitsStr2 + digitsStr1
			hours, minutes = checkNumber(hours, minutes, num)
		}
	}
	hours = removeDuplicate(hours)
	minutes = removeDuplicate(minutes)
	for i := 0; i < len(minutes); i++ {
		for j := 0; j < len(hours); j++ {
			num := hours[j] + minutes[i]
			houMin = append(houMin, num)
		}
	}

	for _, v := range houMin {
		if isPerfectNum(v, digits) {
			result = append(result, v)
		}
	}
	return len(result)
}
func isPerfectNum(v string, digits []int) bool {
	vSli := strings.Split(v, "")
	digitsCopy := make([]int, len(digits))
	copy(digitsCopy, digits)
	for _, v := range vSli {
		for j := 0; j < len(digitsCopy); j++ {
			if numInt, _ := strconv.Atoi(v); numInt == digitsCopy[j] {
				digitsCopy[j] = 99
				break
			}
		}
	}
	count := 0
	for _, v := range digitsCopy {
		if v == 99 {
			count += 1
			continue
		}
	}
	if count == 4 {
		return true
	}
	return false
}

func removeDuplicate(houMin []string) []string {
	for i := 0; i < len(houMin)-1; i++ {
		for j := i + 1; j < len(houMin); j++ {
			if houMin[i] == houMin[j] {
				houMin[j] = "_"
			}
		}
	}
	newString := []string{}
	for _, v := range houMin {
		if v != "_" {
			newString = append(newString, v)
		}
	}
	return newString
}

func checkNumber(hours, minutes []string, num string) ([]string, []string) {
	if numInt, _ := strconv.Atoi(num); numInt < 24 {
		hours = append(hours, num)
	}
	if numInt, _ := strconv.Atoi(num); numInt < 60 {
		minutes = append(minutes, num)
	}
	return hours, minutes
}

// newSlice := make([]string, len(digits))

// for i, vc := range digits {
// 	if strings.Contains(v, strconv.Itoa(vc)) {
// 		newSlice[i] += "1"
// 	}
// }
// newSliceStr := strings.Join(newSlice, "")
// fmt.Println(newSliceStr)
// if strings.Count(newSliceStr, "1") == 4 {
// 	for _, vc1 := range result {
// 		if vc1 == v {
// 			isDuplicate = true
// 			break
// 		}
// 	}

// 	if !isDuplicate {
// 		result = append(result, v)
// 		isDuplicate = false
// 	}
// 	continue
// }
