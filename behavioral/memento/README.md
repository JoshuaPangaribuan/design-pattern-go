# Memento Pattern

## Problem Statement

When you need to save and restore object state:
- Want undo/redo functionality
- Need to capture state without violating encapsulation
- State should be saved externally
- Rollback to previous states required

## Real-World Scenario

**JoshBank Account State Management**: JoshBank needs to support transaction rollback and account state restoration. Memento captures account balance at each transaction, allowing the system to revert to previous states for error recovery or audit purposes without exposing internal account structure.

## Core Components

1. **Originator**: Object whose state needs to be saved (Account)
2. **Memento**: Stores originator's state (AccountMemento)
3. **Caretaker**: Manages mementos (TransactionHistory)

## Diagrams

### Class Diagram

```mermaid
classDiagram
    class Account {
        -accountID string
        -balance float64
        +Deposit(amount)
        +Withdraw(amount)
        +Save() AccountMemento
        +Restore(memento)
    }
    class AccountMemento {
        -balance float64
        -timestamp Time
        +GetTimestamp() Time
    }
    class TransactionHistory {
        -mementos List~AccountMemento~
        -current int
        +Save(memento)
        +Undo() AccountMemento
        +Redo() AccountMemento
    }
    
    Account --> AccountMemento : creates
    TransactionHistory --> AccountMemento : stores
    Account ..> AccountMemento : restores from
```

### Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant Account
    participant History as TransactionHistory
    participant Memento as AccountMemento
    
    Client->>Account: Deposit($500)
    Account->>Account: balance = $1500
    Client->>Account: Save()
    Account->>Memento: Create(balance=$1500)
    Account-->>Client: memento
    Client->>History: Save(memento)
    
    Note over Client,History: Later...
    
    Client->>Account: Deposit($200)
    Account->>Account: balance = $1700
    Client->>Account: Save()
    Account->>Memento: Create(balance=$1700)
    Client->>History: Save(memento)
    
    Note over Client,History: Undo operation...
    
    Client->>History: Undo()
    History-->>Client: memento(balance=$1500)
    Client->>Account: Restore(memento)
    Account->>Account: balance = $1500
```

## When to Use

✅ **Use when:**
- Need to save/restore object state
- Want undo/redo functionality
- Direct state access violates encapsulation

## Running the Example

```bash
cd behavioral/memento
go run main.go
```

## Key Takeaways

- Memento captures and restores object state
- Enables undo/redo functionality
- Preserves encapsulation
- Common pattern for state management
