package main

import (
	"strings"
)

// "regexp"
// "strings"

func Solution_Trees(A []int) int {
	newSlice := make([]int, len(A))
	result := 0

	checkA := checkPerfect(A)
	if checkA {
		return 0
	} else {
		for i, _ := range A {
			copy(newSlice, A)
			newSlice := append(newSlice[:i], newSlice[i+1:]...)
			perfectPa := checkPerfect(newSlice)
			if perfectPa {
				result += 1
			}
		}
	}
	if result == 0 {
		return -1
	}
	return result
}
func checkPerfect(a []int) bool {
	cases := []string{"00", "11", "2", "000", "111", "22"}
	pos := []string{}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			pos = append(pos, "1")
		}
		if a[i] < a[i+1] {
			pos = append(pos, "0")
		}
		if a[i] == a[i+1] {
			pos = append(pos, "2")
		}
	}
	posStr := strings.Join(pos, "")
	for _, v := range cases {
		if strings.Contains(posStr, v) {
			return false
		}
	}
	return true
}

// func handle(A []int) int {
// 	rev := 0
// 	solution := []int{}
// 	for i := 0; i < len(A); i++ {
// 		if i >= len(A)-2 {
// 			if A[i] > A[i] {
// 				if A[i] < A[i+1] {
// 				}
// 			}
// 		}
// 		if i > 0 && i < len(A)-2 {
// 			if rev > A[i] {
// 				if A[i] < A[i+1] {
// 					rev = A[i]
// 					continue
// 				}
// 				if A[i] > A[i+1] {
// 					if i-1 == 0 {
// 						solution = append(solution, i-1)
// 					}
// 					if A[i+2] != 0 {
// 						if A[i+1] > A[i+2] {
// 							return -1
// 						}
// 					}

// 					if A[i] < A[i+2] {
// 						solution = append(solution, i+1)
// 					}
// 					if rev > A[i+1] {
// 						solution = append(solution, i)
// 					}
// 					if A[i-1] < A[i] {
// 						solution = append(solution, i-1)
// 					}
// 				}
// 				return len(solution)
// 			}
// 			if rev < A[i] {
// 				if A[i] > A[i+1] {
// 					rev = A[i]

// 					continue
// 				}
// 				if A[i] < A[i+1] {
// 					if i-1 == 0 {
// 						solution = append(solution, i-1)
// 					}
// 					if A[i+1] < A[i+2] {
// 						return -1
// 					}
// 					if A[i] > A[i+2] {
// 						solution = append(solution, i+1)
// 					}
// 					if rev < A[i+1] {
// 						solution = append(solution, i)
// 					}
// 					if A[i-1] > A[i] {
// 						solution = append(solution, i-1)
// 					}
// 				}
// 				return len(solution)
// 			}
// 			if rev == A[i] {

// 				if A[i] == A[i+1] {
// 					return -1
// 				}
// 				if A[i+1] > A[i+2] {
// 					if A[i] < A[i+1] {
// 						solution = append(solution, i, i-1)
// 					}
// 				}
// 				if A[i+1] < A[i+2] {
// 					if A[i] > A[i+1] {
// 						solution = append(solution, i, i-1)
// 					}
// 				}
// 				return len(solution)
// 			}
// 		}
// 		rev = A[i]
// 	}
// 	return 0
// }

// func Solution_Cut(A []int) int {
// 	pos := []string{}
// 	count := []int{0, 0, 0, 0, 0}
// 	posF := 0
// 	posR := []int{}
// 	cases := []string{"00", "11", "2", "000", "111", "22"}
// 	for i := 0; i < len(A)-1; i++ {
// 		if A[i] > A[i+1] {
// 			pos = append(pos, "1")
// 		}
// 		if A[i] < A[i+1] {
// 			pos = append(pos, "0")
// 		}
// 		if A[i] == A[i+1] {
// 			pos = append(pos, "2")
// 		}
// 	}
// 	fmt.Println(A)
// 	posStr := strings.Join(pos, "")
// 	fmt.Println(posStr)
// 	count[0] = strings.Count(posStr, cases[0])
// 	count[1] = strings.Count(posStr, cases[1])
// 	count[2] = strings.Count(posStr, cases[2])
// 	count[3] = strings.Count(posStr, cases[3])
// 	count[4] = strings.Count(posStr, cases[4])
// 	if count[0] == 0 && count[1] == 0 && count[2] == 0 {
// 		return 0
// 	}
// 	if count[3] > 0 || count[4] > 0 {
// 		return -1
// 	}
// 	if count[0] == 0 && count[1] == 0 && count[2] == 1 {
// 		posF = strings.Index(posStr, cases[2])
// 		solution := []int{}
// 		if posF > 0 {
// 			if A[posF-1] > A[posF] && A[posF+2] > A[posF] {
// 				solution = append(solution, posF, posF+1)
// 			}
// 			if A[posF-1] > A[posF] && A[posF+2] > A[posF] {
// 				solution = append(solution, posF, posF+1)
// 			}
// 		}
// 		if posF == 0 {
// 			if A[posF+2] > A[posF+3] && A[posF+2] > A[posF] {
// 				solution = append(solution, posF, posF+1)
// 			}
// 			if A[posF+2] < A[posF+3] && A[posF+2] < A[posF] {
// 				solution = append(solution, posF, posF+1)
// 			}
// 		}
// 		return len(solution)
// 	}

// 	if count[0] == 1 && count[1] == 0 && count[2] == 0 {
// 		solution := []int{}
// 		fmt.Println(posStr)

// 		posF = strings.Index(posStr, cases[0])
// 		if posF > 0 {
// 			for i, _ := range A {
// 				if i >= posF-1 && i <= posF+3 {
// 					posR = append(posR, i)
// 				}
// 			}
// 			for i := 1; i < len(posR)-1; i++ {
// 				if A[posR[i-1]] == A[posR[i+1]] {
// 					return -1
// 				}
// 				if i%2 != 0 && A[posR[i-1]] > A[posR[i+1]] {
// 					solution = append(solution, posR[i])
// 				} else if i%2 == 0 && A[posR[i-1]] < A[posR[i+1]] {
// 					solution = append(solution, posR[i])
// 				}
// 			}
// 			return len(solution)
// 		}
// 		if posF == 0 {
// 			solution = append(solution, 0)
// 			solution = append(solution, 1)
// 			if A[1] > A[3] {
// 				solution = append(solution, 2)
// 			}
// 			return len(solution)
// 		}

// 	}
// 	if count[0] == 0 && count[1] == 1 && count[2] == 0 {
// 		fmt.Println(posStr)
// 		posF = strings.Index(posStr, cases[1])
// 		solution := []int{}
// 		if posF > 0 {
// 			for i, _ := range A {
// 				if i >= posF-1 && i <= posF+3 {
// 					posR = append(posR, i)
// 				}
// 			}
// 			for i := 1; i < len(posR)-1; i++ {
// 				if A[posR[i-1]] == A[posR[i+1]] {
// 					return -1
// 				}
// 				if i%2 != 0 && A[posR[i-1]] < A[posR[i+1]] {
// 					solution = append(solution, posR[i])
// 				} else if i%2 == 0 && A[posR[i-1]] > A[posR[i+1]] {
// 					solution = append(solution, posR[i])
// 				}
// 			}
// 		}
// 		if posF == 0 {
// 			solution = append(solution, 0)
// 			solution = append(solution, 1)
// 			if A[1] < A[3] {
// 				solution = append(solution, 2)
// 			}
// 		}
// 		return len(solution)
// 	}

// 	return -1
// }
