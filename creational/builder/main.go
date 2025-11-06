package main

import (
	"fmt"
	"time"
)

// CustomerProfile is the complex product we're building.
// It has many optional fields that can be configured.
type CustomerProfile struct {
	customerID      string
	firstName       string
	lastName        string
	email           string
	phone           string
	dateOfBirth     time.Time
	address         string
	city            string
	state           string
	zipCode         string
	accountType     string
	kycStatus       string
	riskLevel       string
	preferences     map[string]string
	metadata        map[string]interface{}
}

// Display simulates displaying the customer profile
func (c *CustomerProfile) Display() {
	fmt.Println("\n--- Customer Profile ---")
	fmt.Printf("Customer ID: %s\n", c.customerID)
	fmt.Printf("Name: %s %s\n", c.firstName, c.lastName)
	fmt.Printf("Email: %s\n", c.email)
	fmt.Printf("Phone: %s\n", c.phone)
	if !c.dateOfBirth.IsZero() {
		fmt.Printf("Date of Birth: %s\n", c.dateOfBirth.Format("2006-01-02"))
	}
	fmt.Printf("Address: %s, %s, %s %s\n", c.address, c.city, c.state, c.zipCode)
	fmt.Printf("Account Type: %s\n", c.accountType)
	fmt.Printf("KYC Status: %s\n", c.kycStatus)
	fmt.Printf("Risk Level: %s\n", c.riskLevel)
	if len(c.preferences) > 0 {
		fmt.Println("Preferences:")
		for k, v := range c.preferences {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
	fmt.Println("✓ Profile created successfully")
}

// CustomerProfileBuilder is the builder that constructs CustomerProfile objects.
// It provides a fluent interface for setting various options.
type CustomerProfileBuilder struct {
	profile *CustomerProfile
}

// NewCustomerProfileBuilder creates a new builder with default values
func NewCustomerProfileBuilder() *CustomerProfileBuilder {
	return &CustomerProfileBuilder{
		profile: &CustomerProfile{
			accountType: "personal",
			kycStatus:   "pending",
			riskLevel:   "low",
			preferences: make(map[string]string),
			metadata:    make(map[string]interface{}),
		},
	}
}

// SetCustomerID sets the customer ID (required)
func (b *CustomerProfileBuilder) SetCustomerID(id string) *CustomerProfileBuilder {
	b.profile.customerID = id
	return b
}

// SetPersonalInfo sets basic personal information
func (b *CustomerProfileBuilder) SetPersonalInfo(firstName, lastName, email string) *CustomerProfileBuilder {
	b.profile.firstName = firstName
	b.profile.lastName = lastName
	b.profile.email = email
	return b
}

// SetPhone sets the phone number
func (b *CustomerProfileBuilder) SetPhone(phone string) *CustomerProfileBuilder {
	b.profile.phone = phone
	return b
}

// SetDateOfBirth sets the date of birth
func (b *CustomerProfileBuilder) SetDateOfBirth(dob time.Time) *CustomerProfileBuilder {
	b.profile.dateOfBirth = dob
	return b
}

// SetAddress sets the address information
func (b *CustomerProfileBuilder) SetAddress(address, city, state, zipCode string) *CustomerProfileBuilder {
	b.profile.address = address
	b.profile.city = city
	b.profile.state = state
	b.profile.zipCode = zipCode
	return b
}

// SetAccountType sets the account type
func (b *CustomerProfileBuilder) SetAccountType(accountType string) *CustomerProfileBuilder {
	b.profile.accountType = accountType
	return b
}

// SetKYCStatus sets the KYC verification status
func (b *CustomerProfileBuilder) SetKYCStatus(status string) *CustomerProfileBuilder {
	b.profile.kycStatus = status
	return b
}

// SetRiskLevel sets the risk assessment level
func (b *CustomerProfileBuilder) SetRiskLevel(level string) *CustomerProfileBuilder {
	b.profile.riskLevel = level
	return b
}

// AddPreference adds a customer preference
func (b *CustomerProfileBuilder) AddPreference(key, value string) *CustomerProfileBuilder {
	b.profile.preferences[key] = value
	return b
}

// AddMetadata adds metadata to the profile
func (b *CustomerProfileBuilder) AddMetadata(key string, value interface{}) *CustomerProfileBuilder {
	b.profile.metadata[key] = value
	return b
}

// Build constructs and returns the final CustomerProfile.
// It validates required fields before returning.
func (b *CustomerProfileBuilder) Build() (*CustomerProfile, error) {
	if b.profile.customerID == "" {
		return nil, fmt.Errorf("customer ID is required")
	}
	if b.profile.firstName == "" || b.profile.lastName == "" {
		return nil, fmt.Errorf("first name and last name are required")
	}
	if b.profile.email == "" {
		return nil, fmt.Errorf("email is required")
	}

	// Return a copy to prevent modification after building
	return b.profile, nil
}

// CustomerProfileDirector (optional) can orchestrate complex building sequences
type CustomerProfileDirector struct {
	builder *CustomerProfileBuilder
}

func NewCustomerProfileDirector(builder *CustomerProfileBuilder) *CustomerProfileDirector {
	return &CustomerProfileDirector{builder: builder}
}

// BuildPersonalAccountProfile creates a standard personal account profile
func (d *CustomerProfileDirector) BuildPersonalAccountProfile(customerID, firstName, lastName, email string) *CustomerProfile {
	dob, _ := time.Parse("2006-01-02", "1990-01-01")
	req, _ := d.builder.
		SetCustomerID(customerID).
		SetPersonalInfo(firstName, lastName, email).
		SetAccountType("personal").
		SetKYCStatus("pending").
		SetRiskLevel("low").
		SetDateOfBirth(dob).
		AddPreference("notifications", "email").
		AddPreference("language", "en").
		Build()
	return req
}

// BuildBusinessAccountProfile creates a business account profile
func (d *CustomerProfileDirector) BuildBusinessAccountProfile(customerID, firstName, lastName, email, businessName string) *CustomerProfile {
	req, _ := d.builder.
		SetCustomerID(customerID).
		SetPersonalInfo(firstName, lastName, email).
		SetAccountType("business").
		SetKYCStatus("verified").
		SetRiskLevel("medium").
		AddMetadata("business_name", businessName).
		AddPreference("notifications", "email,sms").
		Build()
	return req
}

func main() {
	fmt.Println("=== Builder Pattern: JoshBank Customer Profile Builder ===")

	// Example 1: Simple personal profile
	fmt.Println("\n--- Example 1: Simple Personal Profile ---")
	simpleProfile, err := NewCustomerProfileBuilder().
		SetCustomerID("CUST001").
		SetPersonalInfo("John", "Doe", "john.doe@example.com").
		SetPhone("555-0100").
		Build()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		simpleProfile.Display()
	}

	// Example 2: Complex profile with all options
	fmt.Println("\n--- Example 2: Complex Profile with All Options ---")
	dob, _ := time.Parse("2006-01-02", "1985-05-15")
	complexProfile, err := NewCustomerProfileBuilder().
		SetCustomerID("CUST002").
		SetPersonalInfo("Jane", "Smith", "jane.smith@example.com").
		SetPhone("555-0200").
		SetDateOfBirth(dob).
		SetAddress("123 Main St", "New York", "NY", "10001").
		SetAccountType("premium").
		SetKYCStatus("verified").
		SetRiskLevel("low").
		AddPreference("notifications", "email,sms,push").
		AddPreference("language", "en").
		AddPreference("currency", "USD").
		AddMetadata("referral_source", "online").
		AddMetadata("signup_date", time.Now()).
		Build()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		complexProfile.Display()
	}

	// Example 3: Using Director for common profile patterns
	fmt.Println("\n--- Example 3: Using Director for Common Patterns ---")

	director := NewCustomerProfileDirector(NewCustomerProfileBuilder())

	personalProfile := director.BuildPersonalAccountProfile(
		"CUST003",
		"Bob",
		"Johnson",
		"bob.johnson@example.com",
	)
	personalProfile.Display()

	// Reset builder for next profile
	director.builder = NewCustomerProfileBuilder()
	businessProfile := director.BuildBusinessAccountProfile(
		"CUST004",
		"Alice",
		"Williams",
		"alice@company.com",
		"Tech Corp Inc.",
	)
	businessProfile.Display()

	// Example 4: Error handling - missing required field
	fmt.Println("\n--- Example 4: Validation ---")
	invalidProfile, err := NewCustomerProfileBuilder().
		SetPersonalInfo("Test", "User", "test@example.com").
		Build() // Missing Customer ID

	if err != nil {
		fmt.Printf("✓ Validation caught error: %v\n", err)
	} else {
		invalidProfile.Display()
	}

	fmt.Println("\n✓ Builder pattern provides fluent, readable API")
	fmt.Println("✓ Complex customer profiles can be constructed step by step")
	fmt.Println("✓ Director can encapsulate common profile creation patterns")
	fmt.Println("✓ JoshBank can easily create different customer profile types")
}
