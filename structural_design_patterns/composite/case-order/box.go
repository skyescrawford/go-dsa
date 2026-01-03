package main

import "fmt"

type Box struct {
	name     string
	packages []Package
}

func (b *Box) getPrice() float64 {
	var price float64
	for _, pkg := range b.packages {
		price += pkg.getPrice()
	}
	return price
}

func (b *Box) add(p Package) {
	b.packages = append(b.packages, p)
}

func (b *Box) getName() {
	fmt.Println("box:", b.name)
	for _, pkg := range b.packages {
		pkg.getName()
	}
}
