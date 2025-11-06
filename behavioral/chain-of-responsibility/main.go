package main

import "fmt"

// TransactionRequest represents a transaction that needs approval
type TransactionRequest struct {
	ID          string
	CustomerID  string
	Amount      float64
	Type        string // "transfer", "withdrawal", "deposit"
	Description string
	Priority    string // "low", "medium", "high", "critical"
}

// Handler is the interface that all handlers in the chain must implement
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request *TransactionRequest) bool
}

// BaseHandler provides common functionality for all handlers
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) HandleNext(request *TransactionRequest) bool {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return false
}

// --- Concrete Handlers ---

// LowAmountHandler handles low-priority transactions
type LowAmountHandler struct {
	BaseHandler
	name string
}

func NewLowAmountHandler(name string) *LowAmountHandler {
	return &LowAmountHandler{name: name}
}

func (h *LowAmountHandler) Handle(request *TransactionRequest) bool {
	if request.Amount <= 1000.0 {
		fmt.Printf("[%s] Handling transaction %s: $%.2f - %s\n", h.name, request.ID, request.Amount, request.Description)
		fmt.Printf("  → Approved: Low amount transaction\n")
		return true
	}

	fmt.Printf("[%s] Cannot handle transaction %s (amount: $%.2f), escalating...\n",
		h.name, request.ID, request.Amount)
	return h.HandleNext(request)
}

// MediumAmountHandler handles medium-priority transactions
type MediumAmountHandler struct {
	BaseHandler
	name string
}

func NewMediumAmountHandler(name string) *MediumAmountHandler {
	return &MediumAmountHandler{name: name}
}

func (h *MediumAmountHandler) Handle(request *TransactionRequest) bool {
	if request.Amount > 1000.0 && request.Amount <= 10000.0 {
		fmt.Printf("[%s] Handling transaction %s: $%.2f - %s\n", h.name, request.ID, request.Amount, request.Description)
		fmt.Printf("  → Approved: Medium amount transaction\n")
		return true
	}

	fmt.Printf("[%s] Cannot handle transaction %s (amount: $%.2f), escalating...\n",
		h.name, request.ID, request.Amount)
	return h.HandleNext(request)
}

// ManagerHandler handles high-priority transactions
type ManagerHandler struct {
	BaseHandler
	name string
}

func NewManagerHandler(name string) *ManagerHandler {
	return &ManagerHandler{name: name}
}

func (h *ManagerHandler) Handle(request *TransactionRequest) bool {
	if request.Amount > 10000.0 && request.Amount <= 50000.0 {
		fmt.Printf("[%s] Handling transaction %s: $%.2f - %s\n", h.name, request.ID, request.Amount, request.Description)
		fmt.Printf("  → Approved: Manager approval required\n")
		return true
	}

	fmt.Printf("[%s] Cannot handle transaction %s (amount: $%.2f), escalating...\n",
		h.name, request.ID, request.Amount)
	return h.HandleNext(request)
}

// DirectorHandler handles critical transactions
type DirectorHandler struct {
	BaseHandler
	name string
}

func NewDirectorHandler(name string) *DirectorHandler {
	return &DirectorHandler{name: name}
}

func (h *DirectorHandler) Handle(request *TransactionRequest) bool {
	if request.Amount > 50000.0 {
		fmt.Printf("[%s] Handling transaction %s: $%.2f - %s\n", h.name, request.ID, request.Amount, request.Description)
		fmt.Printf("  → Approved: Director approval required\n")
		return true
	}

	fmt.Printf("[%s] No one can handle transaction %s\n", h.name, request.ID)
	return false
}

func main() {
	fmt.Println("=== Chain of Responsibility Pattern: JoshBank Transaction Approval ===")

	// Build the chain
	lowAmount := NewLowAmountHandler("Auto-Approval System")
	mediumAmount := NewMediumAmountHandler("Supervisor")
	manager := NewManagerHandler("Manager")
	director := NewDirectorHandler("Director")

	lowAmount.SetNext(mediumAmount).SetNext(manager).SetNext(director)

	// Create transactions with different amounts
	transactions := []*TransactionRequest{
		{ID: "TXN001", CustomerID: "CUST001", Amount: 500.0, Type: "transfer", Description: "Payment to merchant", Priority: "low"},
		{ID: "TXN002", CustomerID: "CUST002", Amount: 5000.0, Type: "transfer", Description: "Bill payment", Priority: "medium"},
		{ID: "TXN003", CustomerID: "CUST003", Amount: 25000.0, Type: "withdrawal", Description: "Large withdrawal", Priority: "high"},
		{ID: "TXN004", CustomerID: "CUST004", Amount: 100000.0, Type: "transfer", Description: "Business transfer", Priority: "critical"},
	}

	// Process each transaction through the chain
	for _, txn := range transactions {
		fmt.Printf("\n→ Processing transaction %s ($%.2f)\n", txn.ID, txn.Amount)
		handled := lowAmount.Handle(txn)
		if !handled {
			fmt.Printf("  ✗ Transaction %s was not handled\n", txn.ID)
		}
	}

	fmt.Println("\n✓ Chain of Responsibility decouples sender from receiver")
	fmt.Println("✓ Each handler decides to process or pass to next")
	fmt.Println("✓ Chain can be modified dynamically")
	fmt.Println("✓ Useful for transaction approval workflows at JoshBank")
}
