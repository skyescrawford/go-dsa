package main

import "fmt"

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

type MidtransGateway struct{}

func (m MidtransGateway) Pay(amount float64) error {
	fmt.Println("Paying", amount, "using Midtrans")
	return nil
}

type QrisGateway struct{}

func (q QrisGateway) Pay(amount float64) error {
	fmt.Println("Paying", amount, "using QRIS")
	return nil
}
