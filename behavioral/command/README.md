# Command Pattern

## Problem Statement

When you need to parameterize objects with operations:
- Want to queue, log, or undo operations
- Need to decouple invoker from receiver
- Operations should be first-class objects
- Support for macro commands (composite operations)

## Real-World Scenario

**JoshBank Transaction Controller**: JoshBank needs to execute various banking operations (deposits, withdrawals, transfers) that can be undone, logged for audit, or grouped into composite operations (e.g., "Bill Payment" executes multiple withdrawals). The Command pattern encapsulates each operation as an object, enabling undo/redo functionality and transaction logging.

## Core Components

1. **Command Interface**: Declares execute and undo methods
2. **Concrete Commands**: Implement specific operations
3. **Receiver**: The object that performs the actual work (Account, TransferService)
4. **Invoker**: Asks command to execute (BankingController)
5. **Client**: Creates and configures commands

## When to Use

✅ **Use when:**
- Need to parameterize objects with operations
- Want to queue operations
- Support undo/redo functionality
- Log operations for auditing

⚠️ **Cautions:**
- Increases number of classes
- Can be overkill for simple operations

## Running the Example

```bash
cd behavioral/command
go run main.go
```

## Key Takeaways

- Command encapsulates requests as objects
- Enables undo/redo and audit logging
- Decouples invoker from receiver
- Supports macro commands for complex operations
