package main

import "fmt"

type Mk47Gun struct {
	name string
}

func design(name string) {
	fmt.Println("design for gun: ", name)
}

func buildGun(name string) string {
	return name
}
