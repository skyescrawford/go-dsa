package main

import "fmt"

type Printer interface {
	Print(msg string)
}

// Adaptee (existing, already worked)
type LegacyPrinter struct{}

func (l *LegacyPrinter) PrintText(msg string) {
	fmt.Printf("[Legacy Printer] %s\n", msg)
}

// Adapter
type PrinterAdapter struct {
	legacy *LegacyPrinter
}

func (p *PrinterAdapter) Print(msg string) {
	p.legacy.PrintText(msg)
}

// ga bisa passing LegacyPrinter langsung ke StdOut karena adanya method PrintText, bukan Print
// jadi perlu adapter
func StdOut(p Printer, msg string) {
	p.Print(msg)
}

// Legacy printer ga punya method Print
// dengan menggunakan adapter, kita bisa memanggil method Print dengan menjadikan legacy printer sebagai attribut/properti PrinterAdapter
func main() {
	legacy := &LegacyPrinter{}

	adapter := &PrinterAdapter{legacy: legacy}

	StdOut(adapter, "hello world")
}
