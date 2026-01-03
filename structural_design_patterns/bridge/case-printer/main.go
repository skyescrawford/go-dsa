package main

import "fmt"

func main() {
	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}
	canonPrinter := &Canon{}

	macComputer := &Mac{}
	macComputer.SetPrinter((hpPrinter))
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &Windows{}
	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()

	linuxPC := Linux{}
	linuxPC.SetPrinter(hpPrinter)

	linuxPC.Print()
	fmt.Println()

	linuxPC.SetPrinter(canonPrinter)
	linuxPC.Print()
}
