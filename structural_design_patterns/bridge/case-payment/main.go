package main

// GOAL:
// Allow different payment types (one-time, subscription) to work with different payment gateways (Stripe, PayPal) without exploding combinations.

func main() {
	stripe := &StripeGateway{}
	paypal := &PaypalGateway{}
	midtrans := &MidtransGateway{}
	qris := &QrisGateway{}

	oneTime := &OneTimePayment{}
	subscription := &SubscriptionPayment{}

	oneTime.SetPayment(stripe)
	oneTime.Pay(12.45)

	oneTime.SetPayment(paypal)
	oneTime.Pay(22.78)

	subscription.SetPayment(midtrans)
	subscription.Pay(5.55)

	subscription.SetPayment(qris)
	subscription.Pay(11.11)
}
