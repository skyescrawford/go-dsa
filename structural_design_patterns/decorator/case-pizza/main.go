package main

import "fmt"

type Pizza interface {
	getPrice() float64
	getDescription() string
}

type BasePizza struct{}

func (b *BasePizza) getPrice() float64 {
	return 4.0
}
func (b *BasePizza) getDescription() string {
	return "Pizza"
}

type CheeseTopping struct {
	pizza Pizza
}

func (c *CheeseTopping) getPrice() float64 {
	return c.pizza.getPrice() + 1.2
}
func (c *CheeseTopping) getDescription() string {
	return fmt.Sprintf("%s + Cheese", c.pizza.getDescription())
}

type MeatTopping struct {
	pizza Pizza
}

func (m *MeatTopping) getPrice() float64 {
	return m.pizza.getPrice() + 1.3
}
func (m *MeatTopping) getDescription() string {
	return fmt.Sprintf("%s + Meat", m.pizza.getDescription())
}

type PepperoniTopping struct {
	pizza Pizza
}

func (p *PepperoniTopping) getPrice() float64 {
	return p.pizza.getPrice() + 1.45
}
func (p *PepperoniTopping) getDescription() string {
	return fmt.Sprintf("%s + Pepperoni", p.pizza.getDescription())
}

type VeggieTopping struct {
	pizza Pizza
}

func (v *VeggieTopping) getPrice() float64 {
	return v.pizza.getPrice()
}
func (v *VeggieTopping) getDescription() string {
	return fmt.Sprintf("%s + Veggie", v.pizza.getDescription())
}

func main() {
	base := &BasePizza{}
	cheese := &CheeseTopping{pizza: base}
	meat := &MeatTopping{pizza: cheese}
	pepperoni := &PepperoniTopping{pizza: meat}
	veggie := &VeggieTopping{pizza: pepperoni}

	fmt.Printf("Order: %s\n", veggie.getDescription())
	fmt.Printf("Price: %.2f\n", veggie.getPrice())
}
