package main

import "fmt"

func main(){
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer := &mac{}
	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	windowsComputer := &windows{}
	windowsComputer.setPrinter(hpPrinter)
	windowsComputer.print()
	fmt.Println()

	windowsComputer.setPrinter(epsonPrinter)
	windowsComputer.print()
	fmt.Println()

}