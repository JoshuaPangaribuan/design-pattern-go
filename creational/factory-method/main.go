package main

import (
	"fmt"
)

// PaymentMethod is the interface that all payment methods must implement.
// This allows JoshBank to work with different payment types uniformly.
type PaymentMethod interface {
	Validate() error
	Process(amount float64) error
	GetDetails() string
}

// PaymentType represents the type of payment method
type PaymentType string

const (
	CreditCardType   PaymentType = "creditcard"
	BankTransferType PaymentType = "banktransfer"
	CryptoType       PaymentType = "crypto"
)

// CreditCardPayment handles credit card transactions
type CreditCardPayment struct {
	cardNumber string
	cvv        string
	expiry     string
}

func (c *CreditCardPayment) Validate() error {
	if len(c.cardNumber) != 16 {
		return fmt.Errorf("invalid card number length")
	}
	if len(c.cvv) != 3 {
		return fmt.Errorf("invalid CVV")
	}
	return nil
}

func (c *CreditCardPayment) Process(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	fmt.Printf("Card ending in %s\n", c.cardNumber[12:])
	return nil
}

func (c *CreditCardPayment) GetDetails() string {
	return fmt.Sprintf("Credit Card ending in %s", c.cardNumber[12:])
}

// BankTransferPayment handles bank transfer transactions
type BankTransferPayment struct {
	accountNumber string
	routingNumber string
}

func (b *BankTransferPayment) Validate() error {
	if len(b.accountNumber) < 8 {
		return fmt.Errorf("invalid account number")
	}
	if len(b.routingNumber) != 9 {
		return fmt.Errorf("invalid routing number")
	}
	return nil
}

func (b *BankTransferPayment) Process(amount float64) error {
	fmt.Printf("Processing bank transfer of $%.2f\n", amount)
	fmt.Printf("Account: %s, Routing: %s\n", b.accountNumber, b.routingNumber)
	return nil
}

func (b *BankTransferPayment) GetDetails() string {
	return fmt.Sprintf("Bank Transfer - Account: %s", b.accountNumber)
}

// CryptoPayment handles cryptocurrency transactions
type CryptoPayment struct {
	walletAddress string
	currency      string
}

func (c *CryptoPayment) Validate() error {
	if len(c.walletAddress) < 26 {
		return fmt.Errorf("invalid wallet address")
	}
	return nil
}

func (c *CryptoPayment) Process(amount float64) error {
	fmt.Printf("Processing %s payment of $%.2f\n", c.currency, amount)
	fmt.Printf("Wallet: %s...%s\n", c.walletAddress[:6], c.walletAddress[len(c.walletAddress)-4:])
	return nil
}

func (c *CryptoPayment) GetDetails() string {
	return fmt.Sprintf("%s wallet: %s...%s", c.currency,
		c.walletAddress[:6], c.walletAddress[len(c.walletAddress)-4:])
}

// NewPaymentMethod is the factory function that creates the appropriate payment method.
// This is the Factory Method - it encapsulates the creation logic.
func NewPaymentMethod(paymentType PaymentType, details map[string]string) (PaymentMethod, error) {
	switch paymentType {
	case CreditCardType:
		return &CreditCardPayment{
			cardNumber: details["card_number"],
			cvv:        details["cvv"],
			expiry:     details["expiry"],
		}, nil

	case BankTransferType:
		return &BankTransferPayment{
			accountNumber: details["account_number"],
			routingNumber: details["routing_number"],
		}, nil

	case CryptoType:
		return &CryptoPayment{
			walletAddress: details["wallet_address"],
			currency:      details["currency"],
		}, nil

	default:
		return nil, fmt.Errorf("unsupported payment type: %s", paymentType)
	}
}

// processPayment demonstrates how client code uses the factory.
// It doesn't need to know about concrete payment types.
func processPayment(paymentType PaymentType, details map[string]string, amount float64) {
	fmt.Printf("\n--- Processing %s Payment ---\n", paymentType)

	// Use factory to create the payment method
	payment, err := NewPaymentMethod(paymentType, details)
	if err != nil {
		fmt.Printf("Error creating payment: %v\n", err)
		return
	}

	// Validate the payment details
	if err := payment.Validate(); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		return
	}

	// Process the payment
	if err := payment.Process(amount); err != nil {
		fmt.Printf("Processing failed: %v\n", err)
		return
	}

	fmt.Printf("✓ Payment successful via %s\n", payment.GetDetails())
}

func main() {
	fmt.Println("=== Factory Method Pattern: JoshBank Payment Processing ===\n")

	// Example 1: Credit Card Payment
	processPayment(CreditCardType, map[string]string{
		"card_number": "1234567890123456",
		"cvv":         "123",
		"expiry":      "12/25",
	}, 99.99)

	// Example 2: Bank Transfer Payment
	processPayment(BankTransferType, map[string]string{
		"account_number": "1234567890",
		"routing_number": "987654321",
	}, 149.50)

	// Example 3: Cryptocurrency Payment
	processPayment(CryptoType, map[string]string{
		"wallet_address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		"currency":       "Bitcoin",
	}, 299.99)

	// Example 4: Invalid payment type
	fmt.Println("\n--- Attempting Invalid Payment Type ---")
	_, err := NewPaymentMethod("invalid", map[string]string{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n✓ Factory Method allows easy addition of new payment types")
	fmt.Println("✓ Client code doesn't depend on concrete payment classes")
	fmt.Println("✓ JoshBank can easily integrate new payment providers")
}
