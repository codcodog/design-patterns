package strategy

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

func NewPayment(context *PaymentContext, strategy PaymentStrategy) *Payment {
	return &Payment{
		context:  context,
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.pay(p.context)
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

type PaymentStrategy interface {
	pay(*PaymentContext)
}

type Cash struct {
}

func (c *Cash) pay(context *PaymentContext) {
	fmt.Printf("cash pay, name: %s, card id: %s, money: %d", context.Name, context.CardID, context.Money)
}

type Bank struct {
}

func (b *Bank) pay(context *PaymentContext) {
	fmt.Printf("bank pay, name: %s, card id: %s, money: %d", context.Name, context.CardID, context.Money)
}
