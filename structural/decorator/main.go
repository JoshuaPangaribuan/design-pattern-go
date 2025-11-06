package main

import "fmt"

// Transaction is the component interface that all transactions and decorators implement.
// This allows decorators to wrap other decorators or concrete transactions.
type Transaction interface {
	GetDescription() string
	GetAmount() float64
	Process() error
}

// --- Concrete Components (Base Transactions) ---

// BaseTransaction is a concrete component - a basic transaction
type BaseTransaction struct {
	description string
	amount      float64
}

func NewBaseTransaction(description string, amount float64) *BaseTransaction {
	return &BaseTransaction{
		description: description,
		amount:      amount,
	}
}

func (t *BaseTransaction) GetDescription() string {
	return t.description
}

func (t *BaseTransaction) GetAmount() float64 {
	return t.amount
}

func (t *BaseTransaction) Process() error {
	fmt.Printf("Processing transaction: %s - $%.2f\n", t.description, t.amount)
	return nil
}

// --- Base Decorator ---

// TransactionDecorator is the base decorator that wraps a Transaction.
// All concrete decorators will embed this.
type TransactionDecorator struct {
	transaction Transaction
}

// --- Concrete Decorators ---

// LoggingDecorator adds logging to transactions
type LoggingDecorator struct {
	TransactionDecorator
}

func NewLoggingDecorator(transaction Transaction) *LoggingDecorator {
	return &LoggingDecorator{
		TransactionDecorator: TransactionDecorator{transaction: transaction},
	}
}

func (d *LoggingDecorator) GetDescription() string {
	return d.transaction.GetDescription() + " [Logged]"
}

func (d *LoggingDecorator) GetAmount() float64 {
	return d.transaction.GetAmount()
}

func (d *LoggingDecorator) Process() error {
	fmt.Println("  [LOG] Transaction started")
	err := d.transaction.Process()
	if err != nil {
		fmt.Printf("  [LOG] Transaction failed: %v\n", err)
	} else {
		fmt.Println("  [LOG] Transaction completed successfully")
	}
	return err
}

// ValidationDecorator adds validation to transactions
type ValidationDecorator struct {
	TransactionDecorator
}

func NewValidationDecorator(transaction Transaction) *ValidationDecorator {
	return &ValidationDecorator{
		TransactionDecorator: TransactionDecorator{transaction: transaction},
	}
}

func (d *ValidationDecorator) GetDescription() string {
	return d.transaction.GetDescription() + " [Validated]"
}

func (d *ValidationDecorator) GetAmount() float64 {
	return d.transaction.GetAmount()
}

func (d *ValidationDecorator) Process() error {
	fmt.Println("  [VALIDATION] Validating transaction...")
	if d.transaction.GetAmount() <= 0 {
		return fmt.Errorf("invalid amount: must be positive")
	}
	if d.transaction.GetAmount() > 100000 {
		return fmt.Errorf("amount exceeds limit: requires approval")
	}
	fmt.Println("  [VALIDATION] Validation passed")
	return d.transaction.Process()
}

// EncryptionDecorator adds encryption to transactions
type EncryptionDecorator struct {
	TransactionDecorator
}

func NewEncryptionDecorator(transaction Transaction) *EncryptionDecorator {
	return &EncryptionDecorator{
		TransactionDecorator: TransactionDecorator{transaction: transaction},
	}
}

func (d *EncryptionDecorator) GetDescription() string {
	return d.transaction.GetDescription() + " [Encrypted]"
}

func (d *EncryptionDecorator) GetAmount() float64 {
	return d.transaction.GetAmount()
}

func (d *EncryptionDecorator) Process() error {
	fmt.Println("  [ENCRYPTION] Encrypting transaction data...")
	err := d.transaction.Process()
	fmt.Println("  [ENCRYPTION] Transaction data encrypted")
	return err
}

// AuditDecorator adds audit trail to transactions
type AuditDecorator struct {
	TransactionDecorator
}

func NewAuditDecorator(transaction Transaction) *AuditDecorator {
	return &AuditDecorator{
		TransactionDecorator: TransactionDecorator{transaction: transaction},
	}
}

func (d *AuditDecorator) GetDescription() string {
	return d.transaction.GetDescription() + " [Audited]"
}

func (d *AuditDecorator) GetAmount() float64 {
	return d.transaction.GetAmount()
}

func (d *AuditDecorator) Process() error {
	err := d.transaction.Process()
	fmt.Printf("  [AUDIT] Transaction recorded in audit log: %s - $%.2f\n", 
		d.transaction.GetDescription(), d.transaction.GetAmount())
	return err
}

// FeeDecorator adds fees to transactions
type FeeDecorator struct {
	TransactionDecorator
	feePercent float64
}

func NewFeeDecorator(transaction Transaction, feePercent float64) *FeeDecorator {
	return &FeeDecorator{
		TransactionDecorator: TransactionDecorator{transaction: transaction},
		feePercent:          feePercent,
	}
}

func (d *FeeDecorator) GetDescription() string {
	return d.transaction.GetDescription() + fmt.Sprintf(" [Fee: %.1f%%]", d.feePercent*100)
}

func (d *FeeDecorator) GetAmount() float64 {
	fee := d.transaction.GetAmount() * d.feePercent
	return d.transaction.GetAmount() + fee
}

func (d *FeeDecorator) Process() error {
	fee := d.transaction.GetAmount() * d.feePercent
	fmt.Printf("  [FEE] Applying %.1f%% fee: $%.2f\n", d.feePercent*100, fee)
	return d.transaction.Process()
}

// --- Helper Functions ---

func printTransaction(t Transaction) {
	fmt.Printf("\nTransaction: %s\n", t.GetDescription())
	fmt.Printf("Amount: $%.2f\n", t.GetAmount())
	fmt.Println("Processing:")
	if err := t.Process(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	fmt.Println("=== Decorator Pattern: JoshBank Transaction Processing ===")

	// Example 1: Simple transaction
	fmt.Println("\n--- Example 1: Simple Transaction ---")
	transaction1 := NewBaseTransaction("Payment to Merchant", 100.00)
	printTransaction(transaction1)

	// Example 2: Transaction with logging
	fmt.Println("\n--- Example 2: Transaction with Logging ---")
	transaction2 := Transaction(NewBaseTransaction("Transfer to Savings", 500.00))
	transaction2 = NewLoggingDecorator(transaction2)
	printTransaction(transaction2)

	// Example 3: Transaction with validation and logging
	fmt.Println("\n--- Example 3: Transaction with Validation and Logging ---")
	transaction3 := Transaction(NewBaseTransaction("Bill Payment", 75.50))
	transaction3 = NewValidationDecorator(transaction3)
	transaction3 = NewLoggingDecorator(transaction3)
	printTransaction(transaction3)

	// Example 4: Transaction with multiple decorators
	fmt.Println("\n--- Example 4: Transaction with Multiple Decorators ---")
	transaction4 := Transaction(NewBaseTransaction("Wire Transfer", 1000.00))
	transaction4 = NewValidationDecorator(transaction4)
	transaction4 = NewEncryptionDecorator(transaction4)
	transaction4 = NewAuditDecorator(transaction4)
	transaction4 = NewLoggingDecorator(transaction4)
	printTransaction(transaction4)

	// Example 5: Transaction with fee
	fmt.Println("\n--- Example 5: Transaction with Fee ---")
	transaction5 := Transaction(NewBaseTransaction("International Transfer", 200.00))
	transaction5 = NewFeeDecorator(transaction5, 0.02) // 2% fee
	transaction5 = NewValidationDecorator(transaction5)
	transaction5 = NewLoggingDecorator(transaction5)
	printTransaction(transaction5)

	// Example 6: Building transaction step by step
	fmt.Println("\n--- Example 6: Building Transaction Step by Step ---")

	base := NewBaseTransaction("Custom Transaction", 250.00)
	fmt.Printf("Starting with: %s - $%.2f\n", base.GetDescription(), base.GetAmount())

	withValidation := NewValidationDecorator(base)
	fmt.Printf("Added Validation: %s - $%.2f\n", withValidation.GetDescription(), withValidation.GetAmount())

	withEncryption := NewEncryptionDecorator(withValidation)
	fmt.Printf("Added Encryption: %s - $%.2f\n", withEncryption.GetDescription(), withEncryption.GetAmount())

	withAudit := NewAuditDecorator(withEncryption)
	fmt.Printf("Added Audit: %s - $%.2f\n", withAudit.GetDescription(), withAudit.GetAmount())

	fmt.Println("\nProcessing final transaction:")
	withAudit.Process()

	// Example 7: Demonstrating flexibility
	fmt.Println("\n--- Example 7: Same Base, Different Decorations ---")

	baseTransaction := NewBaseTransaction("Payment", 150.00)

	option1 := NewLoggingDecorator(baseTransaction)
	option2 := NewValidationDecorator(NewLoggingDecorator(baseTransaction))
	option3 := NewEncryptionDecorator(NewAuditDecorator(baseTransaction))

	fmt.Println("Three different ways to enhance the same transaction:")
	printTransaction(option1)
	printTransaction(option2)
	printTransaction(option3)

	fmt.Println("\n✓ Decorator pattern allows dynamic addition of responsibilities")
	fmt.Println("✓ Decorators can be stacked in any combination")
	fmt.Println("✓ No need for subclasses for every possible combination")
	fmt.Println("✓ Follows Open/Closed Principle - open for extension, closed for modification")
	fmt.Println("✓ JoshBank can add features like logging, validation, encryption without modifying core transaction code")
}
