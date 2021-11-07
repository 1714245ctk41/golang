package main

import "fmt"

func main() {
	fmt.Println(("Hello evyerone"))
	x := 23 + 23
	y := "Hello world"
	z := 0
	fmt.Println(x, y, z)
	foo()
}

func foo() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
		bar()
	}
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println(("and then we exited"))
}
