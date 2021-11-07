package main

import "fmt"

type Summember struct {
	numberCount int
	sum         int
}
type SumWar struct {
	sumChild []Summember
}

func Solution_Magic_Square(square [][]int) int {
	y := len(square)
	x := len(square[0])
	// squareChildGroup := [][][]int{}
	isSquare := false
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			rowChild := []int{}
			columnChild := []int{}
			squareChild := [][]int{}
			for g := j; g < x; g++ {
				rowChild = append(rowChild, square[j][g])
				fmt.Println("rowChild: ", rowChild)
				columnChild = append(columnChild, square[j][g])
				fmt.Println("columnChild: ", columnChild)

				// if k == g {
				// 	isSquare = true
				// 	break
				// }
			}
			for k := i; k < y; k++ {

				// squareChildGroup = append(squareChildGroup, squareChild)
			}
			if isSquare {
				squareChild = append(squareChild, rowChild)
				squareChild = append(squareChild, columnChild)
				fmt.Println(squareChild)

			}
			fmt.Println("End j: ", j)
		}
		fmt.Println("End i: ", i)
	}

	return len(square)
}

// func Solution_Magic_Square(square [][]int) int {
// 	sumColumnArr := make(map[int][]Summember)
// 	sumRowArr := make(map[int][]Summember)
// 	for i := 0; i < len(square); i++ {
// 		sum := 0
// 		sumColumn := []Summember{}
// 		for j := 0; j < len(square[i]); j++ {
// 			if i == 0 {
// 				sumRow := []Summember{}
// 				column := 0
// 				for k, v := range square {
// 					column += v[j]
// 					sumRow = append(sumRow, Summember{
// 						numberCount: k + 1,
// 						sum:         column,
// 					})
// 				}
// 				sumRowArr[j] = sumRow
// 			}
// 			sum += square[i][j]
// 			sumColumn = append(sumColumn, Summember{
// 				numberCount: j + 1,
// 				sum:         sum,
// 			})
// 			sumColumnArr[i] = sumColumn
// 		}
// 	}
// 	fmt.Println(sumColumnArr)
// 	fmt.Println(len(sumColumnArr))

// 	fmt.Println(sumRowArr)
// 	fmt.Println(len(sumRowArr))

// 	return len(square)
// }
