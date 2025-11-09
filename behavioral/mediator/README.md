# Mediator Pattern

## Problem Statement

When objects communicate directly, it creates tight coupling:
- Many-to-many relationships become complex
- Hard to understand and maintain interactions
- Difficult to reuse objects independently
- Changes ripple through multiple classes

## Real-World Scenario

**JoshBank Transaction Coordination**: JoshBank has multiple services (Payment Service, Notification Service, Audit Service, Compliance Service) that need to coordinate when processing transactions. Instead of each service communicating directly with others, they communicate through a Transaction Coordinator (mediator), which manages all interactions and prevents tight coupling.

## Core Components

1. **Mediator Interface**: Defines communication interface (BankingMediator)
2. **Concrete Mediator**: Coordinates communication between colleagues (TransactionCoordinator)
3. **Colleague**: Objects that communicate through mediator (PaymentService, NotificationService, etc.)

## Diagrams

### Class Diagram

```mermaid
classDiagram
    class BankingMediator {
        <<Interface>>
        +Notify(sender, event, data)
    }
    class Component {
        <<Interface>>
        +SetMediator(mediator)
    }
    class TransactionCoordinator {
        -paymentService PaymentService
        -notificationService NotificationService
        -auditService AuditService
        -complianceService ComplianceService
        +Notify(sender, event, data)
    }
    class PaymentService {
        -mediator BankingMediator
        +SetMediator(mediator)
        +ProcessPayment(id, customerID, amount)
    }
    class NotificationService {
        -mediator BankingMediator
        +SetMediator(mediator)
        +SendNotification(recipient, message)
    }
    class AuditService {
        -mediator BankingMediator
        +SetMediator(mediator)
        +LogTransaction(id, amount)
    }
    class ComplianceService {
        -mediator BankingMediator
        +SetMediator(mediator)
        +CheckTransaction(id, amount)
    }
    
    BankingMediator <|.. TransactionCoordinator
    Component <|.. PaymentService
    Component <|.. NotificationService
    Component <|.. AuditService
    Component <|.. ComplianceService
    TransactionCoordinator --> PaymentService : coordinates
    TransactionCoordinator --> NotificationService : coordinates
    TransactionCoordinator --> AuditService : coordinates
    TransactionCoordinator --> ComplianceService : coordinates
    PaymentService --> BankingMediator : uses
    NotificationService --> BankingMediator : uses
    AuditService --> BankingMediator : uses
    ComplianceService --> BankingMediator : uses
```

### Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant PaymentService
    participant Coordinator as TransactionCoordinator
    participant AuditService
    participant NotificationService
    participant ComplianceService
    
    Client->>PaymentService: ProcessPayment(TXN001, $500)
    PaymentService->>Coordinator: Notify(payment_processed)
    Coordinator->>AuditService: LogTransaction(TXN001, $500)
    Coordinator->>NotificationService: SendNotification(customer, message)
    Coordinator->>ComplianceService: CheckTransaction(TXN001, $500)
    ComplianceService->>ComplianceService: Check amount <= $10000
    ComplianceService-->>Coordinator: Passed
```

## When to Use

✅ **Use when:**
- Objects communicate in complex ways
- Reusing objects is difficult due to dependencies
- Behavior distributed across classes should be customizable

## Running the Example

```bash
cd behavioral/mediator
go run main.go
```

## Key Takeaways

- Mediator centralizes complex communications
- Reduces coupling between objects
- Makes interactions easier to understand and maintain
- Common pattern for coordinating multiple subsystems
