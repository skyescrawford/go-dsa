package main

import "fmt"

type Product struct {
	name  string
	price float64
}

func (p *Product) getPrice() float64 {
	return p.price
}

func (p *Product) getName() {
	fmt.Println("product:", p.name)
}
