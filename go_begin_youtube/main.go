package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Threads: %v\n", runtime.GOMAXPROCS(-1))
}
