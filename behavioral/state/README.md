# State Pattern

## Problem Statement

When object behavior changes based on internal state:
- Large conditional statements based on state
- State-specific behavior scattered across methods
- Adding new states requires modifying existing code
- State transitions are complex

## Real-World Scenario

**JoshBank Account States**: JoshBank accounts have different states (Active, Frozen, Closed). Available operations and behavior depend on current state. Active accounts allow deposits and withdrawals, frozen accounts only allow deposits, and closed accounts allow no operations. State pattern encapsulates state-specific behavior.

## Core Components

1. **Context**: Maintains current state, delegates to state object (Account)
2. **State Interface**: Defines state-specific behavior (AccountState)
3. **Concrete States**: Implement behavior for each state (ActiveState, FrozenState, ClosedState)

## Diagrams

### Class Diagram

```mermaid
classDiagram
    class AccountState {
        <<Interface>>
        +Deposit(amount) string
        +Withdraw(amount) string
        +Close() string
        +GetStateName() string
    }
    class Account {
        -activeState AccountState
        -frozenState AccountState
        -closedState AccountState
        -currentState AccountState
        -accountID string
        -balance float64
        +Deposit(amount)
        +Withdraw(amount)
        +Close()
        +SetState(state)
    }
    class ActiveState {
        -account Account
        +Deposit(amount) string
        +Withdraw(amount) string
        +Close() string
    }
    class FrozenState {
        -account Account
        +Deposit(amount) string
        +Withdraw(amount) string
        +Close() string
    }
    class ClosedState {
        -account Account
        +Deposit(amount) string
        +Withdraw(amount) string
        +Close() string
    }
    
    AccountState <|.. ActiveState
    AccountState <|.. FrozenState
    AccountState <|.. ClosedState
    Account --> AccountState : currentState
    Account --> ActiveState : creates
    Account --> FrozenState : creates
    Account --> ClosedState : creates
    ActiveState --> Account : references
    FrozenState --> Account : references
    ClosedState --> Account : references
```

### State Transition Diagram

```mermaid
stateDiagram-v2
    [*] --> Active
    Active --> Frozen: Balance < 0
    Frozen --> Active: Balance > 0
    Active --> Closed: Close()
    Frozen --> Closed: Close()
    Closed --> [*]
    
    note right of Active
        Can deposit and withdraw
    end note
    
    note right of Frozen
        Can only deposit
        Auto-unfreeze when balance > 0
    end note
    
    note right of Closed
        No operations allowed
    end note
```

## When to Use

✅ **Use when:**
- Object behavior depends on state
- Large conditionals based on state
- State transitions are complex

## Running the Example

```bash
cd behavioral/state
go run main.go
```

## Key Takeaways

- State pattern encapsulates state-specific behavior
- Eliminates complex conditionals
- Easy to add new states
- Makes state transitions explicit
