package main

import "fmt"

// InterestCalculationStrategy defines the interface for interest calculation algorithms
type InterestCalculationStrategy interface {
	CalculateInterest(balance float64) float64
	GetName() string
}

// --- Concrete Strategies ---

type SimpleInterestStrategy struct {
	rate float64
}

func NewSimpleInterestStrategy(rate float64) *SimpleInterestStrategy {
	return &SimpleInterestStrategy{rate: rate}
}

func (s *SimpleInterestStrategy) CalculateInterest(balance float64) float64 {
	return balance * s.rate
}

func (s *SimpleInterestStrategy) GetName() string {
	return fmt.Sprintf("Simple Interest (%.2f%%)", s.rate*100)
}

type CompoundInterestStrategy struct {
	rate     float64
	periods  int
}

func NewCompoundInterestStrategy(rate float64, periods int) *CompoundInterestStrategy {
	return &CompoundInterestStrategy{rate: rate, periods: periods}
}

func (c *CompoundInterestStrategy) CalculateInterest(balance float64) float64 {
	return balance * (c.rate * float64(c.periods))
}

func (c *CompoundInterestStrategy) GetName() string {
	return fmt.Sprintf("Compound Interest (%.2f%%, %d periods)", c.rate*100, c.periods)
}

type TieredInterestStrategy struct {
	tiers []struct {
		minBalance float64
		rate       float64
	}
}

func NewTieredInterestStrategy() *TieredInterestStrategy {
	return &TieredInterestStrategy{
		tiers: []struct {
			minBalance float64
			rate       float64
		}{
			{0, 0.01},      // 1% for balances 0-1000
			{1000, 0.02},   // 2% for balances 1000-10000
			{10000, 0.03},  // 3% for balances 10000+
		},
	}
}

func (t *TieredInterestStrategy) CalculateInterest(balance float64) float64 {
	for i := len(t.tiers) - 1; i >= 0; i-- {
		if balance >= t.tiers[i].minBalance {
			return balance * t.tiers[i].rate
		}
	}
	return 0
}

func (t *TieredInterestStrategy) GetName() string {
	return "Tiered Interest (1-3% based on balance)"
}

// Account is the context that uses a strategy
type Account struct {
	accountID string
	balance   float64
	strategy  InterestCalculationStrategy
}

func NewAccount(accountID string, balance float64, strategy InterestCalculationStrategy) *Account {
	return &Account{
		accountID: accountID,
		balance:   balance,
		strategy:  strategy,
	}
}

func (a *Account) SetStrategy(strategy InterestCalculationStrategy) {
	a.strategy = strategy
	fmt.Printf("  [Account %s] Strategy changed to: %s\n", a.accountID, strategy.GetName())
}

func (a *Account) CalculateInterest() {
	fmt.Printf("\n→ Calculating interest for account %s (balance: $%.2f) using %s:\n",
		a.accountID, a.balance, a.strategy.GetName())
	interest := a.strategy.CalculateInterest(a.balance)
	fmt.Printf("  Interest: $%.2f\n", interest)
}

// --- Another Example: Fee Calculation Strategies ---

type FeeCalculationStrategy interface {
	CalculateFee(amount float64) float64
}

type FlatFeeStrategy struct {
	fee float64
}

func (f *FlatFeeStrategy) CalculateFee(amount float64) float64 {
	return f.fee
}

type PercentageFeeStrategy struct {
	rate float64
}

func (p *PercentageFeeStrategy) CalculateFee(amount float64) float64 {
	return amount * p.rate
}

type TransactionProcessor struct {
	feeStrategy FeeCalculationStrategy
}

func (t *TransactionProcessor) SetFeeStrategy(strategy FeeCalculationStrategy) {
	t.feeStrategy = strategy
}

func (t *TransactionProcessor) ProcessTransaction(amount float64) {
	fee := t.feeStrategy.CalculateFee(amount)
	fmt.Printf("  Transaction amount: $%.2f, Fee: $%.2f, Total: $%.2f\n", amount, fee, amount+fee)
}

func main() {
	fmt.Println("=== Strategy Pattern: JoshBank Interest & Fee Calculation ===")

	// Example 1: Interest calculation strategies
	fmt.Println("\n--- Example 1: Interest Calculation Strategies ---")

	account := NewAccount("ACC001", 5000.0, NewSimpleInterestStrategy(0.02))
	account.CalculateInterest()

	account.SetStrategy(NewCompoundInterestStrategy(0.015, 12))
	account.CalculateInterest()

	account.SetStrategy(NewTieredInterestStrategy())
	account.CalculateInterest()

	// Example 2: Fee calculation strategies
	fmt.Println("\n--- Example 2: Fee Calculation Strategies ---")

	processor := &TransactionProcessor{}

	fmt.Println("\n→ Processing with Flat Fee:")
	processor.SetFeeStrategy(&FlatFeeStrategy{fee: 2.50})
	processor.ProcessTransaction(100.0)
	processor.ProcessTransaction(1000.0)

	fmt.Println("\n→ Processing with Percentage Fee:")
	processor.SetFeeStrategy(&PercentageFeeStrategy{rate: 0.025})
	processor.ProcessTransaction(100.0)
	processor.ProcessTransaction(1000.0)

	fmt.Println("\n✓ Strategy pattern defines family of algorithms")
	fmt.Println("✓ Makes algorithms interchangeable")
	fmt.Println("✓ Eliminates conditional statements")
	fmt.Println("✓ Easy to add new calculation strategies")
	fmt.Println("✓ JoshBank can switch between different interest and fee calculation methods")
}
