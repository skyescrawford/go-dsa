package main

import "fmt"

// GOAL:
// Allow different payment types (one-time, subscription) to work with different payment gateways (Stripe, PayPal) without exploding combinations.

// implementator interface (Bridge)
type PaymentGateway interface {
	Pay(amount float64) error
}

// concrete implementation
type StripeGateway struct{}

func (s StripeGateway) Pay(amount float64) error {
	fmt.Println("Paying", amount, "using Stripe")
	return nil
}

type PaypalGateway struct{}

func (p PaypalGateway) Pay(amount float64) error {
	fmt.Println("Paying", amount, "using Paypal")
	return nil
}

// abstraction (Bridge Holder)
type Payment struct {
	gateway PaymentGateway
}

func NewPayment(p PaymentGateway) *Payment {
	return &Payment{gateway: p}
}

func (p *Payment) Pay(amount float64) error {
	return p.gateway.Pay(amount)
}

type OneTimePayment struct {
	*Payment
}

func (o *OneTimePayment) Pay(amount float64) error {
	fmt.Println("One-time payment")
	return o.gateway.Pay(amount)
}

type SubscriptionPayment struct {
	*Payment
}

func (s *SubscriptionPayment) Pay(amount float64) error {
	fmt.Println("Subscription payment")
	return s.gateway.Pay(amount)
}

func main() {
	stripe := &StripeGateway{}
	paypal := &PaypalGateway{}

	oneTime := &OneTimePayment{
		Payment: NewPayment(stripe),
	}

	subscription := &SubscriptionPayment{
		Payment: NewPayment(paypal),
	}

	oneTime.Pay(99.99)
	subscription.Pay(45.67)
}
