package main

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name      string
	PaymentID string
	Amount    int
}

func NewPayment(name, cardID string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:      name,
			PaymentID: cardID,
			Amount:    money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

func (p *Payment) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type CreditCard struct{}

func (*CreditCard) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d via %s with credit card number %v\n", ctx.Amount, ctx.Name, ctx.PaymentID)
}

type QRCode struct{}

func (*QRCode) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d via %s with qr code number %v\n", ctx.Amount, ctx.Name, ctx.PaymentID)
}
