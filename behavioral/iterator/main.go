package main

import "fmt"

// Transaction represents a banking transaction
type Transaction struct {
	ID          string
	Amount      float64
	Type        string
	Description string
}

// Iterator interface defines traversal methods
type Iterator interface {
	HasNext() bool
	Next() *Transaction
	Reset()
}

// Collection interface declares method to create iterator
type Collection interface {
	CreateIterator() Iterator
	Add(transaction *Transaction)
}

// --- Array-based Transaction History ---

type ArrayTransactionHistory struct {
	transactions []*Transaction
}

func NewArrayTransactionHistory() *ArrayTransactionHistory {
	return &ArrayTransactionHistory{transactions: make([]*Transaction, 0)}
}

func (h *ArrayTransactionHistory) Add(transaction *Transaction) {
	h.transactions = append(h.transactions, transaction)
}

func (h *ArrayTransactionHistory) CreateIterator() Iterator {
	return &ArrayIterator{history: h, index: 0}
}

type ArrayIterator struct {
	history *ArrayTransactionHistory
	index   int
}

func (i *ArrayIterator) HasNext() bool {
	return i.index < len(i.history.transactions)
}

func (i *ArrayIterator) Next() *Transaction {
	if i.HasNext() {
		transaction := i.history.transactions[i.index]
		i.index++
		return transaction
	}
	return nil
}

func (i *ArrayIterator) Reset() {
	i.index = 0
}

// --- Linked List-based Transaction History ---

type TransactionNode struct {
	transaction *Transaction
	next        *TransactionNode
}

type LinkedListTransactionHistory struct {
	head *TransactionNode
}

func NewLinkedListTransactionHistory() *LinkedListTransactionHistory {
	return &LinkedListTransactionHistory{}
}

func (h *LinkedListTransactionHistory) Add(transaction *Transaction) {
	newNode := &TransactionNode{transaction: transaction}
	if h.head == nil {
		h.head = newNode
		return
	}

	current := h.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (h *LinkedListTransactionHistory) CreateIterator() Iterator {
	return &LinkedListIterator{current: h.head, head: h.head}
}

type LinkedListIterator struct {
	current *TransactionNode
	head    *TransactionNode
}

func (i *LinkedListIterator) HasNext() bool {
	return i.current != nil
}

func (i *LinkedListIterator) Next() *Transaction {
	if i.HasNext() {
		transaction := i.current.transaction
		i.current = i.current.next
		return transaction
	}
	return nil
}

func (i *LinkedListIterator) Reset() {
	i.current = i.head
}

// Helper function to print transaction history
func printTransactionHistory(collection Collection, name string) {
	fmt.Printf("\n%s:\n", name)
	iterator := collection.CreateIterator()
	count := 1
	for iterator.HasNext() {
		txn := iterator.Next()
		fmt.Printf("  %d. %s: $%.2f - %s (%s)\n", count, txn.ID, txn.Amount, txn.Description, txn.Type)
		count++
	}
}

func main() {
	fmt.Println("=== Iterator Pattern: JoshBank Transaction History ===")

	// Create different transaction history implementations
	arrayHistory := NewArrayTransactionHistory()
	arrayHistory.Add(&Transaction{ID: "TXN001", Amount: 100.0, Type: "deposit", Description: "Salary"})
	arrayHistory.Add(&Transaction{ID: "TXN002", Amount: 50.0, Type: "withdrawal", Description: "ATM"})
	arrayHistory.Add(&Transaction{ID: "TXN003", Amount: 250.0, Type: "transfer", Description: "Bill payment"})

	linkedHistory := NewLinkedListTransactionHistory()
	linkedHistory.Add(&Transaction{ID: "TXN004", Amount: 500.0, Type: "deposit", Description: "Refund"})
	linkedHistory.Add(&Transaction{ID: "TXN005", Amount: 75.0, Type: "withdrawal", Description: "Purchase"})
	linkedHistory.Add(&Transaction{ID: "TXN006", Amount: 1000.0, Type: "transfer", Description: "Investment"})

	// Example 1: Traverse different collections uniformly
	fmt.Println("\n--- Example 1: Uniform Traversal ---")
	printTransactionHistory(arrayHistory, "Account History (Array)")
	printTransactionHistory(linkedHistory, "Account History (Linked List)")

	// Example 2: Multiple iterations
	fmt.Println("\n--- Example 2: Multiple Iterations ---")
	iter := arrayHistory.CreateIterator()
	fmt.Println("First pass:")
	for iter.HasNext() {
		txn := iter.Next()
		fmt.Printf("  - %s: $%.2f\n", txn.ID, txn.Amount)
	}

	iter.Reset()
	fmt.Println("\nSecond pass:")
	for iter.HasNext() {
		txn := iter.Next()
		fmt.Printf("  - %s: $%.2f\n", txn.ID, txn.Amount)
	}

	fmt.Println("\n✓ Iterator provides uniform way to traverse transaction collections")
	fmt.Println("✓ Hides internal structure of collections")
	fmt.Println("✓ Supports multiple simultaneous traversals")
	fmt.Println("✓ JoshBank can iterate through transactions regardless of storage implementation")
}
