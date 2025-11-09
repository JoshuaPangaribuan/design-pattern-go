# Strategy Pattern

## Problem Statement

When you have multiple algorithms for a task:
- Need to switch between algorithms at runtime
- Algorithms should be interchangeable
- Don't want conditional logic for algorithm selection
- Want to isolate algorithm implementation

## Real-World Scenario

**JoshBank Interest & Fee Calculation**: JoshBank offers different interest calculation methods (Simple, Compound, Tiered) and fee structures (Flat Fee, Percentage Fee). Customers can choose or the bank can switch strategies based on account type. Strategy pattern encapsulates each calculation algorithm, making them interchangeable.

## Core Components

1. **Strategy Interface**: Defines algorithm interface (InterestCalculationStrategy)
2. **Concrete Strategies**: Implement different algorithms (SimpleInterestStrategy, CompoundInterestStrategy)
3. **Context**: Uses a strategy to execute algorithm (Account)

## Diagrams

### Class Diagram

```mermaid
classDiagram
    class InterestCalculationStrategy {
        <<Interface>>
        +Calculate(balance, rate, time) float64
    }
    class SimpleInterestStrategy {
        +Calculate(balance, rate, time) float64
    }
    class CompoundInterestStrategy {
        +Calculate(balance, rate, time) float64
    }
    class TieredInterestStrategy {
        +Calculate(balance, rate, time) float64
    }
    class Account {
        -accountID string
        -balance float64
        -strategy InterestCalculationStrategy
        +SetStrategy(strategy)
        +CalculateInterest(rate, time) float64
    }
    
    InterestCalculationStrategy <|.. SimpleInterestStrategy
    InterestCalculationStrategy <|.. CompoundInterestStrategy
    InterestCalculationStrategy <|.. TieredInterestStrategy
    Account --> InterestCalculationStrategy : uses
```

### Sequence Diagram

```mermaid
sequenceDiagram
    participant Client
    participant Account
    participant Strategy as SimpleInterestStrategy
    
    Client->>Account: SetStrategy(SimpleInterestStrategy)
    Account->>Account: strategy = SimpleInterestStrategy
    
    Note over Client,Strategy: Later...
    
    Client->>Account: CalculateInterest(rate=0.05, time=1)
    Account->>Strategy: Calculate(balance, rate, time)
    Strategy->>Strategy: balance * rate * time
    Strategy-->>Account: interest amount
    Account-->>Client: interest
```

## When to Use

✅ **Use when:**
- Many related classes differ only in behavior
- Need different variants of an algorithm
- Algorithm uses data clients shouldn't know about

## Running the Example

```bash
cd behavioral/strategy
go run main.go
```

## Key Takeaways

- Strategy defines family of algorithms
- Makes algorithms interchangeable
- Eliminates conditional statements
- Easy to add new strategies
