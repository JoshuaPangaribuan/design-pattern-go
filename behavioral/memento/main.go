package main

import (
	"fmt"
	"time"
)

// AccountMemento stores the state of the Account
type AccountMemento struct {
	balance   float64
	timestamp time.Time
}

func (m *AccountMemento) GetTimestamp() time.Time {
	return m.timestamp
}

// Account is the originator that creates mementos
type Account struct {
	accountID string
	balance   float64
}

func NewAccount(accountID string, initialBalance float64) *Account {
	return &Account{
		accountID: accountID,
		balance:   initialBalance,
	}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited $%.2f | Balance: $%.2f\n", amount, a.balance)
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	fmt.Printf("Withdrew $%.2f | Balance: $%.2f\n", amount, a.balance)
	return nil
}

func (a *Account) Save() *AccountMemento {
	fmt.Println("  [Saving account state...]")
	return &AccountMemento{
		balance:   a.balance,
		timestamp: time.Now(),
	}
}

func (a *Account) Restore(m *AccountMemento) {
	a.balance = m.balance
	fmt.Printf("  [Restored balance to: $%.2f]\n", a.balance)
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

// TransactionHistory is the caretaker that manages mementos
type TransactionHistory struct {
	mementos []*AccountMemento
	current  int
}

func NewTransactionHistory() *TransactionHistory {
	return &TransactionHistory{
		mementos: make([]*AccountMemento, 0),
		current:  -1,
	}
}

func (h *TransactionHistory) Save(m *AccountMemento) {
	// Remove any mementos after current position (for redo)
	h.mementos = h.mementos[:h.current+1]
	h.mementos = append(h.mementos, m)
	h.current++
}

func (h *TransactionHistory) Undo() *AccountMemento {
	if h.current > 0 {
		h.current--
		return h.mementos[h.current]
	}
	return nil
}

func (h *TransactionHistory) Redo() *AccountMemento {
	if h.current < len(h.mementos)-1 {
		h.current++
		return h.mementos[h.current]
	}
	return nil
}

func (h *TransactionHistory) ShowHistory() {
	fmt.Println("\nTransaction History:")
	for i, m := range h.mementos {
		marker := " "
		if i == h.current {
			marker = "→"
		}
		fmt.Printf("  %s %d. Balance: $%.2f (%s)\n", marker, i+1, m.balance,
			m.timestamp.Format("15:04:05"))
	}
}

func main() {
	fmt.Println("=== Memento Pattern: JoshBank Account State Management ===")

	account := NewAccount("ACC001", 1000.0)
	history := NewTransactionHistory()

	// Example 1: Transactions and save
	fmt.Println("\n--- Example 1: Transactions and Saving ---")
	account.Deposit(500.0)
	history.Save(account.Save())

	account.Withdraw(200.0)
	history.Save(account.Save())

	account.Deposit(1000.0)
	history.Save(account.Save())

	// Example 2: Undo
	fmt.Println("\n--- Example 2: Undo Operations ---")
	if m := history.Undo(); m != nil {
		account.Restore(m)
	}

	if m := history.Undo(); m != nil {
		account.Restore(m)
	}

	// Example 3: Redo
	fmt.Println("\n--- Example 3: Redo Operations ---")
	if m := history.Redo(); m != nil {
		account.Restore(m)
	}

	// Example 4: View history
	history.ShowHistory()

	fmt.Println("\n✓ Memento captures and restores account state")
	fmt.Println("✓ Enables undo/redo functionality for transactions")
	fmt.Println("✓ Preserves encapsulation")
	fmt.Println("✓ JoshBank can implement transaction rollback and audit trails")
}
