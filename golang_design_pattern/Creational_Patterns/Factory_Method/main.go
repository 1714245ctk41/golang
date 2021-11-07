package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun){
	fmt.Println("Gun Name: ", g.getName())
	fmt.Println()
	fmt.Println("Power Gun: ", g.getPower())
	fmt.Println()
}