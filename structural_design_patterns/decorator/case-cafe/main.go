package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func drawLine(length int) {
	line := ""
	for i := 0; i < length; i++ {
		line += "-"
	}
	fmt.Printf("%s\n", line)
}

func createOrder(coffee *Coffee, reader *bufio.Reader) {
	for {
		fmt.Println("Order")
		fmt.Println("1. Coffee")
		fmt.Println("2. Milk")
		fmt.Println("3. Cheese")
		fmt.Println("9. Set Order")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")
		line, _, err := reader.ReadLine()

		if string(line) == "done" {
			break
		}

		if err != nil {
			os.Exit(1)
		}

		selected, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Println("invalid order")
			os.Exit(1)
		}

		switch selected {
		case 1:
			*coffee = &SimpleCoffee{}
		case 2:
			*coffee = &MilkDecorator{coffee: *coffee}
		case 3:
			*coffee = &CheeseDecorator{coffee: *coffee}
		case 9:
			drawLine(30)
			fmt.Println("Calculating order...🚀")
			time.Sleep(3 * time.Second)
			return
		case 0:
			fmt.Println("Cancelled ❌")
			os.Exit(1)
		}

		clearScreen()
	}

}

func main() {
	clearScreen()

	var coffee Coffee

	reader := bufio.NewReader(os.Stdin)

	createOrder(&coffee, reader)

	fmt.Printf("Order info:\t%s = $%d\n", coffee.GetDescription(), coffee.GetCost())
}
