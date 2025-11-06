package main

import "fmt"

// KYCVerification defines the template method and abstract operations
type KYCVerification interface {
	CollectDocuments()
	VerifyIdentity()
	CheckCompliance()
	ApproveAccount()
	RejectAccount()
}

// BaseKYCVerification provides the template method
type BaseKYCVerification struct {
	verification KYCVerification
}

// Verify is the template method that defines the algorithm structure
func (b *BaseKYCVerification) Verify() {
	fmt.Println("=== Starting KYC Verification Process ===")
	b.verification.CollectDocuments()
	b.verification.VerifyIdentity()
	b.verification.CheckCompliance()
	
	// Decision point - can be overridden
	if b.shouldApprove() {
		b.verification.ApproveAccount()
	} else {
		b.verification.RejectAccount()
	}
	fmt.Println("=== KYC Verification Complete ===\n")
}

func (b *BaseKYCVerification) shouldApprove() bool {
	// Default logic - can be overridden
	return true
}

// --- Concrete Implementations ---

type PersonalAccountKYC struct {
	BaseKYCVerification
	customerName string
}

func NewPersonalAccountKYC(customerName string) *PersonalAccountKYC {
	kyc := &PersonalAccountKYC{customerName: customerName}
	kyc.BaseKYCVerification.verification = kyc
	return kyc
}

func (p *PersonalAccountKYC) CollectDocuments() {
	fmt.Printf("  [Personal KYC] Collecting ID and proof of address for %s\n", p.customerName)
}

func (p *PersonalAccountKYC) VerifyIdentity() {
	fmt.Println("  [Personal KYC] Verifying identity documents")
}

func (p *PersonalAccountKYC) CheckCompliance() {
	fmt.Println("  [Personal KYC] Running basic compliance checks")
}

func (p *PersonalAccountKYC) ApproveAccount() {
	fmt.Printf("  [Personal KYC] Account approved for %s\n", p.customerName)
}

func (p *PersonalAccountKYC) RejectAccount() {
	fmt.Printf("  [Personal KYC] Account rejected for %s\n", p.customerName)
}

type BusinessAccountKYC struct {
	BaseKYCVerification
	businessName string
}

func NewBusinessAccountKYC(businessName string) *BusinessAccountKYC {
	kyc := &BusinessAccountKYC{businessName: businessName}
	kyc.BaseKYCVerification.verification = kyc
	return kyc
}

func (b *BusinessAccountKYC) CollectDocuments() {
	fmt.Printf("  [Business KYC] Collecting business license, tax ID, and ownership documents for %s\n", b.businessName)
}

func (b *BusinessAccountKYC) VerifyIdentity() {
	fmt.Println("  [Business KYC] Verifying business registration and authorized signatories")
}

func (b *BusinessAccountKYC) CheckCompliance() {
	fmt.Println("  [Business KYC] Running enhanced compliance checks (AML, PEP screening)")
}

func (b *BusinessAccountKYC) ApproveAccount() {
	fmt.Printf("  [Business KYC] Business account approved for %s\n", b.businessName)
}

func (b *BusinessAccountKYC) RejectAccount() {
	fmt.Printf("  [Business KYC] Business account rejected for %s\n", b.businessName)
}

// --- Another Example: Loan Approval ---

type LoanApproval interface {
	CheckCreditScore()
	VerifyIncome()
	AssessCollateral()
	ApproveLoan()
	RejectLoan()
}

type BaseLoanApproval struct {
	approval LoanApproval
}

func (b *BaseLoanApproval) Process() {
	fmt.Println("\n--- Loan Approval Process ---")
	b.approval.CheckCreditScore()
	b.approval.VerifyIncome()
	b.approval.AssessCollateral()
	
	if b.shouldApprove() {
		b.approval.ApproveLoan()
	} else {
		b.approval.RejectLoan()
	}
}

func (b *BaseLoanApproval) shouldApprove() bool {
	return true
}

type PersonalLoanApproval struct {
	BaseLoanApproval
	applicantName string
}

func NewPersonalLoanApproval(applicantName string) *PersonalLoanApproval {
	approval := &PersonalLoanApproval{applicantName: applicantName}
	approval.BaseLoanApproval.approval = approval
	return approval
}

func (p *PersonalLoanApproval) CheckCreditScore() {
	fmt.Println("  [Personal Loan] Checking credit score")
}

func (p *PersonalLoanApproval) VerifyIncome() {
	fmt.Println("  [Personal Loan] Verifying employment and income")
}

func (p *PersonalLoanApproval) AssessCollateral() {
	fmt.Println("  [Personal Loan] Assessing personal assets")
}

func (p *PersonalLoanApproval) ApproveLoan() {
	fmt.Printf("  [Personal Loan] Loan approved for %s\n", p.applicantName)
}

func (p *PersonalLoanApproval) RejectLoan() {
	fmt.Printf("  [Personal Loan] Loan rejected for %s\n", p.applicantName)
}

type BusinessLoanApproval struct {
	BaseLoanApproval
	businessName string
}

func NewBusinessLoanApproval(businessName string) *BusinessLoanApproval {
	approval := &BusinessLoanApproval{businessName: businessName}
	approval.BaseLoanApproval.approval = approval
	return approval
}

func (b *BusinessLoanApproval) CheckCreditScore() {
	fmt.Println("  [Business Loan] Checking business credit history")
}

func (b *BusinessLoanApproval) VerifyIncome() {
	fmt.Println("  [Business Loan] Verifying business financial statements")
}

func (b *BusinessLoanApproval) AssessCollateral() {
	fmt.Println("  [Business Loan] Assessing business assets and guarantees")
}

func (b *BusinessLoanApproval) ApproveLoan() {
	fmt.Printf("  [Business Loan] Loan approved for %s\n", b.businessName)
}

func (b *BusinessLoanApproval) RejectLoan() {
	fmt.Printf("  [Business Loan] Loan rejected for %s\n", b.businessName)
}

func main() {
	fmt.Println("=== Template Method Pattern: JoshBank KYC & Loan Approval ===")

	// Example 1: KYC Verification
	fmt.Println("\n--- Example 1: KYC Verification ---")

	personalKYC := NewPersonalAccountKYC("John Doe")
	personalKYC.Verify()

	businessKYC := NewBusinessAccountKYC("Tech Corp Inc.")
	businessKYC.Verify()

	// Example 2: Loan Approval
	fmt.Println("\n--- Example 2: Loan Approval Process ---")

	personalLoan := NewPersonalLoanApproval("Jane Smith")
	personalLoan.Process()

	businessLoan := NewBusinessLoanApproval("Manufacturing LLC")
	businessLoan.Process()

	fmt.Println("\n✓ Template method defines algorithm skeleton")
	fmt.Println("✓ Subclasses override specific steps")
	fmt.Println("✓ Promotes code reuse")
	fmt.Println("✓ Enforces algorithm structure")
	fmt.Println("✓ JoshBank can standardize processes while allowing customization")
}
