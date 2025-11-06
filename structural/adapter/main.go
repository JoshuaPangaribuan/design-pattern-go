package main

import (
	"fmt"
	"time"
)

// JoshBankPaymentProcessor is the target interface that our application expects.
// All payment methods in JoshBank should implement this interface.
type JoshBankPaymentProcessor interface {
	ProcessPayment(amount float64, currency string) (*PaymentResult, error)
	RefundPayment(transactionID string, amount float64) error
}

// PaymentResult represents the result of a payment transaction
type PaymentResult struct {
	TransactionID string
	Status        string
	ProcessedAt   time.Time
}

// --- JoshBank Internal Payment System (already implemented) ---

// JoshBankInternalPaymentSystem is our existing payment processor
type JoshBankInternalPaymentSystem struct {
	merchantID string
}

func (p *JoshBankInternalPaymentSystem) ProcessPayment(amount float64, currency string) (*PaymentResult, error) {
	fmt.Printf("[JoshBank Internal] Processing payment: %.2f %s\n", amount, currency)
	return &PaymentResult{
		TransactionID: "JOSH-" + fmt.Sprintf("%d", time.Now().Unix()),
		Status:        "completed",
		ProcessedAt:   time.Now(),
	}, nil
}

func (p *JoshBankInternalPaymentSystem) RefundPayment(transactionID string, amount float64) error {
	fmt.Printf("[JoshBank Internal] Refunding transaction %s: %.2f\n", transactionID, amount)
	return nil
}

// --- Third-Party: Legacy Bank System (Adaptee - incompatible interface) ---

// LegacyBankAPI represents a legacy bank system with different interface.
// This is the "adaptee" - it has methods we need but different signatures.
type LegacyBankAPI struct {
	apiKey string
}

// CreateTransaction is Legacy Bank's method for processing payments (different from our interface)
func (l *LegacyBankAPI) CreateTransaction(amountInCents int, curr string, description string) (string, error) {
	fmt.Printf("[Legacy Bank] Creating transaction: %d cents %s - %s\n", amountInCents, curr, description)
	transactionID := fmt.Sprintf("LEGACY_%d", time.Now().Unix())
	return transactionID, nil
}

// ProcessRefund is Legacy Bank's method for refunds (different from our interface)
func (l *LegacyBankAPI) ProcessRefund(transactionID string, amountInCents int) error {
	fmt.Printf("[Legacy Bank] Processing refund for transaction %s: %d cents\n", transactionID, amountInCents)
	return nil
}

// --- Adapter: Makes Legacy Bank compatible with our interface ---

// LegacyBankAdapter adapts the Legacy Bank API to our JoshBankPaymentProcessor interface.
// This is the adapter that bridges the incompatible interfaces.
type LegacyBankAdapter struct {
	legacyBank *LegacyBankAPI
}

func NewLegacyBankAdapter(apiKey string) *LegacyBankAdapter {
	return &LegacyBankAdapter{
		legacyBank: &LegacyBankAPI{apiKey: apiKey},
	}
}

// ProcessPayment adapts our interface to Legacy Bank's CreateTransaction method
func (a *LegacyBankAdapter) ProcessPayment(amount float64, currency string) (*PaymentResult, error) {
	// Convert dollars to cents (Legacy Bank uses cents)
	amountInCents := int(amount * 100)

	// Call Legacy Bank's method with adapted parameters
	transactionID, err := a.legacyBank.CreateTransaction(amountInCents, currency, "Payment via JoshBank adapter")
	if err != nil {
		return nil, err
	}

	// Convert Legacy Bank's response to our format
	return &PaymentResult{
		TransactionID: transactionID,
		Status:        "completed",
		ProcessedAt:   time.Now(),
	}, nil
}

// RefundPayment adapts our interface to Legacy Bank's ProcessRefund method
func (a *LegacyBankAdapter) RefundPayment(transactionID string, amount float64) error {
	// Convert dollars to cents
	amountInCents := int(amount * 100)

	// Call Legacy Bank's refund method
	return a.legacyBank.ProcessRefund(transactionID, amountInCents)
}

// --- Third-Party: External Payment Gateway (Another Adaptee) ---

// ExternalPaymentGateway represents another third-party provider with yet another interface
type ExternalPaymentGateway struct {
	clientID string
}

func (e *ExternalPaymentGateway) ExecutePayment(amountStr string, currencyCode string) (map[string]interface{}, error) {
	fmt.Printf("[External Gateway] Executing payment: %s %s\n", amountStr, currencyCode)
	return map[string]interface{}{
		"payment_id": fmt.Sprintf("EXT-%d", time.Now().Unix()),
		"state":      "approved",
	}, nil
}

func (e *ExternalPaymentGateway) RefundTransaction(paymentID string, refundAmount string) error {
	fmt.Printf("[External Gateway] Refunding payment %s: %s\n", paymentID, refundAmount)
	return nil
}

// ExternalGatewayAdapter adapts External Gateway SDK to our interface
type ExternalGatewayAdapter struct {
	gateway *ExternalPaymentGateway
}

func NewExternalGatewayAdapter(clientID string) *ExternalGatewayAdapter {
	return &ExternalGatewayAdapter{
		gateway: &ExternalPaymentGateway{clientID: clientID},
	}
}

func (a *ExternalGatewayAdapter) ProcessPayment(amount float64, currency string) (*PaymentResult, error) {
	// Convert amount to string (External Gateway expects string)
	amountStr := fmt.Sprintf("%.2f", amount)

	// Call External Gateway's method
	response, err := a.gateway.ExecutePayment(amountStr, currency)
	if err != nil {
		return nil, err
	}

	// Convert External Gateway's response to our format
	return &PaymentResult{
		TransactionID: response["payment_id"].(string),
		Status:        response["state"].(string),
		ProcessedAt:   time.Now(),
	}, nil
}

func (a *ExternalGatewayAdapter) RefundPayment(transactionID string, amount float64) error {
	amountStr := fmt.Sprintf("%.2f", amount)
	return a.gateway.RefundTransaction(transactionID, amountStr)
}

// --- Client Code ---

// processOrder demonstrates client code that works with any JoshBankPaymentProcessor.
// It doesn't need to know about Legacy Bank, External Gateway, or their specific APIs.
func processOrder(processor JoshBankPaymentProcessor, amount float64, currency string) {
	fmt.Println("\n--- Processing Order ---")

	// Process payment using the common interface
	result, err := processor.ProcessPayment(amount, currency)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}

	fmt.Printf("✓ Payment successful!\n")
	fmt.Printf("  Transaction ID: %s\n", result.TransactionID)
	fmt.Printf("  Status: %s\n", result.Status)

	// Simulate a refund scenario
	fmt.Println("\n--- Processing Refund ---")
	err = processor.RefundPayment(result.TransactionID, amount/2)
	if err != nil {
		fmt.Printf("Refund failed: %v\n", err)
		return
	}
	fmt.Printf("✓ Refund successful!\n")
}

func main() {
	fmt.Println("=== Adapter Pattern: JoshBank Payment Gateway Integration ===")

	// Example 1: Using JoshBank internal payment system (no adapter needed)
	fmt.Println("\n=== Example 1: JoshBank Internal Payment System ===")
	internalProcessor := &JoshBankInternalPaymentSystem{merchantID: "JOSH123"}
	processOrder(internalProcessor, 99.99, "USD")

	// Example 2: Using Legacy Bank through adapter
	fmt.Println("\n=== Example 2: Legacy Bank System (via Adapter) ===")
	legacyProcessor := NewLegacyBankAdapter("legacy_api_key")
	processOrder(legacyProcessor, 149.99, "USD")

	// Example 3: Using External Gateway through adapter
	fmt.Println("\n=== Example 3: External Payment Gateway (via Adapter) ===")
	externalProcessor := NewExternalGatewayAdapter("external_client_id")
	processOrder(externalProcessor, 199.99, "USD")

	// Example 4: Switching payment providers at runtime
	fmt.Println("\n=== Example 4: Runtime Provider Selection ===")

	providers := map[string]JoshBankPaymentProcessor{
		"internal": internalProcessor,
		"legacy":   legacyProcessor,
		"external": externalProcessor,
	}

	// Customer selects payment method
	selectedProvider := "legacy"
	fmt.Printf("Customer selected: %s\n", selectedProvider)

	processor := providers[selectedProvider]
	processOrder(processor, 299.99, "USD")

	fmt.Println("\n✓ Adapter pattern enables seamless integration of different payment providers")
	fmt.Println("✓ Client code remains unchanged when adding new providers")
	fmt.Println("✓ Each provider's unique API is hidden behind common interface")
	fmt.Println("✓ JoshBank can integrate legacy systems without modifying core code")
}
