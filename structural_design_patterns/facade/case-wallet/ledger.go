package main

import "fmt"

type Ledger struct{}

func line(length int) {
	dash := ""
	for i := 0; i < length; i++ {
		dash += "-"
	}
	fmt.Println(dash)
}

func (l *Ledger) makeEntry(id string, transactionType string, amount int) {
	fmt.Printf("make ledger entry for id %s with trancaction type %s for amount %d\n", id, transactionType, amount)
	line(20)
}
