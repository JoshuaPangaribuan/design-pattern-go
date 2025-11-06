package main

import "fmt"

// BankingMediator interface defines communication methods
type BankingMediator interface {
	Notify(sender Component, event string, data interface{})
}

// Component is the base for all colleagues
type Component interface {
	SetMediator(mediator BankingMediator)
}

// --- Concrete Mediator ---

type TransactionCoordinator struct {
	paymentService    *PaymentService
	notificationService *NotificationService
	auditService     *AuditService
	complianceService *ComplianceService
}

func NewTransactionCoordinator() *TransactionCoordinator {
	coordinator := &TransactionCoordinator{
		paymentService:     &PaymentService{},
		notificationService: &NotificationService{},
		auditService:        &AuditService{},
		complianceService:   &ComplianceService{},
	}

	coordinator.paymentService.SetMediator(coordinator)
	coordinator.notificationService.SetMediator(coordinator)
	coordinator.auditService.SetMediator(coordinator)
	coordinator.complianceService.SetMediator(coordinator)

	return coordinator
}

func (t *TransactionCoordinator) Notify(sender Component, event string, data interface{}) {
	switch event {
	case "payment_processed":
		paymentData := data.(map[string]interface{})
		fmt.Printf("[Coordinator] Payment processed: %s - $%.2f\n", 
			paymentData["transactionID"], paymentData["amount"])
		t.auditService.LogTransaction(paymentData["transactionID"].(string), paymentData["amount"].(float64))
		t.notificationService.SendNotification(paymentData["customerID"].(string), 
			fmt.Sprintf("Payment of $%.2f processed", paymentData["amount"].(float64)))
		t.complianceService.CheckTransaction(paymentData["transactionID"].(string), paymentData["amount"].(float64))

	case "compliance_flag":
		flagData := data.(map[string]interface{})
		fmt.Printf("[Coordinator] Compliance flag raised: %s\n", flagData["transactionID"])
		t.auditService.LogComplianceFlag(flagData["transactionID"].(string))
		t.notificationService.SendAlert("compliance@joshbank.com", 
			fmt.Sprintf("Compliance review needed for transaction %s", flagData["transactionID"]))
	}
}

// --- Colleagues ---

type PaymentService struct {
	mediator BankingMediator
}

func (p *PaymentService) SetMediator(mediator BankingMediator) {
	p.mediator = mediator
}

func (p *PaymentService) ProcessPayment(transactionID, customerID string, amount float64) {
	fmt.Printf("[PaymentService] Processing payment: %s - $%.2f\n", transactionID, amount)
	p.mediator.Notify(p, "payment_processed", map[string]interface{}{
		"transactionID": transactionID,
		"customerID":    customerID,
		"amount":        amount,
	})
}

type NotificationService struct {
	mediator BankingMediator
}

func (n *NotificationService) SetMediator(mediator BankingMediator) {
	n.mediator = mediator
}

func (n *NotificationService) SendNotification(recipient, message string) {
	fmt.Printf("[NotificationService] Sending to %s: %s\n", recipient, message)
}

func (n *NotificationService) SendAlert(recipient, message string) {
	fmt.Printf("[NotificationService] ALERT to %s: %s\n", recipient, message)
}

type AuditService struct {
	mediator BankingMediator
}

func (a *AuditService) SetMediator(mediator BankingMediator) {
	a.mediator = mediator
}

func (a *AuditService) LogTransaction(transactionID string, amount float64) {
	fmt.Printf("[AuditService] Logging transaction: %s - $%.2f\n", transactionID, amount)
}

func (a *AuditService) LogComplianceFlag(transactionID string) {
	fmt.Printf("[AuditService] Logging compliance flag: %s\n", transactionID)
}

type ComplianceService struct {
	mediator BankingMediator
}

func (c *ComplianceService) SetMediator(mediator BankingMediator) {
	c.mediator = mediator
}

func (c *ComplianceService) CheckTransaction(transactionID string, amount float64) {
	if amount > 10000 {
		fmt.Printf("[ComplianceService] Flagging transaction %s for review\n", transactionID)
		c.mediator.Notify(c, "compliance_flag", map[string]interface{}{
			"transactionID": transactionID,
		})
	} else {
		fmt.Printf("[ComplianceService] Transaction %s passed compliance check\n", transactionID)
	}
}

func main() {
	fmt.Println("=== Mediator Pattern: JoshBank Transaction Coordination ===")

	coordinator := NewTransactionCoordinator()

	fmt.Println("\n--- Processing Transactions ---")
	coordinator.paymentService.ProcessPayment("TXN001", "CUST001", 500.0)
	fmt.Println()
	coordinator.paymentService.ProcessPayment("TXN002", "CUST002", 15000.0)

	fmt.Println("\n✓ Mediator centralizes complex communications")
	fmt.Println("✓ Reduces coupling between banking services")
	fmt.Println("✓ Easy to understand and maintain interactions")
	fmt.Println("✓ JoshBank services coordinate through mediator instead of direct communication")
}
