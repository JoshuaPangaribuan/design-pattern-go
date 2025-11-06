package main

import "fmt"

// Account is the abstract product interface for bank accounts
type Account interface {
	GetAccountType() string
	GetInterestRate() float64
	GetMinimumBalance() float64
}

// Card is the abstract product interface for payment cards
type Card interface {
	GetCardType() string
	GetCreditLimit() float64
	GetAnnualFee() float64
}

// Loan is the abstract product interface for loans
type Loan interface {
	GetLoanType() string
	GetInterestRate() float64
	GetMaxAmount() float64
}

// --- Personal Banking Products ---

// PersonalAccount is a concrete product for personal banking
type PersonalAccount struct {
	accountType string
}

func (a *PersonalAccount) GetAccountType() string {
	return "Personal Checking Account"
}

func (a *PersonalAccount) GetInterestRate() float64 {
	return 0.01 // 1% annual interest
}

func (a *PersonalAccount) GetMinimumBalance() float64 {
	return 0.0 // No minimum balance
}

// PersonalCard is a concrete product for personal banking
type PersonalCard struct {
	cardType string
}

func (c *PersonalCard) GetCardType() string {
	return "Personal Debit Card"
}

func (c *PersonalCard) GetCreditLimit() float64 {
	return 0.0 // Debit card, no credit limit
}

func (c *PersonalCard) GetAnnualFee() float64 {
	return 0.0 // No annual fee
}

// PersonalLoan is a concrete product for personal banking
type PersonalLoan struct {
	loanType string
}

func (l *PersonalLoan) GetLoanType() string {
	return "Personal Loan"
}

func (l *PersonalLoan) GetInterestRate() float64 {
	return 0.08 // 8% annual interest
}

func (l *PersonalLoan) GetMaxAmount() float64 {
	return 50000.0 // Max $50,000
}

// --- Business Banking Products ---

// BusinessAccount is a concrete product for business banking
type BusinessAccount struct {
	accountType string
}

func (a *BusinessAccount) GetAccountType() string {
	return "Business Checking Account"
}

func (a *BusinessAccount) GetInterestRate() float64 {
	return 0.02 // 2% annual interest
}

func (a *BusinessAccount) GetMinimumBalance() float64 {
	return 5000.0 // $5,000 minimum balance
}

// BusinessCard is a concrete product for business banking
type BusinessCard struct {
	cardType string
}

func (c *BusinessCard) GetCardType() string {
	return "Business Credit Card"
}

func (c *BusinessCard) GetCreditLimit() float64 {
	return 100000.0 // $100,000 credit limit
}

func (c *BusinessCard) GetAnnualFee() float64 {
	return 500.0 // $500 annual fee
}

// BusinessLoan is a concrete product for business banking
type BusinessLoan struct {
	loanType string
}

func (l *BusinessLoan) GetLoanType() string {
	return "Business Loan"
}

func (l *BusinessLoan) GetInterestRate() float64 {
	return 0.06 // 6% annual interest
}

func (l *BusinessLoan) GetMaxAmount() float64 {
	return 1000000.0 // Max $1,000,000
}

// BankingProductFactory is the abstract factory interface.
// It declares methods for creating each type of banking product.
type BankingProductFactory interface {
	CreateAccount() Account
	CreateCard() Card
	CreateLoan() Loan
	GetProductFamily() string
}

// PersonalBankingFactory is a concrete factory that creates personal banking products
type PersonalBankingFactory struct{}

func (f *PersonalBankingFactory) CreateAccount() Account {
	return &PersonalAccount{}
}

func (f *PersonalBankingFactory) CreateCard() Card {
	return &PersonalCard{}
}

func (f *PersonalBankingFactory) CreateLoan() Loan {
	return &PersonalLoan{}
}

func (f *PersonalBankingFactory) GetProductFamily() string {
	return "Personal Banking"
}

// BusinessBankingFactory is a concrete factory that creates business banking products
type BusinessBankingFactory struct{}

func (f *BusinessBankingFactory) CreateAccount() Account {
	return &BusinessAccount{}
}

func (f *BusinessBankingFactory) CreateCard() Card {
	return &BusinessCard{}
}

func (f *BusinessBankingFactory) CreateLoan() Loan {
	return &BusinessLoan{}
}

func (f *BusinessBankingFactory) GetProductFamily() string {
	return "Business Banking"
}

// BankingApplication represents the client code that uses the factory.
// It works with factories and products through their interfaces.
type BankingApplication struct {
	factory BankingProductFactory
}

func NewBankingApplication(factory BankingProductFactory) *BankingApplication {
	return &BankingApplication{factory: factory}
}

// CreateCustomerPackage demonstrates using the factory to create a complete banking package
func (a *BankingApplication) CreateCustomerPackage() {
	fmt.Printf("\n=== Creating Banking Package (%s) ===\n", a.factory.GetProductFamily())

	// Create all products using the same factory
	// This ensures all products belong to the same banking family
	account := a.factory.CreateAccount()
	card := a.factory.CreateCard()
	loan := a.factory.CreateLoan()

	// Display the package details
	fmt.Printf("Account: %s (Interest: %.1f%%, Min Balance: $%.2f)\n",
		account.GetAccountType(), account.GetInterestRate()*100, account.GetMinimumBalance())
	fmt.Printf("Card: %s (Credit Limit: $%.2f, Annual Fee: $%.2f)\n",
		card.GetCardType(), card.GetCreditLimit(), card.GetAnnualFee())
	fmt.Printf("Loan: %s (Interest: %.1f%%, Max Amount: $%.2f)\n",
		loan.GetLoanType(), loan.GetInterestRate()*100, loan.GetMaxAmount())
}

// GetFactory returns the appropriate factory based on customer type
func GetFactory(customerType string) BankingProductFactory {
	switch customerType {
	case "business":
		return &BusinessBankingFactory{}
	case "personal":
		return &PersonalBankingFactory{}
	default:
		return &PersonalBankingFactory{}
	}
}

func main() {
	fmt.Println("=== Abstract Factory Pattern: JoshBank Product Families ===")

	// Scenario 1: Personal customer
	fmt.Println("\n--- Customer Type: Personal ---")
	personalFactory := GetFactory("personal")
	personalApp := NewBankingApplication(personalFactory)
	personalApp.CreateCustomerPackage()

	// Scenario 2: Business customer
	fmt.Println("\n--- Customer Type: Business ---")
	businessFactory := GetFactory("business")
	businessApp := NewBankingApplication(businessFactory)
	businessApp.CreateCustomerPackage()

	// Scenario 3: Switching customer types at runtime
	fmt.Println("\n--- Switching Customer Type at Runtime ---")
	app := NewBankingApplication(personalFactory)
	fmt.Println("Initial customer type: Personal")
	app.CreateCustomerPackage()

	// Switch to business
	app.factory = businessFactory
	fmt.Println("\nSwitched to Business customer:")
	app.CreateCustomerPackage()

	fmt.Println("\n✓ Abstract Factory ensures all products match the customer type")
	fmt.Println("✓ Easy to add new banking product families without changing client code")
	fmt.Println("✓ Products from the same family are guaranteed to be compatible")
	fmt.Println("✓ JoshBank can easily extend to new customer segments (e.g., Premium)")
}
