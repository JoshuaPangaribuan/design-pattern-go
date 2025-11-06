package main

import "fmt"

// BankingCommand interface declares methods for executing and undoing banking operations
type BankingCommand interface {
	Execute()
	Undo()
	GetDescription() string
}

// --- Receivers (banking services that perform actual work) ---

type Account struct {
	accountID string
	balance   float64
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("  [Account %s] Deposited $%.2f, Balance: $%.2f\n", a.accountID, amount, a.balance)
}

func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	fmt.Printf("  [Account %s] Withdrew $%.2f, Balance: $%.2f\n", a.accountID, amount, a.balance)
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

type TransferService struct{}

func (t *TransferService) Transfer(from, to *Account, amount float64) error {
	if from.balance < amount {
		return fmt.Errorf("insufficient funds in source account")
	}
	from.balance -= amount
	to.balance += amount
	fmt.Printf("  [Transfer] Transferred $%.2f from %s to %s\n", amount, from.accountID, to.accountID)
	return nil
}

// --- Concrete Commands ---

type DepositCommand struct {
	account *Account
	amount  float64
}

func (c *DepositCommand) Execute() {
	c.account.Deposit(c.amount)
}

func (c *DepositCommand) Undo() {
	c.account.Withdraw(c.amount)
}

func (c *DepositCommand) GetDescription() string {
	return fmt.Sprintf("Deposit $%.2f to account %s", c.amount, c.account.accountID)
}

type WithdrawCommand struct {
	account *Account
	amount  float64
	success bool
}

func (c *WithdrawCommand) Execute() {
	err := c.account.Withdraw(c.amount)
	c.success = (err == nil)
}

func (c *WithdrawCommand) Undo() {
	if c.success {
		c.account.Deposit(c.amount)
	}
}

func (c *WithdrawCommand) GetDescription() string {
	return fmt.Sprintf("Withdraw $%.2f from account %s", c.amount, c.account.accountID)
}

type TransferCommand struct {
	transferService *TransferService
	from            *Account
	to              *Account
	amount          float64
	success         bool
}

func (c *TransferCommand) Execute() {
	err := c.transferService.Transfer(c.from, c.to, c.amount)
	c.success = (err == nil)
}

func (c *TransferCommand) Undo() {
	if c.success {
		c.transferService.Transfer(c.to, c.from, c.amount)
	}
}

func (c *TransferCommand) GetDescription() string {
	return fmt.Sprintf("Transfer $%.2f from %s to %s", c.amount, c.from.accountID, c.to.accountID)
}

// MacroCommand executes multiple commands
type MacroCommand struct {
	commands    []BankingCommand
	description string
}

func NewMacroCommand(description string, commands []BankingCommand) *MacroCommand {
	return &MacroCommand{commands: commands, description: description}
}

func (m *MacroCommand) Execute() {
	for _, cmd := range m.commands {
		cmd.Execute()
	}
}

func (m *MacroCommand) Undo() {
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
}

func (m *MacroCommand) GetDescription() string {
	return m.description
}

// --- Invoker ---

type BankingController struct {
	history []BankingCommand
}

func NewBankingController() *BankingController {
	return &BankingController{history: make([]BankingCommand, 0)}
}

func (b *BankingController) ExecuteCommand(cmd BankingCommand) {
	fmt.Printf("→ Executing: %s\n", cmd.GetDescription())
	cmd.Execute()
	b.history = append(b.history, cmd)
}

func (b *BankingController) UndoLast() {
	if len(b.history) == 0 {
		fmt.Println("  Nothing to undo")
		return
	}

	cmd := b.history[len(b.history)-1]
	b.history = b.history[:len(b.history)-1]

	fmt.Printf("→ Undoing: %s\n", cmd.GetDescription())
	cmd.Undo()
}

func main() {
	fmt.Println("=== Command Pattern: JoshBank Transaction Controller ===")

	// Create accounts
	account1 := &Account{accountID: "ACC001", balance: 1000.0}
	account2 := &Account{accountID: "ACC002", balance: 500.0}
	transferService := &TransferService{}

	// Create commands
	deposit1 := &DepositCommand{account: account1, amount: 200.0}
	withdraw1 := &WithdrawCommand{account: account1, amount: 150.0}
	transfer1 := &TransferCommand{transferService: transferService, from: account1, to: account2, amount: 100.0}

	// Create invoker
	controller := NewBankingController()

	// Example 1: Execute individual commands
	fmt.Println("\n--- Example 1: Individual Commands ---")
	controller.ExecuteCommand(deposit1)
	controller.ExecuteCommand(withdraw1)
	controller.ExecuteCommand(transfer1)

	// Example 2: Undo commands
	fmt.Println("\n--- Example 2: Undo Operations ---")
	controller.UndoLast()
	controller.UndoLast()

	// Example 3: Macro command (Bill Payment)
	fmt.Println("\n--- Example 3: Macro Command (Bill Payment) ---")
	billPayment := NewMacroCommand("Bill Payment", []BankingCommand{
		&WithdrawCommand{account: account1, amount: 50.0},
		&WithdrawCommand{account: account1, amount: 75.0},
		&WithdrawCommand{account: account1, amount: 25.0},
	})

	controller.ExecuteCommand(billPayment)

	fmt.Println("\n--- Undo Bill Payment ---")
	controller.UndoLast()

	fmt.Println("\n✓ Command pattern encapsulates banking operations as objects")
	fmt.Println("✓ Supports undo/redo operations")
	fmt.Println("✓ Commands can be queued and logged for audit")
	fmt.Println("✓ Macro commands combine multiple operations")
	fmt.Println("✓ JoshBank can implement transaction rollback and audit trails")
}
