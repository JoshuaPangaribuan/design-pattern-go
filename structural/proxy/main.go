package main

import (
	"fmt"
	"time"
)

// RiskAssessmentService is the subject interface that both RealRiskService and RiskProxy implement
type RiskAssessmentService interface {
	AssessRisk(transactionID string, amount float64) (*RiskResult, error)
	GetRiskScore(customerID string) (int, error)
}

// RiskResult represents the result of a risk assessment
type RiskResult struct {
	TransactionID string
	RiskLevel     string
	Score         int
	Recommendation string
}

// --- RealSubject ---

// RealRiskAssessmentService represents the actual risk service that is expensive to call
type RealRiskAssessmentService struct {
	serviceURL string
}

func NewRealRiskAssessmentService(serviceURL string) *RealRiskAssessmentService {
	service := &RealRiskAssessmentService{serviceURL: serviceURL}
	service.initialize()
	return service
}

// initialize simulates expensive service initialization
func (r *RealRiskAssessmentService) initialize() {
	fmt.Printf("  [RealRiskService] Initializing connection to %s...\n", r.serviceURL)
	time.Sleep(500 * time.Millisecond) // Simulate slow connection
	fmt.Printf("  [RealRiskService] Connected and ready\n")
}

// AssessRisk performs actual risk assessment (expensive operation)
func (r *RealRiskAssessmentService) AssessRisk(transactionID string, amount float64) (*RiskResult, error) {
	fmt.Printf("  [RealRiskService] Assessing risk for transaction %s (amount: $%.2f)...\n", transactionID, amount)
	time.Sleep(300 * time.Millisecond) // Simulate API call latency
	
	// Simulate risk calculation
	riskLevel := "low"
	score := 20
	if amount > 10000 {
		riskLevel = "high"
		score = 80
	} else if amount > 5000 {
		riskLevel = "medium"
		score = 50
	}
	
	return &RiskResult{
		TransactionID: transactionID,
		RiskLevel:     riskLevel,
		Score:         score,
		Recommendation: fmt.Sprintf("Transaction %s assessed as %s risk", transactionID, riskLevel),
	}, nil
}

func (r *RealRiskAssessmentService) GetRiskScore(customerID string) (int, error) {
	fmt.Printf("  [RealRiskService] Fetching risk score for customer %s...\n", customerID)
	time.Sleep(200 * time.Millisecond)
	return 35, nil
}

// --- Virtual Proxy (Lazy Loading) ---

// LazyRiskProxy is a virtual proxy that delays loading until needed
type LazyRiskProxy struct {
	serviceURL string
	realService *RealRiskAssessmentService
}

func NewLazyRiskProxy(serviceURL string) *LazyRiskProxy {
	fmt.Printf("  [LazyRiskProxy] Created proxy for risk service (not initialized yet)\n")
	return &LazyRiskProxy{serviceURL: serviceURL}
}

// AssessRisk loads the real service on first access (lazy loading)
func (p *LazyRiskProxy) AssessRisk(transactionID string, amount float64) (*RiskResult, error) {
	if p.realService == nil {
		fmt.Printf("  [LazyRiskProxy] First access - initializing real service...\n")
		p.realService = NewRealRiskAssessmentService(p.serviceURL)
	}
	return p.realService.AssessRisk(transactionID, amount)
}

func (p *LazyRiskProxy) GetRiskScore(customerID string) (int, error) {
	if p.realService == nil {
		fmt.Printf("  [LazyRiskProxy] First access - initializing real service...\n")
		p.realService = NewRealRiskAssessmentService(p.serviceURL)
	}
	return p.realService.GetRiskScore(customerID)
}

// --- Protection Proxy (Access Control) ---

// User represents a user with permissions
type User struct {
	name        string
	role        string
	canAssessRisk bool
}

// ComplianceService is the subject interface for compliance operations
type ComplianceService interface {
	CheckCompliance(transactionID string) (bool, error)
	GenerateReport(period string) (string, error)
}

// RealComplianceService is the actual compliance service
type RealComplianceService struct{}

func (c *RealComplianceService) CheckCompliance(transactionID string) (bool, error) {
	fmt.Printf("  [RealCompliance] Checking compliance for transaction %s\n", transactionID)
	return true, nil
}

func (c *RealComplianceService) GenerateReport(period string) (string, error) {
	fmt.Printf("  [RealCompliance] Generating compliance report for %s\n", period)
	return fmt.Sprintf("Compliance report for %s", period), nil
}

// ComplianceProxy is a protection proxy that controls access
type ComplianceProxy struct {
	service *RealComplianceService
	user    *User
}

func NewComplianceProxy(user *User) *ComplianceProxy {
	return &ComplianceProxy{
		service: &RealComplianceService{},
		user:    user,
	}
}

func (p *ComplianceProxy) CheckCompliance(transactionID string) (bool, error) {
	fmt.Printf("  [ComplianceProxy] User '%s' checking compliance\n", p.user.name)
	return p.service.CheckCompliance(transactionID)
}

func (p *ComplianceProxy) GenerateReport(period string) (string, error) {
	if !p.user.canAssessRisk {
		fmt.Printf("  [ComplianceProxy] Access denied: User '%s' cannot generate reports\n", p.user.name)
		return "", fmt.Errorf("access denied: user does not have permission to generate reports")
	}
	fmt.Printf("  [ComplianceProxy] User '%s' generating report\n", p.user.name)
	return p.service.GenerateReport(period)
}

// --- Caching Proxy ---

// ExchangeRateService is the subject interface for exchange rate operations
type ExchangeRateService interface {
	GetRate(fromCurrency, toCurrency string) (float64, error)
}

// RealExchangeRateService represents an expensive remote service
type RealExchangeRateService struct{}

func (s *RealExchangeRateService) GetRate(fromCurrency, toCurrency string) (float64, error) {
	fmt.Printf("  [RealExchangeRate] Fetching rate %s/%s from remote API...\n", fromCurrency, toCurrency)
	time.Sleep(300 * time.Millisecond) // Simulate network latency
	// Simulate exchange rate
	rate := 1.0
	if fromCurrency == "USD" && toCurrency == "EUR" {
		rate = 0.85
	} else if fromCurrency == "USD" && toCurrency == "GBP" {
		rate = 0.75
	}
	return rate, nil
}

// CachingProxy caches results to avoid repeated expensive calls
type CachingProxy struct {
	service *RealExchangeRateService
	cache    map[string]float64
}

func NewCachingProxy() *CachingProxy {
	return &CachingProxy{
		service: &RealExchangeRateService{},
		cache:   make(map[string]float64),
	}
}

func (p *CachingProxy) GetRate(fromCurrency, toCurrency string) (float64, error) {
	key := fmt.Sprintf("%s-%s", fromCurrency, toCurrency)
	// Check cache first
	if rate, exists := p.cache[key]; exists {
		fmt.Printf("  [CachingProxy] Cache hit for %s/%s\n", fromCurrency, toCurrency)
		return rate, nil
	}

	// Cache miss - fetch from real service
	fmt.Printf("  [CachingProxy] Cache miss for %s/%s\n", fromCurrency, toCurrency)
	rate, err := p.service.GetRate(fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}
	p.cache[key] = rate
	return rate, nil
}

// --- Logging Proxy ---

// LoggingProxy adds logging to any RiskAssessmentService
type LoggingProxy struct {
	service     RiskAssessmentService
	accessCount int
}

func NewLoggingProxy(service RiskAssessmentService) *LoggingProxy {
	return &LoggingProxy{service: service}
}

func (p *LoggingProxy) AssessRisk(transactionID string, amount float64) (*RiskResult, error) {
	p.accessCount++
	fmt.Printf("  [LoggingProxy] Risk assessment #%d at %s\n", p.accessCount, time.Now().Format("15:04:05"))
	return p.service.AssessRisk(transactionID, amount)
}

func (p *LoggingProxy) GetRiskScore(customerID string) (int, error) {
	p.accessCount++
	fmt.Printf("  [LoggingProxy] Risk score lookup #%d at %s\n", p.accessCount, time.Now().Format("15:04:05"))
	return p.service.GetRiskScore(customerID)
}

func main() {
	fmt.Println("=== Proxy Pattern: JoshBank Risk & Compliance Services ===")

	// Example 1: Virtual Proxy (Lazy Loading)
	fmt.Println("\n--- Example 1: Virtual Proxy (Lazy Loading) ---")
	fmt.Println("Creating risk service proxies (services not initialized yet):")

	riskProxy1 := NewLazyRiskProxy("https://risk-api.joshbank.com")
	riskProxy2 := NewLazyRiskProxy("https://risk-api.joshbank.com")

	fmt.Println("\nAssessing risk for first transaction:")
	result1, _ := riskProxy1.AssessRisk("TXN001", 5000.0)
	fmt.Printf("Result: %s (Risk Level: %s, Score: %d)\n", result1.Recommendation, result1.RiskLevel, result1.Score)

	fmt.Println("\nOther proxy still not initialized:")
	score, _ := riskProxy2.GetRiskScore("CUST001")
	fmt.Printf("Risk Score: %d\n", score)

	fmt.Println("\nAssessing risk again (already initialized):")
	riskProxy1.AssessRisk("TXN002", 15000.0)

	// Example 2: Protection Proxy (Access Control)
	fmt.Println("\n--- Example 2: Protection Proxy (Access Control) ---")

	adminUser := &User{name: "Alice", role: "admin", canAssessRisk: true}
	analystUser := &User{name: "Bob", role: "analyst", canAssessRisk: false}

	adminCompliance := NewComplianceProxy(adminUser)
	analystCompliance := NewComplianceProxy(analystUser)

	fmt.Println("\nAdmin user operations:")
	adminCompliance.CheckCompliance("TXN001")
	adminCompliance.GenerateReport("Q1 2024")

	fmt.Println("\nAnalyst user operations:")
	analystCompliance.CheckCompliance("TXN002")
	_, err := analystCompliance.GenerateReport("Q1 2024")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Example 3: Caching Proxy
	fmt.Println("\n--- Example 3: Caching Proxy ---")

	rateService := NewCachingProxy()

	fmt.Println("\nFirst access (cache miss):")
	rate1, _ := rateService.GetRate("USD", "EUR")
	fmt.Printf("Rate: %.2f\n", rate1)

	fmt.Println("\nSecond access (cache hit):")
	rate2, _ := rateService.GetRate("USD", "EUR")
	fmt.Printf("Rate: %.2f\n", rate2)

	fmt.Println("\nDifferent currency (cache miss):")
	rate3, _ := rateService.GetRate("USD", "GBP")
	fmt.Printf("Rate: %.2f\n", rate3)

	// Example 4: Logging Proxy
	fmt.Println("\n--- Example 4: Logging Proxy ---")

	realRiskService := NewRealRiskAssessmentService("https://risk-api.joshbank.com")
	loggedRiskService := NewLoggingProxy(realRiskService)

	fmt.Println("\nAccessing risk service multiple times:")
	loggedRiskService.AssessRisk("TXN003", 3000.0)
	loggedRiskService.GetRiskScore("CUST002")
	loggedRiskService.AssessRisk("TXN004", 8000.0)

	// Example 5: Combining Proxies
	fmt.Println("\n--- Example 5: Combining Multiple Proxies ---")

	// Create a proxy chain: Logging -> Virtual Proxy
	virtualProxy := NewLazyRiskProxy("https://risk-api.joshbank.com")
	combinedProxy := NewLoggingProxy(virtualProxy)

	fmt.Println("\nFirst access (lazy load + logging):")
	combinedProxy.AssessRisk("TXN005", 2000.0)

	fmt.Println("\nSecond access (already loaded + logging):")
	combinedProxy.AssessRisk("TXN006", 4000.0)

	fmt.Println("\n✓ Proxy pattern provides controlled access to expensive services")
	fmt.Println("✓ Virtual proxy enables lazy loading")
	fmt.Println("✓ Protection proxy enforces access control")
	fmt.Println("✓ Caching proxy improves performance")
	fmt.Println("✓ Logging proxy adds monitoring")
	fmt.Println("✓ Proxies can be combined for multiple concerns")
	fmt.Println("✓ JoshBank can optimize expensive operations without changing core code")
}
