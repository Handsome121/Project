/*
里氏替换原则（LSP） 里氏替换原则指派生类的对象应该能够替换基类的对象而不影响程序的正确性。
*/

package solid

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type CreditCardProcessor struct{}

func (c *CreditCardProcessor) ProcessPayment(amount float64) error {
	// Process credit card payment
	return nil
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) ProcessPayment(amount float64) error {
	// Process PayPal payment
	return nil
}

func main() {
	creditCardProcessor := &CreditCardProcessor{}
	payPalProcessor := &PayPalProcessor{}

	processPayment(creditCardProcessor, 100)
	processPayment(payPalProcessor, 200)
}

func processPayment(processor PaymentProcessor, amount float64) {
	err := processor.ProcessPayment(amount)
	if err != nil {
		return
	}
}
