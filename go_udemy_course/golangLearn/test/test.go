package main

import (
	"fmt"
)

func	main(){
	fmt.Println(" sum = ", mySum(2,5,3,3))
}

func mySum(xi ...int) int{
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}