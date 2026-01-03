package main

import "fmt"

type Package interface {
	getPrice() float64
	getName()
}

func main() {
	product1 := &Product{name: "baju", price: 2.1}
	product2 := &Product{name: "celana", price: 3.2}
	product3 := &Product{name: "topi", price: 1.3}

	box1 := &Box{name: "box 1"}
	box1.add(product1)
	box1.add(product2)

	box2 := &Box{name: "box 2"}
	box2.add(box1)
	box2.add(product3)

	fmt.Printf("total price: %.2f\n", box2.getPrice())

	box2.getName()
}
