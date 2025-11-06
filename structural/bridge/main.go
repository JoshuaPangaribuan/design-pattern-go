package main

import "fmt"

// PaymentProcessor is the implementation interface.
// This represents the "implementation" side of the bridge.
// Different payment processing methods implement this interface.
type PaymentProcessor interface {
	ProcessPayment(amount float64, merchantID string) error
	GetProcessorName() string
}

// --- Concrete Implementations (Payment Processors) ---

// CreditCardProcessor implements payment processing for credit cards
type CreditCardProcessor struct{}

func (p *CreditCardProcessor) ProcessPayment(amount float64, merchantID string) error {
	fmt.Printf("[Credit Card] Processing $%.2f payment via credit card network\n", amount)
	return nil
}

func (p *CreditCardProcessor) GetProcessorName() string {
	return "Credit Card"
}

// BankTransferProcessor implements payment processing for bank transfers
type BankTransferProcessor struct{}

func (p *BankTransferProcessor) ProcessPayment(amount float64, merchantID string) error {
	fmt.Printf("[Bank Transfer] Processing $%.2f payment via ACH network\n", amount)
	return nil
}

func (p *BankTransferProcessor) GetProcessorName() string {
	return "Bank Transfer"
}

// CryptoProcessor implements payment processing for cryptocurrency
type CryptoProcessor struct{}

func (p *CryptoProcessor) ProcessPayment(amount float64, merchantID string) error {
	fmt.Printf("[Crypto] Processing $%.2f payment via blockchain network\n", amount)
	return nil
}

func (p *CryptoProcessor) GetProcessorName() string {
	return "Cryptocurrency"
}

// --- Abstraction ---

// Account is the abstraction that uses the implementation.
// This is the "abstraction" side of the bridge.
type Account struct {
	processor PaymentProcessor
}

func NewAccount(processor PaymentProcessor) *Account {
	return &Account{processor: processor}
}

// ProcessPayment is a basic method that delegates to the implementation
func (a *Account) ProcessPayment(amount float64, merchantID string) error {
	fmt.Printf("Processing payment using %s processor...\n", a.processor.GetProcessorName())
	return a.processor.ProcessPayment(amount, merchantID)
}

// --- Refined Abstractions (Account Types) ---

// CheckingAccount is a refined abstraction for checking accounts
type CheckingAccount struct {
	*Account
	accountNumber string
	balance       float64
}

func NewCheckingAccount(processor PaymentProcessor, accountNumber string, balance float64) *CheckingAccount {
	return &CheckingAccount{
		Account:       NewAccount(processor),
		accountNumber: accountNumber,
		balance:       balance,
	}
}

func (c *CheckingAccount) ProcessPayment(amount float64, merchantID string) error {
	if c.balance < amount {
		return fmt.Errorf("insufficient funds in checking account")
	}
	fmt.Printf("[Checking Account %s] Balance: $%.2f\n", c.accountNumber, c.balance)
	c.balance -= amount
	return c.Account.ProcessPayment(amount, merchantID)
}

// SavingsAccount is a refined abstraction for savings accounts
type SavingsAccount struct {
	*Account
	accountNumber string
	balance       float64
	minBalance    float64
}

func NewSavingsAccount(processor PaymentProcessor, accountNumber string, balance float64, minBalance float64) *SavingsAccount {
	return &SavingsAccount{
		Account:       NewAccount(processor),
		accountNumber: accountNumber,
		balance:       balance,
		minBalance:    minBalance,
	}
}

func (s *SavingsAccount) ProcessPayment(amount float64, merchantID string) error {
	if s.balance-amount < s.minBalance {
		return fmt.Errorf("payment would violate minimum balance requirement")
	}
	fmt.Printf("[Savings Account %s] Balance: $%.2f, Min Balance: $%.2f\n", s.accountNumber, s.balance, s.minBalance)
	s.balance -= amount
	return s.Account.ProcessPayment(amount, merchantID)
}

// InvestmentAccount is a refined abstraction for investment accounts
type InvestmentAccount struct {
	*Account
	accountNumber string
	balance       float64
}

func NewInvestmentAccount(processor PaymentProcessor, accountNumber string, balance float64) *InvestmentAccount {
	return &InvestmentAccount{
		Account:       NewAccount(processor),
		accountNumber: accountNumber,
		balance:       balance,
	}
}

func (i *InvestmentAccount) ProcessPayment(amount float64, merchantID string) error {
	fmt.Printf("[Investment Account %s] Processing investment transaction\n", i.accountNumber)
	return i.Account.ProcessPayment(amount, merchantID)
}

func main() {
	fmt.Println("=== Bridge Pattern: JoshBank Account & Payment Processing ===")

	// Create different payment processors (implementations)
	creditCardProcessor := &CreditCardProcessor{}
	bankTransferProcessor := &BankTransferProcessor{}
	cryptoProcessor := &CryptoProcessor{}

	// Example 1: Checking accounts with different processors
	fmt.Println("\n--- Example 1: Checking Accounts ---")

	checking1 := NewCheckingAccount(creditCardProcessor, "CHK001", 5000.0)
	checking1.ProcessPayment(100.0, "MERCH001")

	fmt.Println()
	checking2 := NewCheckingAccount(bankTransferProcessor, "CHK002", 3000.0)
	checking2.ProcessPayment(250.0, "MERCH002")

	// Example 2: Savings accounts with different processors
	fmt.Println("\n--- Example 2: Savings Accounts ---")

	savings1 := NewSavingsAccount(creditCardProcessor, "SAV001", 10000.0, 1000.0)
	savings1.ProcessPayment(500.0, "MERCH003")

	fmt.Println()
	savings2 := NewSavingsAccount(cryptoProcessor, "SAV002", 15000.0, 2000.0)
	savings2.ProcessPayment(1000.0, "MERCH004")

	// Example 3: Investment accounts
	fmt.Println("\n--- Example 3: Investment Accounts ---")

	investment1 := NewInvestmentAccount(bankTransferProcessor, "INV001", 50000.0)
	investment1.ProcessPayment(5000.0, "MERCH005")

	fmt.Println()
	investment2 := NewInvestmentAccount(cryptoProcessor, "INV002", 75000.0)
	investment2.ProcessPayment(10000.0, "MERCH006")

	// Example 4: Switching processors at runtime
	fmt.Println("\n--- Example 4: Runtime Processor Switching ---")

	account := NewAccount(creditCardProcessor)
	account.ProcessPayment(50.0, "MERCH007")

	// Switch to bank transfer
	account.processor = bankTransferProcessor
	account.ProcessPayment(50.0, "MERCH007")

	// Switch to crypto
	account.processor = cryptoProcessor
	account.ProcessPayment(50.0, "MERCH007")

	fmt.Println("\n✓ Bridge pattern separates account type from payment processing")
	fmt.Println("✓ Avoided creating 9+ classes (3 account types × 3 processors)")
	fmt.Println("✓ Easy to add new account types or processors independently")
	fmt.Println("✓ Processors can be switched at runtime")
	fmt.Println("✓ JoshBank can easily support new payment methods without modifying account classes")
}
