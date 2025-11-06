package main

import "fmt"

// Visitor interface declares visit methods for each account type
type Visitor interface {
	VisitCheckingAccount(account *CheckingAccount) float64
	VisitSavingsAccount(account *SavingsAccount) float64
	VisitInvestmentAccount(account *InvestmentAccount) float64
}

// Account interface declares accept method
type Account interface {
	Accept(visitor Visitor) float64
	GetAccountID() string
	GetBalance() float64
}

// --- Concrete Account Types ---

type CheckingAccount struct {
	accountID string
	balance   float64
}

func (c *CheckingAccount) Accept(visitor Visitor) float64 {
	return visitor.VisitCheckingAccount(c)
}

func (c *CheckingAccount) GetAccountID() string {
	return c.accountID
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}

type SavingsAccount struct {
	accountID string
	balance   float64
}

func (s *SavingsAccount) Accept(visitor Visitor) float64 {
	return visitor.VisitSavingsAccount(s)
}

func (s *SavingsAccount) GetAccountID() string {
	return s.accountID
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.balance
}

type InvestmentAccount struct {
	accountID string
	balance   float64
}

func (i *InvestmentAccount) Accept(visitor Visitor) float64 {
	return visitor.VisitInvestmentAccount(i)
}

func (i *InvestmentAccount) GetAccountID() string {
	return i.accountID
}

func (i *InvestmentAccount) GetBalance() float64 {
	return i.balance
}

// --- Concrete Visitors ---

type InterestCalculationVisitor struct{}

func (i *InterestCalculationVisitor) VisitCheckingAccount(account *CheckingAccount) float64 {
	interest := account.balance * 0.01 // 1% interest
	fmt.Printf("  [Interest] Checking Account %s: $%.2f (1%%)\n", account.accountID, interest)
	return interest
}

func (i *InterestCalculationVisitor) VisitSavingsAccount(account *SavingsAccount) float64 {
	interest := account.balance * 0.025 // 2.5% interest
	fmt.Printf("  [Interest] Savings Account %s: $%.2f (2.5%%)\n", account.accountID, interest)
	return interest
}

func (i *InterestCalculationVisitor) VisitInvestmentAccount(account *InvestmentAccount) float64 {
	interest := account.balance * 0.05 // 5% interest
	fmt.Printf("  [Interest] Investment Account %s: $%.2f (5%%)\n", account.accountID, interest)
	return interest
}

type FeeCalculationVisitor struct{}

func (f *FeeCalculationVisitor) VisitCheckingAccount(account *CheckingAccount) float64 {
	fee := 0.0 // No fee for checking accounts
	fmt.Printf("  [Fee] Checking Account %s: $%.2f (no fee)\n", account.accountID, fee)
	return fee
}

func (f *FeeCalculationVisitor) VisitSavingsAccount(account *SavingsAccount) float64 {
	fee := 5.0 // Flat fee for savings accounts
	fmt.Printf("  [Fee] Savings Account %s: $%.2f (flat fee)\n", account.accountID, fee)
	return fee
}

func (f *FeeCalculationVisitor) VisitInvestmentAccount(account *InvestmentAccount) float64 {
	fee := account.balance * 0.01 // 1% fee for investment accounts
	fmt.Printf("  [Fee] Investment Account %s: $%.2f (1%%)\n", account.accountID, fee)
	return fee
}

type RiskAssessmentVisitor struct{}

func (r *RiskAssessmentVisitor) VisitCheckingAccount(account *CheckingAccount) float64 {
	risk := 0.1 // Low risk
	fmt.Printf("  [Risk] Checking Account %s: Risk Level %.1f (Low)\n", account.accountID, risk)
	return risk
}

func (r *RiskAssessmentVisitor) VisitSavingsAccount(account *SavingsAccount) float64 {
	risk := 0.2 // Low-Medium risk
	fmt.Printf("  [Risk] Savings Account %s: Risk Level %.1f (Low-Medium)\n", account.accountID, risk)
	return risk
}

func (r *RiskAssessmentVisitor) VisitInvestmentAccount(account *InvestmentAccount) float64 {
	risk := 0.7 // High risk
	fmt.Printf("  [Risk] Investment Account %s: Risk Level %.1f (High)\n", account.accountID, risk)
	return risk
}

// AccountPortfolio manages accounts
type AccountPortfolio struct {
	accounts []Account
}

func NewAccountPortfolio() *AccountPortfolio {
	return &AccountPortfolio{accounts: make([]Account, 0)}
}

func (p *AccountPortfolio) AddAccount(account Account) {
	p.accounts = append(p.accounts, account)
}

func (p *AccountPortfolio) CalculateTotal(visitor Visitor) float64 {
	total := 0.0
	for _, account := range p.accounts {
		total += account.Accept(visitor)
	}
	return total
}

func (p *AccountPortfolio) ShowAccounts() {
	fmt.Println("\nPortfolio Accounts:")
	for _, account := range p.accounts {
		fmt.Printf("  - %s: $%.2f\n", account.GetAccountID(), account.GetBalance())
	}
}

func main() {
	fmt.Println("=== Visitor Pattern: JoshBank Account Analysis ===")

	// Create account portfolio
	portfolio := NewAccountPortfolio()
	portfolio.AddAccount(&CheckingAccount{accountID: "CHK001", balance: 5000.0})
	portfolio.AddAccount(&SavingsAccount{accountID: "SAV001", balance: 10000.0})
	portfolio.AddAccount(&InvestmentAccount{accountID: "INV001", balance: 50000.0})

	portfolio.ShowAccounts()

	// Example 1: Calculate interest
	fmt.Println("\n--- Example 1: Interest Calculation ---")
	interestVisitor := &InterestCalculationVisitor{}
	totalInterest := portfolio.CalculateTotal(interestVisitor)
	fmt.Printf("Total Interest: $%.2f\n", totalInterest)

	// Example 2: Calculate fees
	fmt.Println("\n--- Example 2: Fee Calculation ---")
	feeVisitor := &FeeCalculationVisitor{}
	totalFees := portfolio.CalculateTotal(feeVisitor)
	fmt.Printf("Total Fees: $%.2f\n", totalFees)

	// Example 3: Risk assessment
	fmt.Println("\n--- Example 3: Risk Assessment ---")
	riskVisitor := &RiskAssessmentVisitor{}
	portfolio.CalculateTotal(riskVisitor)

	// Example 4: Calculate net value
	fmt.Println("\n--- Example 4: Net Portfolio Value ---")
	totalBalance := 0.0
	for _, account := range portfolio.accounts {
		totalBalance += account.GetBalance()
	}
	netValue := totalBalance + totalInterest - totalFees
	fmt.Printf("Total Balance: $%.2f\n", totalBalance)
	fmt.Printf("Total Interest: $%.2f\n", totalInterest)
	fmt.Printf("Total Fees: $%.2f\n", totalFees)
	fmt.Printf("Net Value: $%.2f\n", netValue)

	fmt.Println("\n✓ Visitor pattern adds operations without modifying account types")
	fmt.Println("✓ Separates algorithms from account structure")
	fmt.Println("✓ Easy to add new analysis operations")
	fmt.Println("✓ Operations are centralized in visitor classes")
	fmt.Println("✓ JoshBank can perform various analyses on accounts without modifying account classes")
}
