package main

import "fmt"

// Target interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

type InternalPaymentService struct{}

// its already worked 🚀🚀🚀
func (i *InternalPaymentService) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Internal service processed $%.2f", amount)
}

// third party service with incompatible interface
type StripePaymentService struct{}

// this method incompatible with our internal payment method
func (s *StripePaymentService) MakePayment(dollars float64, cents int) string {
	total := dollars + float64(cents)/100
	return fmt.Sprintf("Stripe charged $%.2f", total)
}

type PaypalService struct{}

func (p *PaypalService) SendMoney(amount float64, currency string) string {
	return fmt.Sprintf("PayPal sent %.2f %s", amount, currency)
}

type StripeAdapter struct {
	stripe *StripePaymentService
}

func (sa *StripeAdapter) ProcessPayment(amount float64) string {
	dollars := int(amount)
	cents := int((amount - float64(dollars)) * 100)
	return sa.stripe.MakePayment(float64(dollars), cents)
}

type PaypalAdapter struct {
	paypal *PaypalService
}

func (pa *PaypalAdapter) ProcessPayment(amount float64) string {
	return pa.paypal.SendMoney(amount, "USD")
}

// Client code that works with the PaymentProcessor interface
func checkout(processor PaymentProcessor, amount float64) {
	result := processor.ProcessPayment(amount)
	fmt.Println(result)
}

func main() {
	internal := &InternalPaymentService{}
	checkout(internal, 6.99)

	stripe := &StripePaymentService{}
	stripeAdapter := &StripeAdapter{stripe: stripe}
	checkout(stripeAdapter, 4.99)

	paypal := &PaypalService{}
	paypalAdapter := &PaypalAdapter{paypal: paypal}
	checkout(paypalAdapter, 5.55)
}
