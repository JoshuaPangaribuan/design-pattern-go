# Design Patterns in Go - JoshBank Fintech Platform

A comprehensive collection of Gang of Four (GoF) design patterns implemented in Go, demonstrated through real-world fintech scenarios in the **JoshBank** platform.

## 🎯 Purpose

This repository serves as a practical guide to understanding and implementing design patterns in Go. Each pattern is demonstrated through **JoshBank**, a modern fintech platform that handles banking operations, payments, risk management, compliance, and customer services. Each pattern includes:
- Real-world fintech scenarios from JoshBank
- Clear explanations of core components
- When and why to use the pattern
- Implementation best practices

## 🏦 About JoshBank

**JoshBank** is a comprehensive fintech platform that demonstrates design patterns through various banking and financial services:

- **Customer Management**: Account creation, KYC (Know Your Customer) verification, customer onboarding
- **Payment Processing**: Multiple payment gateways, transaction processing, refunds
- **Risk & Compliance**: Fraud detection, transaction monitoring, regulatory compliance checks
- **Financial Products**: Account types (checking, savings, investment), loan processing, trading
- **Notifications**: Transaction alerts, account updates, system notifications
- **Integration**: Legacy bank systems, third-party payment providers, external services

## 📚 Pattern Categories

Design patterns are organized into three main categories:

### 1. Creational Patterns
Focus on object creation mechanisms, increasing flexibility and reuse of existing code.

- **Abstract Factory** - Create families of related banking products (accounts, loans, cards)
- **Builder** - Construct complex customer profiles and account configurations step by step
- **Factory Method** - Define interface for creating different payment methods and account types
- **Prototype** - Clone existing account templates and transaction configurations
- **Singleton** - Ensure a single instance of configuration manager, audit logger, or risk service

### 2. Structural Patterns
Deal with object composition, creating relationships between objects to form larger structures.

- **Adapter** - Make incompatible interfaces work together (legacy bank systems, payment gateways)
- **Bridge** - Separate account abstraction from payment processing implementation
- **Composite** - Compose account hierarchies and transaction groups into tree structures
- **Decorator** - Add responsibilities to transactions dynamically (logging, validation, encryption)
- **Facade** - Provide simplified interface to complex banking subsystems
- **Flyweight** - Share common state between multiple transaction objects (currency, exchange rates)
- **Proxy** - Provide placeholder for expensive operations (compliance checks, risk assessments)

### 3. Behavioral Patterns
Concerned with algorithms and assignment of responsibilities between objects.

- **Chain of Responsibility** - Pass transaction requests along fraud detection and approval chains
- **Command** - Encapsulate banking operations as objects (transfers, withdrawals, deposits)
- **Iterator** - Access transaction history and account lists sequentially without exposing representation
- **Mediator** - Reduce chaotic dependencies between banking services (payments, notifications, compliance)
- **Memento** - Save and restore account state for transaction rollbacks and audit trails
- **Observer** - Notify multiple services about transaction events and account changes
- **State** - Alter account behavior when internal state changes (active, frozen, closed)
- **Strategy** - Define family of algorithms and make them interchangeable (interest calculation, fee structures)
- **Template Method** - Define skeleton of algorithm in base class (KYC verification, loan approval)
- **Visitor** - Separate algorithms from objects they operate on (transaction reporting, account analysis)
- **Interpreter** - Define grammar and interpreter for a language (transaction query language, rule engine)

## 🚀 Getting Started

### Prerequisites
- Go 1.24.3 or higher

### Installation
```bash
git clone https://github.com/JoshuaPangaribuan/design-pattern-go.git
cd design-pattern-go
```

### Running Examples

Each pattern directory contains runnable examples. Navigate to any pattern folder and run:

```bash
# Example: Running the Singleton pattern
cd creational/singleton
go run main.go
```

### Building All Patterns

To verify all patterns compile correctly:

```bash
go build ./...
```

## 📖 How to Use This Repository

1. **Browse by Category**: Navigate to `creational/`, `structural/`, or `behavioral/` folders
2. **Read Pattern Documentation**: Each pattern has its own `README.md` explaining:
   - The problem it solves
   - Core components and structure
   - Real-world use cases
   - Implementation walkthrough
3. **Study the Code**: Review the Go implementation with inline comments
4. **Run Examples**: Execute the code to see patterns in action

## 🤝 Contributing

Contributions are welcome! Feel free to:
- Report issues or suggest improvements
- Add more real-world examples
- Improve documentation
- Fix bugs or typos

## 📄 License

This project is open source and available for educational purposes.

## 📚 References

- "Design Patterns: Elements of Reusable Object-Oriented Software" by Gang of Four
- Go best practices and idioms

