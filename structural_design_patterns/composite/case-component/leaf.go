package main

import "fmt"

type Leaf struct {
	name string
}

func (l *Leaf) Operation() {
	fmt.Printf("%s\n", l.name)
}
