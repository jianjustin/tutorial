package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) error
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) error {
	fmt.Printf("Paying %.2f with credit card\n", amount)
	return nil
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Paying %.2f with PayPal\n", amount)
	return nil
}

func processPayment(amount float64, paymentStrategy PaymentStrategy) error {
	return paymentStrategy.Pay(amount)
}

func main() {
	creditCard := &CreditCard{}
	payPal := &PayPal{}

	err := processPayment(100, creditCard)
	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}

	err = processPayment(50, payPal)
	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}
}
