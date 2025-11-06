package main

import "fmt"

// --- Subsystem Components (Complex classes) ---

// AccountService manages customer accounts
type AccountService struct {
	accounts map[string]float64
}

func NewAccountService() *AccountService {
	return &AccountService{
		accounts: make(map[string]float64),
	}
}

func (a *AccountService) GetBalance(accountID string) float64 {
	return a.accounts[accountID]
}

func (a *AccountService) SetBalance(accountID string, balance float64) {
	a.accounts[accountID] = balance
	fmt.Printf("Account Service: Set balance for %s to $%.2f\n", accountID, balance)
}

func (a *AccountService) VerifyAccount(accountID string) bool {
	_, exists := a.accounts[accountID]
	if exists {
		fmt.Printf("Account Service: Account %s verified\n", accountID)
	} else {
		fmt.Printf("Account Service: Account %s not found\n", accountID)
	}
	return exists
}

// PaymentService manages payment processing
type PaymentService struct{}

func NewPaymentService() *PaymentService {
	return &PaymentService{}
}

func (p *PaymentService) ProcessPayment(fromAccount, toAccount string, amount float64) error {
	fmt.Printf("Payment Service: Processing $%.2f from %s to %s\n", amount, fromAccount, toAccount)
	return nil
}

func (p *PaymentService) ValidatePayment(amount float64) bool {
	if amount <= 0 {
		fmt.Println("Payment Service: Invalid amount")
		return false
	}
	if amount > 100000 {
		fmt.Println("Payment Service: Amount exceeds limit, requires approval")
		return false
	}
	fmt.Println("Payment Service: Payment validated")
	return true
}

// NotificationService manages notifications
type NotificationService struct{}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n *NotificationService) SendEmail(recipient, subject, message string) {
	fmt.Printf("Notification Service: Sending email to %s\n", recipient)
	fmt.Printf("  Subject: %s\n", subject)
	fmt.Printf("  Message: %s\n", message)
}

func (n *NotificationService) SendSMS(phoneNumber, message string) {
	fmt.Printf("Notification Service: Sending SMS to %s\n", phoneNumber)
	fmt.Printf("  Message: %s\n", message)
}

// ComplianceService manages compliance checks
type ComplianceService struct{}

func NewComplianceService() *ComplianceService {
	return &ComplianceService{}
}

func (c *ComplianceService) CheckAML(accountID string, amount float64) bool {
	fmt.Printf("Compliance Service: Running AML check for account %s, amount $%.2f\n", accountID, amount)
	if amount > 10000 {
		fmt.Println("Compliance Service: Flagged for review (large transaction)")
		return false
	}
	fmt.Println("Compliance Service: AML check passed")
	return true
}

func (c *ComplianceService) LogTransaction(transactionID, accountID string, amount float64) {
	fmt.Printf("Compliance Service: Logging transaction %s: $%.2f from account %s\n", transactionID, amount, accountID)
}

// AuditService manages audit trails
type AuditService struct{}

func NewAuditService() *AuditService {
	return &AuditService{}
}

func (a *AuditService) RecordTransaction(transactionID, details string) {
	fmt.Printf("Audit Service: Recording transaction %s: %s\n", transactionID, details)
}

// --- Facade ---

// BankingFacade provides a simplified interface to the banking subsystem.
// This is the facade that hides the complexity of coordinating multiple components.
type BankingFacade struct {
	accountService    *AccountService
	paymentService    *PaymentService
	notificationService *NotificationService
	complianceService *ComplianceService
	auditService      *AuditService
}

func NewBankingFacade() *BankingFacade {
	return &BankingFacade{
		accountService:     NewAccountService(),
		paymentService:     NewPaymentService(),
		notificationService: NewNotificationService(),
		complianceService:  NewComplianceService(),
		auditService:       NewAuditService(),
	}
}

// TransferMoney is a high-level method that coordinates all subsystems.
// This single method replaces 10+ individual method calls.
func (b *BankingFacade) TransferMoney(fromAccount, toAccount string, amount float64, transactionID string) error {
	fmt.Println("\n💰 Initiating money transfer...")
	fmt.Println("----------------------------------------")

	// Verify accounts
	if !b.accountService.VerifyAccount(fromAccount) {
		return fmt.Errorf("source account not found")
	}
	if !b.accountService.VerifyAccount(toAccount) {
		return fmt.Errorf("destination account not found")
	}

	// Check balance
	balance := b.accountService.GetBalance(fromAccount)
	if balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	// Validate payment
	if !b.paymentService.ValidatePayment(amount) {
		return fmt.Errorf("payment validation failed")
	}

	// Compliance check
	if !b.complianceService.CheckAML(fromAccount, amount) {
		return fmt.Errorf("compliance check failed")
	}

	// Process payment
	if err := b.paymentService.ProcessPayment(fromAccount, toAccount, amount); err != nil {
		return err
	}

	// Update balances
	b.accountService.SetBalance(fromAccount, balance-amount)
	toBalance := b.accountService.GetBalance(toAccount)
	b.accountService.SetBalance(toAccount, toBalance+amount)

	// Audit trail
	b.auditService.RecordTransaction(transactionID, fmt.Sprintf("Transfer $%.2f from %s to %s", amount, fromAccount, toAccount))
	b.complianceService.LogTransaction(transactionID, fromAccount, amount)

	// Notifications
	b.notificationService.SendEmail(fromAccount+"@email.com", "Transfer Completed", 
		fmt.Sprintf("You transferred $%.2f to %s", amount, toAccount))
	b.notificationService.SendSMS("+1234567890", fmt.Sprintf("Transfer of $%.2f completed", amount))

	fmt.Println("----------------------------------------")
	fmt.Println("✓ Transfer completed successfully")
	return nil
}

// OpenAccount is another high-level method for account creation
func (b *BankingFacade) OpenAccount(accountID string, initialBalance float64, email, phone string) error {
	fmt.Println("\n🏦 Opening new account...")
	fmt.Println("----------------------------------------")

	b.accountService.SetBalance(accountID, initialBalance)
	b.notificationService.SendEmail(email, "Welcome to JoshBank", 
		fmt.Sprintf("Your account %s has been opened with initial balance $%.2f", accountID, initialBalance))
	b.notificationService.SendSMS(phone, fmt.Sprintf("JoshBank: Account %s opened", accountID))
	b.auditService.RecordTransaction("ACCOUNT_OPEN", fmt.Sprintf("Account %s opened with $%.2f", accountID, initialBalance))

	fmt.Println("----------------------------------------")
	fmt.Println("✓ Account opened successfully")
	return nil
}

// GetAccountBalance provides simplified access to account balance
func (b *BankingFacade) GetAccountBalance(accountID string) (float64, error) {
	if !b.accountService.VerifyAccount(accountID) {
		return 0, fmt.Errorf("account not found")
	}
	return b.accountService.GetBalance(accountID), nil
}

// GetAccountService allows direct access to subsystem if needed (optional)
func (b *BankingFacade) GetAccountService() *AccountService {
	return b.accountService
}

func main() {
	fmt.Println("=== Facade Pattern: JoshBank Banking System ===")

	// Create the facade
	banking := NewBankingFacade()

	// Example 1: Open accounts using the facade
	fmt.Println("\n=== Example 1: Opening Accounts (with Facade) ===")
	banking.OpenAccount("ACC001", 5000.00, "alice@example.com", "+1234567890")
	banking.OpenAccount("ACC002", 3000.00, "bob@example.com", "+0987654321")

	// Example 2: Transfer money using the facade
	fmt.Println("\n=== Example 2: Transferring Money (with Facade) ===")
	banking.TransferMoney("ACC001", "ACC002", 500.00, "TXN001")

	// Check balances
	balance1, _ := banking.GetAccountBalance("ACC001")
	balance2, _ := banking.GetAccountBalance("ACC002")
	fmt.Printf("\nAccount ACC001 balance: $%.2f\n", balance1)
	fmt.Printf("Account ACC002 balance: $%.2f\n", balance2)

	// Example 3: Direct access to subsystem (when needed)
	fmt.Println("\n=== Example 3: Direct Subsystem Access ===")
	fmt.Println("Advanced user wants to check account directly:")
	accountService := banking.GetAccountService()
	fmt.Printf("Direct balance check: $%.2f\n", accountService.GetBalance("ACC001"))

	// Example 4: Compare with and without facade
	fmt.Println("\n=== Example 4: Without Facade (Complex) ===")
	fmt.Println("To transfer money without facade, client would need to:")
	fmt.Println("1. accountService.VerifyAccount(fromAccount)")
	fmt.Println("2. accountService.VerifyAccount(toAccount)")
	fmt.Println("3. accountService.GetBalance(fromAccount)")
	fmt.Println("4. paymentService.ValidatePayment(amount)")
	fmt.Println("5. complianceService.CheckAML(fromAccount, amount)")
	fmt.Println("6. paymentService.ProcessPayment(...)")
	fmt.Println("7. accountService.SetBalance(fromAccount, ...)")
	fmt.Println("8. accountService.SetBalance(toAccount, ...)")
	fmt.Println("9. auditService.RecordTransaction(...)")
	fmt.Println("10. complianceService.LogTransaction(...)")
	fmt.Println("11. notificationService.SendEmail(...)")
	fmt.Println("12. notificationService.SendSMS(...)")
	fmt.Println("\nWith facade: Just call banking.TransferMoney(...)!")

	fmt.Println("\n✓ Facade simplifies complex subsystem interactions")
	fmt.Println("✓ Client code is much simpler and easier to understand")
	fmt.Println("✓ Subsystems are decoupled from client code")
	fmt.Println("✓ Direct access still available when needed")
	fmt.Println("✓ JoshBank provides simple API while maintaining complex internal operations")
}
