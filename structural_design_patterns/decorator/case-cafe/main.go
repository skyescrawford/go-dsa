package main

import "fmt"

type Coffee interface {
	GetCost() int
	GetDescription() string
}

type SimpleCoffee struct{}

func (s *SimpleCoffee) GetCost() int {
	return 5
}

func (s *SimpleCoffee) GetDescription() string {
	return "Coffee"
}

type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) GetCost() int {
	return 2 + m.coffee.GetCost()
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", milk"
}

type CheeseDecorator struct {
	coffee Coffee
}

func (c *CheeseDecorator) GetCost() int {
	return 3 + c.coffee.GetCost()
}

func (c *CheeseDecorator) GetDescription() string {
	return c.coffee.GetDescription() + ", cheese"
}

func main() {
	coffee := &SimpleCoffee{}
	coffeeWithMilk := &MilkDecorator{coffee: coffee}
	coffeeWithMilkCheese := &CheeseDecorator{coffee: coffeeWithMilk}

	fmt.Printf("%s = $%d\n", coffeeWithMilkCheese.GetDescription(), coffeeWithMilkCheese.GetCost())
}
