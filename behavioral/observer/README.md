# Observer Pattern

## Problem Statement

When one object's state change should notify multiple dependents:
- Many objects depend on another object's state
- Don't know number of dependents in advance
- Want loose coupling between subject and observers
- Broadcast communication needed

## Real-World Scenario

**JoshBank Transaction Monitoring**: When JoshBank processes a transaction, multiple services need to be notified: Notification Service (send alerts), Audit Service (log transaction), Compliance Service (check regulations), Analytics Service (update metrics). Observer pattern allows these services to subscribe to transaction events and react automatically.

## Core Components

1. **Subject**: Maintains list of observers, notifies them of changes (TransactionService)
2. **Observer Interface**: Defines update method
3. **Concrete Observers**: Implement update to react to changes (NotificationService, AuditService, etc.)

## Diagrams

### Class Diagram

```mermaid
classDiagram
    class Subject {
        <<Interface>>
        +RegisterObserver(observer)
        +RemoveObserver(observer)
        +NotifyObservers(id, amount, status)
    }
    class Observer {
        <<Interface>>
        +Update(id, amount, status)
        +GetName() string
    }
    class TransactionService {
        -observers List~Observer~
        +RegisterObserver(observer)
        +RemoveObserver(observer)
        +NotifyObservers(id, amount, status)
        +ProcessTransaction(id, amount)
    }
    class NotificationService {
        +Update(id, amount, status)
        +GetName() string
    }
    class AuditService {
        +Update(id, amount, status)
        +GetName() string
    }
    class ComplianceService {
        +Update(id, amount, status)
        +GetName() string
    }
    class AnalyticsService {
        +Update(id, amount, status)
        +GetName() string
    }
    
    Subject <|.. TransactionService
    Observer <|.. NotificationService
    Observer <|.. AuditService
    Observer <|.. ComplianceService
    Observer <|.. AnalyticsService
    TransactionService --> Observer : notifies
```

### Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant Subject as TransactionService
    participant Observer1 as NotificationService
    participant Observer2 as AuditService
    participant Observer3 as ComplianceService
    
    Client->>Observer1: RegisterObserver()
    Client->>Observer2: RegisterObserver()
    Client->>Observer3: RegisterObserver()
    
    Note over Client,Observer3: Observers registered
    
    Client->>Subject: ProcessTransaction(TXN001, $500)
    Subject->>Subject: Process transaction
    Subject->>Subject: NotifyObservers(TXN001, $500, completed)
    Subject->>Observer1: Update(TXN001, $500, completed)
    Observer1->>Observer1: Send notification
    Subject->>Observer2: Update(TXN001, $500, completed)
    Observer2->>Observer2: Log transaction
    Subject->>Observer3: Update(TXN001, $500, completed)
    Observer3->>Observer3: Check compliance
```

## When to Use

✅ **Use when:**
- Change to one object requires changing others
- Object should notify others without knowing who they are
- Loose coupling between subject and observers needed

## Running the Example

```bash
cd behavioral/observer
go run main.go
```

## Key Takeaways

- Observer enables one-to-many dependencies
- Subject and observers are loosely coupled
- Observers can be added/removed dynamically
- Common pattern for event-driven systems
