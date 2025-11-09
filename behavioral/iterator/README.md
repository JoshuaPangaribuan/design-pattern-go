# Iterator Pattern

## Problem Statement

When you need to traverse a collection without exposing its internal structure:
- Different collections have different traversal methods
- Want multiple simultaneous traversals
- Need uniform interface for different collections
- Should hide collection implementation details

## Real-World Scenario

**JoshBank Transaction History**: JoshBank stores transaction history in different data structures (arrays, linked lists, databases). Iterator provides uniform way to traverse transactions regardless of underlying storage, allowing clients to iterate through transaction history without knowing implementation details.

## Core Components

1. **Iterator Interface**: Defines traversal methods (Next, HasNext, Reset)
2. **Concrete Iterator**: Implements traversal for specific collection
3. **Aggregate Interface**: Declares method to create iterator (Collection)
4. **Concrete Aggregate**: Returns appropriate iterator

## Diagrams

### Class Diagram

```mermaid
classDiagram
    class Iterator {
        <<Interface>>
        +HasNext() bool
        +Next() Transaction
        +Reset()
    }
    class Collection {
        <<Interface>>
        +CreateIterator() Iterator
        +Add(transaction)
    }
    class ArrayIterator {
        -history ArrayTransactionHistory
        -index int
        +HasNext() bool
        +Next() Transaction
        +Reset()
    }
    class LinkedListIterator {
        -current TransactionNode
        -head TransactionNode
        +HasNext() bool
        +Next() Transaction
        +Reset()
    }
    class ArrayTransactionHistory {
        -transactions List~Transaction~
        +CreateIterator() Iterator
        +Add(transaction)
    }
    class LinkedListTransactionHistory {
        -head TransactionNode
        +CreateIterator() Iterator
        +Add(transaction)
    }
    class Transaction {
        +ID string
        +Amount float64
        +Type string
    }
    
    Iterator <|.. ArrayIterator
    Iterator <|.. LinkedListIterator
    Collection <|.. ArrayTransactionHistory
    Collection <|.. LinkedListTransactionHistory
    ArrayIterator --> ArrayTransactionHistory : uses
    LinkedListIterator --> LinkedListTransactionHistory : uses
    ArrayTransactionHistory --> Transaction : contains
    LinkedListTransactionHistory --> Transaction : contains
```

### Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant Collection as ArrayTransactionHistory
    participant Iterator as ArrayIterator
    
    Client->>Collection: CreateIterator()
    Collection-->>Client: Iterator instance
    
    loop While HasNext()
        Client->>Iterator: HasNext()
        Iterator-->>Client: true
        Client->>Iterator: Next()
        Iterator-->>Client: Transaction
    end
    
    Client->>Iterator: Reset()
    Iterator->>Iterator: index = 0
```

## When to Use

✅ **Use when:**
- Need to traverse collection without exposing internals
- Support multiple traversal algorithms
- Provide uniform interface for different collections

## Running the Example

```bash
cd behavioral/iterator
go run main.go
```

## Key Takeaways

- Iterator provides uniform traversal interface
- Hides collection implementation details
- Supports multiple simultaneous traversals
- Common pattern in Go for range loops and collections
