package main

import (
	"fmt"
)

// CurrencyInfo is the flyweight that stores intrinsic state (shared data).
// Currency code, exchange rate, and symbol are shared among many transactions.
type CurrencyInfo struct {
	code       string
	symbol     string
	exchangeRate float64 // Exchange rate to USD
	decimalPlaces int
}

// FormatAmount formats an amount using this currency's formatting rules.
// Amount is extrinsic state (unique per transaction).
func (ci *CurrencyInfo) FormatAmount(amount float64) string {
	formatted := fmt.Sprintf("%.2f", amount)
	if ci.decimalPlaces == 0 {
		formatted = fmt.Sprintf("%.0f", amount)
	}
	return fmt.Sprintf("%s%s", ci.symbol, formatted)
}

func (ci *CurrencyInfo) GetDescription() string {
	return fmt.Sprintf("%s (%s, Rate: %.4f)", ci.code, ci.symbol, ci.exchangeRate)
}

// CurrencyFactory is the flyweight factory that manages currency information.
// It ensures that currencies with the same code are shared.
type CurrencyFactory struct {
	currencies map[string]*CurrencyInfo
}

func NewCurrencyFactory() *CurrencyFactory {
	factory := &CurrencyFactory{
		currencies: make(map[string]*CurrencyInfo),
	}
	// Pre-populate with common currencies
	factory.initializeCurrencies()
	return factory
}

func (cf *CurrencyFactory) initializeCurrencies() {
	currencies := []*CurrencyInfo{
		{code: "USD", symbol: "$", exchangeRate: 1.0, decimalPlaces: 2},
		{code: "EUR", symbol: "€", exchangeRate: 0.85, decimalPlaces: 2},
		{code: "GBP", symbol: "£", exchangeRate: 0.75, decimalPlaces: 2},
		{code: "JPY", symbol: "¥", exchangeRate: 110.0, decimalPlaces: 0},
		{code: "CNY", symbol: "¥", exchangeRate: 6.5, decimalPlaces: 2},
	}
	for _, curr := range currencies {
		cf.currencies[curr.code] = curr
	}
}

// GetCurrency returns a shared currency or creates a new one if it doesn't exist.
// This is the key method that enables flyweight sharing.
func (cf *CurrencyFactory) GetCurrency(code string) (*CurrencyInfo, error) {
	// Check if we already have this currency
	if currency, exists := cf.currencies[code]; exists {
		return currency, nil
	}

	// Create new currency if it doesn't exist (with default values)
	currency := &CurrencyInfo{
		code:          code,
		symbol:        code,
		exchangeRate:  1.0,
		decimalPlaces: 2,
	}
	cf.currencies[code] = currency
	fmt.Printf("  [Creating new currency: %s]\n", code)
	return currency, nil
}

func (cf *CurrencyFactory) GetCurrencyCount() int {
	return len(cf.currencies)
}

// Transaction represents a transaction with its unique amount and currency reference.
// This stores the extrinsic state (amount) and references the flyweight.
type Transaction struct {
	id       string
	amount   float64
	currency *CurrencyInfo // Reference to shared flyweight
}

func NewTransaction(id string, amount float64, currency *CurrencyInfo) *Transaction {
	return &Transaction{
		id:       id,
		amount:   amount,
		currency: currency,
	}
}

func (t *Transaction) Display() {
	formattedAmount := t.currency.FormatAmount(t.amount)
	fmt.Printf("Transaction %s: %s (%s)\n", t.id, formattedAmount, t.currency.code)
}

func (t *Transaction) GetAmountInUSD() float64 {
	return t.amount / t.currency.exchangeRate
}

// TransactionLedger manages a collection of transactions
type TransactionLedger struct {
	transactions []*Transaction
	factory      *CurrencyFactory
}

func NewTransactionLedger() *TransactionLedger {
	return &TransactionLedger{
		transactions: make([]*Transaction, 0),
		factory:      NewCurrencyFactory(),
	}
}

func (tl *TransactionLedger) AddTransaction(id string, amount float64, currencyCode string) error {
	currency, err := tl.factory.GetCurrency(currencyCode)
	if err != nil {
		return err
	}
	transaction := NewTransaction(id, amount, currency)
	tl.transactions = append(tl.transactions, transaction)
	return nil
}

func (tl *TransactionLedger) DisplayAll() {
	fmt.Println("\n--- Transaction Ledger ---")
	for _, txn := range tl.transactions {
		txn.Display()
	}
}

func (tl *TransactionLedger) GetTotalInUSD() float64 {
	var total float64
	for _, txn := range tl.transactions {
		total += txn.GetAmountInUSD()
	}
	return total
}

func (tl *TransactionLedger) GetStats() {
	fmt.Printf("\nLedger Statistics:\n")
	fmt.Printf("  Total transactions: %d\n", len(tl.transactions))
	fmt.Printf("  Unique currencies: %d\n", tl.factory.GetCurrencyCount())
	fmt.Printf("  Memory saved: ~%.1f%% (without flyweight would need %d currency objects)\n",
		(1.0-float64(tl.factory.GetCurrencyCount())/float64(len(tl.transactions)))*100,
		len(tl.transactions))
}

func main() {
	fmt.Println("=== Flyweight Pattern: JoshBank Currency Management ===")

	ledger := NewTransactionLedger()

	// Example 1: Adding transactions with same currency
	fmt.Println("\n--- Example 1: Adding Transactions with Same Currency ---")
	ledger.AddTransaction("TXN001", 100.50, "USD")
	ledger.AddTransaction("TXN002", 250.00, "USD")
	ledger.AddTransaction("TXN003", 75.25, "USD")
	// Notice: All USD transactions share the same CurrencyInfo object

	// Example 2: Adding transactions with different currencies
	fmt.Println("\n--- Example 2: Adding Transactions with Different Currencies ---")
	ledger.AddTransaction("TXN004", 500.00, "EUR")
	ledger.AddTransaction("TXN005", 300.00, "GBP")
	ledger.AddTransaction("TXN006", 10000.00, "JPY")

	// Example 3: Reusing existing currencies
	fmt.Println("\n--- Example 3: Reusing Existing Currencies ---")
	ledger.AddTransaction("TXN007", 150.00, "USD") // Reuses USD currency
	ledger.AddTransaction("TXN008", 200.00, "EUR") // Reuses EUR currency

	// Example 4: New currency
	fmt.Println("\n--- Example 4: Adding New Currency ---")
	ledger.AddTransaction("TXN009", 1000.00, "CNY")

	// Display all transactions
	ledger.DisplayAll()

	// Show statistics
	ledger.GetStats()

	// Example 5: Demonstrate memory savings
	fmt.Println("\n--- Example 5: Memory Savings Demonstration ---")

	largeLedger := NewTransactionLedger()

	// Simulate a ledger with 1000 transactions using only 5 different currencies
	currencyCodes := []string{"USD", "EUR", "GBP", "JPY", "CNY"}
	amounts := []float64{100.0, 250.0, 500.0, 1000.0, 2000.0}

	fmt.Println("Creating ledger with 1000 transactions...")
	for i := 0; i < 1000; i++ {
		currencyCode := currencyCodes[i%len(currencyCodes)]
		amount := amounts[i%len(amounts)]
		largeLedger.AddTransaction(fmt.Sprintf("TXN%04d", i+1), amount, currencyCode)
	}

	largeLedger.GetStats()
	fmt.Printf("Total amount in USD: $%.2f\n", largeLedger.GetTotalInUSD())

	fmt.Println("\n✓ Flyweight pattern significantly reduces memory usage")
	fmt.Println("✓ Shared intrinsic state (currency info) among many transactions")
	fmt.Println("✓ Extrinsic state (amount) remains unique per transaction")
	fmt.Println("✓ Factory ensures proper sharing of flyweight objects")
	fmt.Println("✓ JoshBank can handle millions of transactions efficiently")
}


