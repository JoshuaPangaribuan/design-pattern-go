package main

import (
	"fmt"
	"strings"
)

// TransactionNode is the component interface that both individual transactions and transaction groups implement.
// This allows treating individual transactions and transaction groups uniformly.
type TransactionNode interface {
	GetID() string
	GetAmount() float64
	Display(indent string)
	GetTotalAmount() float64
	Search(transactionID string) TransactionNode
}

// --- Leaf: Individual Transaction ---

// Transaction represents a leaf node (individual transaction)
type Transaction struct {
	id          string
	amount      float64
	description string
	timestamp   string
}

func NewTransaction(id string, amount float64, description string, timestamp string) *Transaction {
	return &Transaction{
		id:          id,
		amount:      amount,
		description: description,
		timestamp:   timestamp,
	}
}

func (t *Transaction) GetID() string {
	return t.id
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}

func (t *Transaction) Display(indent string) {
	fmt.Printf("%s💳 %s: $%.2f - %s (%s)\n", indent, t.id, t.amount, t.description, t.timestamp)
}

func (t *Transaction) GetTotalAmount() float64 {
	return t.amount
}

func (t *Transaction) Search(transactionID string) TransactionNode {
	if t.id == transactionID {
		return t
	}
	return nil
}

// --- Composite: Transaction Group ---

// TransactionGroup represents a composite node that can contain transactions and other groups
type TransactionGroup struct {
	name        string
	transactions []TransactionNode
}

func NewTransactionGroup(name string) *TransactionGroup {
	return &TransactionGroup{
		name:        name,
		transactions: make([]TransactionNode, 0),
	}
}

func (g *TransactionGroup) GetID() string {
	return g.name
}

func (g *TransactionGroup) GetAmount() float64 {
	return g.GetTotalAmount()
}

// GetTotalAmount recursively calculates total amount of all children
func (g *TransactionGroup) GetTotalAmount() float64 {
	var total float64
	for _, transaction := range g.transactions {
		total += transaction.GetTotalAmount()
	}
	return total
}

// Display recursively displays the transaction group tree
func (g *TransactionGroup) Display(indent string) {
	fmt.Printf("%s📁 %s (Total: $%.2f)\n", indent, g.name, g.GetTotalAmount())
	for _, transaction := range g.transactions {
		transaction.Display(indent + "  ")
	}
}

// Search recursively searches for a transaction by ID
func (g *TransactionGroup) Search(transactionID string) TransactionNode {
	for _, transaction := range g.transactions {
		if result := transaction.Search(transactionID); result != nil {
			return result
		}
	}
	return nil
}

// Add adds a transaction or group to this group
func (g *TransactionGroup) Add(node TransactionNode) {
	g.transactions = append(g.transactions, node)
}

// Remove removes a transaction from this group
func (g *TransactionGroup) Remove(transactionID string) bool {
	for i, transaction := range g.transactions {
		if transaction.GetID() == transactionID {
			g.transactions = append(g.transactions[:i], g.transactions[i+1:]...)
			return true
		}
	}
	return false
}

// GetTransactions returns all transactions (useful for additional operations)
func (g *TransactionGroup) GetTransactions() []TransactionNode {
	return g.transactions
}

// --- Helper Functions ---

// printSeparator prints a visual separator
func printSeparator(title string) {
	fmt.Printf("\n%s %s %s\n", strings.Repeat("=", 20), title, strings.Repeat("=", 20))
}

func main() {
	fmt.Println("=== Composite Pattern: JoshBank Transaction Groups ===")

	// Build a transaction structure
	printSeparator("Building Transaction Structure")

	// Root: All transactions
	allTransactions := NewTransactionGroup("All Transactions")

	// January transactions
	january := NewTransactionGroup("January 2024")
	january.Add(NewTransaction("TXN001", 100.50, "Grocery Store", "2024-01-05"))
	january.Add(NewTransaction("TXN002", 45.00, "Gas Station", "2024-01-10"))
	january.Add(NewTransaction("TXN003", 250.00, "Restaurant", "2024-01-15"))

	// February transactions
	february := NewTransactionGroup("February 2024")
	february.Add(NewTransaction("TXN004", 75.25, "Grocery Store", "2024-02-03"))
	february.Add(NewTransaction("TXN005", 120.00, "Online Purchase", "2024-02-08"))
	february.Add(NewTransaction("TXN006", 30.00, "Coffee Shop", "2024-02-12"))

	// Categories
	food := NewTransactionGroup("Food & Dining")
	food.Add(NewTransaction("TXN007", 85.00, "Restaurant", "2024-01-20"))
	food.Add(NewTransaction("TXN008", 15.50, "Fast Food", "2024-02-05"))

	shopping := NewTransactionGroup("Shopping")
	shopping.Add(NewTransaction("TXN009", 200.00, "Department Store", "2024-01-25"))
	shopping.Add(NewTransaction("TXN010", 50.00, "Bookstore", "2024-02-10"))

	// Add categories to months
	january.Add(food)
	february.Add(shopping)

	// Add months to root
	allTransactions.Add(january)
	allTransactions.Add(february)

	fmt.Println("✓ Transaction structure created")

	// Example 1: Display entire tree
	printSeparator("Example 1: Display Transaction Tree")
	allTransactions.Display("")

	// Example 2: Calculate totals
	printSeparator("Example 2: Calculate Totals")
	fmt.Printf("Total amount (all transactions): $%.2f\n", allTransactions.GetTotalAmount())
	fmt.Printf("Total amount (January): $%.2f\n", january.GetTotalAmount())
	fmt.Printf("Total amount (February): $%.2f\n", february.GetTotalAmount())
	fmt.Printf("Total amount (Food & Dining): $%.2f\n", food.GetTotalAmount())
	fmt.Printf("Amount of single transaction (TXN001): $%.2f\n", allTransactions.Search("TXN001").GetTotalAmount())

	// Example 3: Search for transactions
	printSeparator("Example 3: Search Operations")

	searchTerms := []string{"TXN001", "TXN005", "TXN009", "TXN999"}
	for _, term := range searchTerms {
		result := allTransactions.Search(term)
		if result != nil {
			fmt.Printf("✓ Found '%s' (amount: $%.2f)\n", term, result.GetTotalAmount())
		} else {
			fmt.Printf("✗ '%s' not found\n", term)
		}
	}

	// Example 4: Modify structure
	printSeparator("Example 4: Modify Structure")

	fmt.Println("Adding new transaction to January...")
	january.Add(NewTransaction("TXN011", 60.00, "Pharmacy", "2024-01-28"))

	fmt.Println("Removing transaction from February...")
	february.Remove("TXN006")

	fmt.Println("\nUpdated structure:")
	allTransactions.Display("")

	// Example 5: Uniform treatment
	printSeparator("Example 5: Uniform Treatment")

	// Function that works with any TransactionNode
	printTransactionInfo := func(node TransactionNode) {
		fmt.Printf("Transaction/Group: %s\n", node.GetID())
		fmt.Printf("Total Amount: $%.2f\n", node.GetTotalAmount())
		fmt.Println("Structure:")
		node.Display("  ")
		fmt.Println()
	}

	fmt.Println("Treating group and individual transaction uniformly:")
	printTransactionInfo(february)                                    // Composite
	printTransactionInfo(NewTransaction("TXN012", 25.00, "Standalone", "2024-03-01")) // Leaf

	fmt.Println("\n✓ Composite pattern enables uniform treatment of transactions and groups")
	fmt.Println("✓ Operations work recursively through the tree")
	fmt.Println("✓ Easy to organize transactions by date, category, or any criteria")
	fmt.Println("✓ Client code doesn't need to distinguish between leaf and composite")
	fmt.Println("✓ JoshBank can organize transactions hierarchically for better reporting")
}
