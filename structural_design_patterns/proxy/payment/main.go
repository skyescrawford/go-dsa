package main

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(int) error
}

type Money struct {
	balance int
}

func (m *Money) Pay(amount int) error {
	if m.balance < amount {
		return errors.New("insufficient balance")
	}

	m.balance -= amount
	fmt.Printf("paying for %d\n", amount)
	return nil
}

type CreditCard struct {
	money *Money
}

func (c *CreditCard) Pay(amount int) error {
	if c.money == nil {
		c.money = &Money{balance: 50000}
	}

	fmt.Println("paying with credit card...")

	err := c.money.Pay(amount)
	if err != nil {
		return err
	}

	fmt.Println("paying with credit card successfully")
	return nil
}

func main() {
	creditCard := &CreditCard{}
	err := creditCard.Pay(50001)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}
