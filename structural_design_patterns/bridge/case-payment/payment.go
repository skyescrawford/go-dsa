package main

import "fmt"

// abstraction (Bridge Holder)
type Payment interface {
	SetPayment(PaymentGateway)
	Pay(float64) error
}

type OneTimePayment struct {
	gateway PaymentGateway
}

func (o *OneTimePayment) SetPayment(gateway PaymentGateway) {
	o.gateway = gateway
}

func (o *OneTimePayment) Pay(amount float64) error {
	fmt.Println("One-time payment")
	return o.gateway.Pay(amount)
}

type SubscriptionPayment struct {
	gateway PaymentGateway
}

func (s *SubscriptionPayment) SetPayment(gateway PaymentGateway) {
	s.gateway = gateway
}

func (s *SubscriptionPayment) Pay(amount float64) error {
	fmt.Println("Subscription payment")
	return s.gateway.Pay(amount)
}
