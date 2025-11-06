package main

import "fmt"

// Observer interface defines the update method
type Observer interface {
	Update(transactionID string, amount float64, status string)
	GetName() string
}

// Subject interface defines methods for managing observers
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers(transactionID string, amount float64, status string)
}

// --- Concrete Subject ---

type TransactionService struct {
	observers []Observer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{observers: make([]Observer, 0)}
}

func (t *TransactionService) RegisterObserver(o Observer) {
	t.observers = append(t.observers, o)
	fmt.Printf("  [TransactionService] %s subscribed\n", o.GetName())
}

func (t *TransactionService) RemoveObserver(o Observer) {
	for i, observer := range t.observers {
		if observer == o {
			t.observers = append(t.observers[:i], t.observers[i+1:]...)
			fmt.Printf("  [TransactionService] %s unsubscribed\n", o.GetName())
			return
		}
	}
}

func (t *TransactionService) NotifyObservers(transactionID string, amount float64, status string) {
	fmt.Println("  [TransactionService] Notifying all observers...")
	for _, observer := range t.observers {
		observer.Update(transactionID, amount, status)
	}
}

func (t *TransactionService) ProcessTransaction(transactionID string, amount float64) {
	fmt.Printf("\n→ Processing transaction %s: $%.2f\n", transactionID, amount)
	// Simulate processing
	status := "completed"
	if amount > 10000 {
		status = "pending_approval"
	}
	t.NotifyObservers(transactionID, amount, status)
}

// --- Concrete Observers ---

type NotificationService struct {
	name string
}

func NewNotificationService() *NotificationService {
	return &NotificationService{name: "Notification Service"}
}

func (n *NotificationService) Update(transactionID string, amount float64, status string) {
	fmt.Printf("  [%s] Sending notification: Transaction %s - $%.2f (%s)\n", n.name, transactionID, amount, status)
}

func (n *NotificationService) GetName() string {
	return n.name
}

type AuditService struct {
	name string
}

func NewAuditService() *AuditService {
	return &AuditService{name: "Audit Service"}
}

func (a *AuditService) Update(transactionID string, amount float64, status string) {
	fmt.Printf("  [%s] Logging transaction: %s - $%.2f (%s)\n", a.name, transactionID, amount, status)
}

func (a *AuditService) GetName() string {
	return a.name
}

type ComplianceService struct {
	name string
}

func NewComplianceService() *ComplianceService {
	return &ComplianceService{name: "Compliance Service"}
}

func (c *ComplianceService) Update(transactionID string, amount float64, status string) {
	if amount > 10000 {
		fmt.Printf("  [%s] Flagging transaction %s for compliance review (amount: $%.2f)\n", c.name, transactionID, amount)
	} else {
		fmt.Printf("  [%s] Transaction %s passed compliance check\n", c.name, transactionID)
	}
}

func (c *ComplianceService) GetName() string {
	return c.name
}

type AnalyticsService struct {
	name string
}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{name: "Analytics Service"}
}

func (a *AnalyticsService) Update(transactionID string, amount float64, status string) {
	fmt.Printf("  [%s] Recording transaction metrics: %s - $%.2f\n", a.name, transactionID, amount)
}

func (a *AnalyticsService) GetName() string {
	return a.name
}

func main() {
	fmt.Println("=== Observer Pattern: JoshBank Transaction Monitoring ===")

	// Create subject
	transactionService := NewTransactionService()

	// Create observers
	notificationService := NewNotificationService()
	auditService := NewAuditService()
	complianceService := NewComplianceService()
	analyticsService := NewAnalyticsService()

	// Example 1: Register observers
	fmt.Println("\n--- Example 1: Registering Observers ---")
	transactionService.RegisterObserver(notificationService)
	transactionService.RegisterObserver(auditService)
	transactionService.RegisterObserver(complianceService)
	transactionService.RegisterObserver(analyticsService)

	// Example 2: Process transactions
	fmt.Println("\n--- Example 2: Transaction Updates ---")
	transactionService.ProcessTransaction("TXN001", 500.0)
	transactionService.ProcessTransaction("TXN002", 15000.0)
	transactionService.ProcessTransaction("TXN003", 250.0)

	// Example 3: Remove observer
	fmt.Println("\n--- Example 3: Unsubscribing Observer ---")
	transactionService.RemoveObserver(analyticsService)
	transactionService.ProcessTransaction("TXN004", 750.0)

	fmt.Println("\n✓ Observer pattern enables one-to-many dependencies")
	fmt.Println("✓ Subject and observers are loosely coupled")
	fmt.Println("✓ Observers can be added/removed dynamically")
	fmt.Println("✓ JoshBank services are automatically notified of transaction events")
}
