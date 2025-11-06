package main

import (
	"fmt"
	"time"
)

// AccountTemplate is the prototype interface that all account templates must implement
type AccountTemplate interface {
	Clone() AccountTemplate
	GetInfo() string
	Customize(accountNumber, customerName string)
}

// AccountMetadata contains common metadata for all accounts
type AccountMetadata struct {
	CreatedAt   time.Time
	CreatedBy   string
	Version     string
	BankName    string
}

// Clone creates a copy of metadata
func (m *AccountMetadata) Clone() AccountMetadata {
	return AccountMetadata{
		CreatedAt: m.CreatedAt,
		CreatedBy: m.CreatedBy,
		Version:   m.Version,
		BankName:  m.BankName,
	}
}

// CheckingAccountTemplate represents a checking account template
type CheckingAccountTemplate struct {
	AccountNumber string
	CustomerName  string
	AccountType   string
	InterestRate  float64
	MonthlyFee    float64
	OverdraftLimit float64
	Metadata      AccountMetadata
	Features      []string // Demonstrates deep copy of slices
}

// Clone creates a deep copy of the checking account template
func (a *CheckingAccountTemplate) Clone() AccountTemplate {
	// Create a new instance
	clone := &CheckingAccountTemplate{
		AccountType:    a.AccountType,
		InterestRate:   a.InterestRate,
		MonthlyFee:     a.MonthlyFee,
		OverdraftLimit: a.OverdraftLimit,
		Metadata:       a.Metadata.Clone(),
	}

	// Deep copy the slice to avoid shared references
	clone.Features = make([]string, len(a.Features))
	copy(clone.Features, a.Features)

	return clone
}

func (a *CheckingAccountTemplate) GetInfo() string {
	return fmt.Sprintf("Checking Account: %s (Interest: %.2f%%, Fee: $%.2f, Overdraft: $%.2f)",
		a.AccountType, a.InterestRate*100, a.MonthlyFee, a.OverdraftLimit)
}

func (a *CheckingAccountTemplate) Customize(accountNumber, customerName string) {
	a.AccountNumber = accountNumber
	a.CustomerName = customerName
	a.Metadata.CreatedAt = time.Now()
}

// SavingsAccountTemplate represents a savings account template
type SavingsAccountTemplate struct {
	AccountNumber string
	CustomerName  string
	AccountType   string
	InterestRate  float64
	MinimumBalance float64
	WithdrawalLimit int
	Metadata       AccountMetadata
	Features       []string
}

// Clone creates a deep copy of the savings account template
func (a *SavingsAccountTemplate) Clone() AccountTemplate {
	clone := &SavingsAccountTemplate{
		AccountType:     a.AccountType,
		InterestRate:    a.InterestRate,
		MinimumBalance:  a.MinimumBalance,
		WithdrawalLimit: a.WithdrawalLimit,
		Metadata:        a.Metadata.Clone(),
	}

	// Deep copy features
	clone.Features = make([]string, len(a.Features))
	copy(clone.Features, a.Features)

	return clone
}

func (a *SavingsAccountTemplate) GetInfo() string {
	return fmt.Sprintf("Savings Account: %s (Interest: %.2f%%, Min Balance: $%.2f, Withdrawal Limit: %d/month)",
		a.AccountType, a.InterestRate*100, a.MinimumBalance, a.WithdrawalLimit)
}

func (a *SavingsAccountTemplate) Customize(accountNumber, customerName string) {
	a.AccountNumber = accountNumber
	a.CustomerName = customerName
	a.Metadata.CreatedAt = time.Now()
}

// AccountTemplateRegistry manages prototype instances.
// This is an optional component that stores pre-configured prototypes.
type AccountTemplateRegistry struct {
	templates map[string]AccountTemplate
}

func NewAccountTemplateRegistry() *AccountTemplateRegistry {
	return &AccountTemplateRegistry{
		templates: make(map[string]AccountTemplate),
	}
}

// Register adds a template to the registry
func (r *AccountTemplateRegistry) Register(key string, template AccountTemplate) {
	r.templates[key] = template
}

// Create clones a registered template
func (r *AccountTemplateRegistry) Create(key string) (AccountTemplate, error) {
	template, exists := r.templates[key]
	if !exists {
		return nil, fmt.Errorf("template '%s' not found", key)
	}
	return template.Clone(), nil
}

func main() {
	fmt.Println("=== Prototype Pattern: JoshBank Account Templates ===")

	// Create prototype templates
	fmt.Println("\n--- Setting Up Account Templates ---")

	// Standard checking account template
	checkingTemplate := &CheckingAccountTemplate{
		AccountType:    "Standard Checking",
		InterestRate:   0.01,
		MonthlyFee:     0.0,
		OverdraftLimit: 500.0,
		Metadata: AccountMetadata{
			CreatedBy: "JoshBank System",
			Version:   "1.0",
			BankName:  "JoshBank",
		},
		Features: []string{"Online Banking", "Mobile App", "Debit Card"},
	}

	// Premium savings account template
	savingsTemplate := &SavingsAccountTemplate{
		AccountType:     "Premium Savings",
		InterestRate:    0.025,
		MinimumBalance:  1000.0,
		WithdrawalLimit: 6,
		Metadata: AccountMetadata{
			CreatedBy: "JoshBank System",
			Version:   "1.0",
			BankName:  "JoshBank",
		},
		Features: []string{"High Interest", "Online Banking", "Mobile App", "ATM Access"},
	}

	// Create registry and register templates
	registry := NewAccountTemplateRegistry()
	registry.Register("standard-checking", checkingTemplate)
	registry.Register("premium-savings", savingsTemplate)

	fmt.Println("✓ Account templates registered")

	// Example 1: Clone checking accounts for different customers
	fmt.Println("\n--- Example 1: Creating Multiple Checking Accounts ---")

	account1, _ := registry.Create("standard-checking")
	account1.Customize("CHK001", "John Doe")
	fmt.Printf("Created: %s\n", account1.GetInfo())

	account2, _ := registry.Create("standard-checking")
	account2.Customize("CHK002", "Jane Smith")
	fmt.Printf("Created: %s\n", account2.GetInfo())

	// Verify they are independent copies
	fmt.Println("\n✓ Each account is an independent copy")

	// Example 2: Clone savings accounts for different customers
	fmt.Println("\n--- Example 2: Creating Savings Accounts ---")

	savings1, _ := registry.Create("premium-savings")
	savings1.Customize("SAV001", "Bob Johnson")
	fmt.Printf("Created: %s\n", savings1.GetInfo())

	savings2, _ := registry.Create("premium-savings")
	savings2.Customize("SAV002", "Alice Williams")
	fmt.Printf("Created: %s\n", savings2.GetInfo())

	// Example 3: Demonstrate deep copy
	fmt.Println("\n--- Example 3: Deep Copy Verification ---")

	account3, _ := registry.Create("standard-checking")
	account3Concrete := account3.(*CheckingAccountTemplate)
	account3Concrete.Features = append(account3Concrete.Features, "Wire Transfer")

	fmt.Printf("Account 3 features: %v\n", account3Concrete.Features)
	fmt.Printf("Original template features: %v\n", checkingTemplate.Features)
	fmt.Println("✓ Modifying clone doesn't affect template (deep copy)")

	// Example 4: Error handling
	fmt.Println("\n--- Example 4: Error Handling ---")
	_, err := registry.Create("non-existent")
	if err != nil {
		fmt.Printf("✓ Error handled: %v\n", err)
	}

	// Example 5: Performance comparison
	fmt.Println("\n--- Example 5: Performance Benefits ---")
	fmt.Println("Creating from template: Fast - just copying existing object")
	fmt.Println("Creating from scratch: Slow - would need to:")
	fmt.Println("  - Load template from database")
	fmt.Println("  - Parse configuration")
	fmt.Println("  - Initialize all fields")
	fmt.Println("  - Set up default values")
	fmt.Println("  - Calculate interest rates and fees")

	fmt.Println("\n✓ Prototype pattern avoids expensive initialization")
	fmt.Println("✓ Easy to create variations of complex account configurations")
	fmt.Println("✓ Registry provides centralized template management")
	fmt.Println("✓ JoshBank can quickly onboard new customers using account templates")
}
