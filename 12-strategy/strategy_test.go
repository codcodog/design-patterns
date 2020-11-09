package strategy

func ExamplePayCrash() {
	context := &PaymentContext{
		Name:   "John",
		CardID: "xxx",
		Money:  10,
	}
	p := NewPayment(context, &Cash{})
	p.Pay()
	// Output:
	// cash pay, name: John, card id: xxx, money: 10
}

func ExamplePayBank() {
	context := &PaymentContext{
		Name:   "Bob",
		CardID: "xxx",
		Money:  8,
	}
	p := NewPayment(context, &Bank{})
	p.Pay()
	// Output:
	// bank pay, name: Bob, card id: xxx, money: 8
}
