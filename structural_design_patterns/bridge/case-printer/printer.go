package main

import "fmt"

type Printer interface {
	PrintFile()
}

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct{}

func (h *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}

type Canon struct{}

func (c *Canon) PrintFile() {
	fmt.Println("Printing by Canon Printer")
}
