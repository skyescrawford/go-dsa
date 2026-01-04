package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VegieMania struct{}

func (v *VegieMania) getPrice() int {
	return 7
}

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	return t.pizza.getPrice() + 3
}

type CheesePizza struct {
	pizza IPizza
}

func (c *CheesePizza) getPrice() int {
	return c.pizza.getPrice() + 4
}

func main() {
	base := &VegieMania{}
	fmt.Printf("base: %d\n", base.getPrice())

	tomato := &TomatoTopping{pizza: base}
	fmt.Printf("tomato: %d\n", tomato.getPrice())

	tomatoCheese := &CheesePizza{pizza: tomato}
	fmt.Printf("tomato-cheese: %d\n", tomatoCheese.getPrice())

}
