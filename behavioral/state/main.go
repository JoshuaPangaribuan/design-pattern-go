package main

import "fmt"

// AccountState interface defines state-specific behavior
type AccountState interface {
	Deposit(amount float64) string
	Withdraw(amount float64) string
	Close() string
	GetStateName() string
}

// Account is the context that maintains current state
type Account struct {
	activeState    AccountState
	frozenState    AccountState
	closedState    AccountState

	currentState AccountState
	accountID    string
	balance      float64
}

func NewAccount(accountID string, initialBalance float64) *Account {
	account := &Account{
		accountID: accountID,
		balance:  initialBalance,
	}

	account.activeState = &ActiveState{account: account}
	account.frozenState = &FrozenState{account: account}
	account.closedState = &ClosedState{account: account}

	account.currentState = account.activeState

	return account
}

func (a *Account) Deposit(amount float64) {
	fmt.Println(a.currentState.Deposit(amount))
}

func (a *Account) Withdraw(amount float64) {
	fmt.Println(a.currentState.Withdraw(amount))
}

func (a *Account) Close() {
	fmt.Println(a.currentState.Close())
}

func (a *Account) SetState(state AccountState) {
	a.currentState = state
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) SetBalance(balance float64) {
	a.balance = balance
}

// --- Concrete States ---

type ActiveState struct {
	account *Account
}

func (s *ActiveState) Deposit(amount float64) string {
	s.account.balance += amount
	return fmt.Sprintf("Deposited $%.2f. New balance: $%.2f", amount, s.account.balance)
}

func (s *ActiveState) Withdraw(amount float64) string {
	if s.account.balance >= amount {
		s.account.balance -= amount
		return fmt.Sprintf("Withdrew $%.2f. New balance: $%.2f", amount, s.account.balance)
	}
	return "Insufficient funds"
}

func (s *ActiveState) Close() string {
	s.account.SetState(s.account.closedState)
	return "Account closed"
}

func (s *ActiveState) GetStateName() string {
	return "Active"
}

type FrozenState struct {
	account *Account
}

func (s *FrozenState) Deposit(amount float64) string {
	s.account.balance += amount
	// Auto-unfreeze if balance becomes positive
	if s.account.balance > 0 {
		s.account.SetState(s.account.activeState)
		return fmt.Sprintf("Deposited $%.2f. Account unfrozen. New balance: $%.2f", amount, s.account.balance)
	}
	return fmt.Sprintf("Deposited $%.2f. Account still frozen. Balance: $%.2f", amount, s.account.balance)
}

func (s *FrozenState) Withdraw(amount float64) string {
	return "Account is frozen. Cannot withdraw"
}

func (s *FrozenState) Close() string {
	s.account.SetState(s.account.closedState)
	return "Account closed"
}

func (s *FrozenState) GetStateName() string {
	return "Frozen"
}

type ClosedState struct {
	account *Account
}

func (s *ClosedState) Deposit(amount float64) string {
	return "Account is closed. Cannot deposit"
}

func (s *ClosedState) Withdraw(amount float64) string {
	return "Account is closed. Cannot withdraw"
}

func (s *ClosedState) Close() string {
	return "Account is already closed"
}

func (s *ClosedState) GetStateName() string {
	return "Closed"
}

func main() {
	fmt.Println("=== State Pattern: JoshBank Account States ===")

	// Create account
	account := NewAccount("ACC001", 1000.0)

	// Example 1: Normal operations (Active state)
	fmt.Println("\n--- Example 1: Active Account Operations ---")
	account.Deposit(500.0)
	account.Withdraw(200.0)
	account.Withdraw(1500.0) // Insufficient funds

	// Example 2: Freeze account (simulate negative balance)
	fmt.Println("\n--- Example 2: Account Frozen ---")
	account.SetBalance(-100.0)
	account.SetState(account.frozenState)
	account.Withdraw(50.0) // Cannot withdraw when frozen
	account.Deposit(150.0)  // Auto-unfreeze

	// Example 3: Close account
	fmt.Println("\n--- Example 3: Close Account ---")
	account.Close()
	account.Deposit(100.0) // Cannot deposit when closed
	account.Withdraw(50.0)  // Cannot withdraw when closed

	fmt.Println("\n✓ State pattern encapsulates state-specific behavior")
	fmt.Println("✓ Eliminates complex conditionals")
	fmt.Println("✓ Easy to add new account states")
	fmt.Println("✓ JoshBank accounts behave differently based on their state")
}
