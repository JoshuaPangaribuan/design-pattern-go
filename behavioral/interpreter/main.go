package main

import (
	"fmt"
	"strconv"
)

// Expression is the abstract expression interface
type Expression interface {
	Interpret(context map[string]float64) float64
	ToString() string
}

// --- Terminal Expressions ---

// AccountBalance is a terminal expression representing an account balance
type AccountBalance struct {
	accountID string
}

func (a *AccountBalance) Interpret(context map[string]float64) float64 {
	return context[a.accountID]
}

func (a *AccountBalance) ToString() string {
	return a.accountID
}

// Number is a terminal expression
type Number struct {
	value float64
}

func (n *Number) Interpret(context map[string]float64) float64 {
	return n.value
}

func (n *Number) ToString() string {
	return strconv.FormatFloat(n.value, 'f', 2, 64)
}

// --- Non-Terminal Expressions ---

// Add represents addition operation
type Add struct {
	left  Expression
	right Expression
}

func (a *Add) Interpret(context map[string]float64) float64 {
	return a.left.Interpret(context) + a.right.Interpret(context)
}

func (a *Add) ToString() string {
	return fmt.Sprintf("(%s + %s)", a.left.ToString(), a.right.ToString())
}

// Subtract represents subtraction operation
type Subtract struct {
	left  Expression
	right Expression
}

func (s *Subtract) Interpret(context map[string]float64) float64 {
	return s.left.Interpret(context) - s.right.Interpret(context)
}

func (s *Subtract) ToString() string {
	return fmt.Sprintf("(%s - %s)", s.left.ToString(), s.right.ToString())
}

// Multiply represents multiplication operation
type Multiply struct {
	left  Expression
	right Expression
}

func (m *Multiply) Interpret(context map[string]float64) float64 {
	return m.left.Interpret(context) * m.right.Interpret(context)
}

func (m *Multiply) ToString() string {
	return fmt.Sprintf("(%s * %s)", m.left.ToString(), m.right.ToString())
}

// GreaterThan represents comparison operation
type GreaterThan struct {
	left  Expression
	right Expression
}

func (g *GreaterThan) Interpret(context map[string]float64) float64 {
	if g.left.Interpret(context) > g.right.Interpret(context) {
		return 1.0
	}
	return 0.0
}

func (g *GreaterThan) ToString() string {
	return fmt.Sprintf("(%s > %s)", g.left.ToString(), g.right.ToString())
}

func evaluateExpression(expr Expression, context map[string]float64) {
	fmt.Printf("Expression: %s\n", expr.ToString())
	fmt.Printf("Context: %v\n", context)
	result := expr.Interpret(context)
	fmt.Printf("Result: %.2f\n\n", result)
}

func main() {
	fmt.Println("=== Interpreter Pattern: JoshBank Transaction Query Language ===")

	// Example 1: Simple balance calculation
	fmt.Println("\n--- Example 1: Balance Calculation ---")
	
	// ACC001 + 500
	expr1 := &Add{
		left:  &AccountBalance{accountID: "ACC001"},
		right: &Number{value: 500.0},
	}
	
	context1 := map[string]float64{"ACC001": 1000.0}
	evaluateExpression(expr1, context1)

	// Example 2: Complex calculation
	fmt.Println("--- Example 2: Complex Calculation ---")
	
	// (ACC001 + ACC002) * 0.1
	expr2 := &Multiply{
		left: &Add{
			left:  &AccountBalance{accountID: "ACC001"},
			right: &AccountBalance{accountID: "ACC002"},
		},
		right: &Number{value: 0.1},
	}
	
	context2 := map[string]float64{"ACC001": 5000.0, "ACC002": 3000.0}
	evaluateExpression(expr2, context2)

	// Example 3: Balance comparison
	fmt.Println("--- Example 3: Balance Comparison ---")
	
	// ACC001 > 10000
	expr3 := &GreaterThan{
		left:  &AccountBalance{accountID: "ACC001"},
		right: &Number{value: 10000.0},
	}
	
	context3 := map[string]float64{"ACC001": 15000.0}
	evaluateExpression(expr3, context3)
	
	context4 := map[string]float64{"ACC001": 5000.0}
	evaluateExpression(expr3, context4)

	// Example 4: Net worth calculation
	fmt.Println("--- Example 4: Net Worth Calculation ---")
	
	// (ACC001 + ACC002) - ACC003
	expr4 := &Subtract{
		left: &Add{
			left:  &AccountBalance{accountID: "ACC001"},
			right: &AccountBalance{accountID: "ACC002"},
		},
		right: &AccountBalance{accountID: "ACC003"},
	}
	
	context5 := map[string]float64{
		"ACC001": 10000.0,
		"ACC002": 5000.0,
		"ACC003": 2000.0,
	}
	evaluateExpression(expr4, context5)

	fmt.Println("✓ Interpreter pattern represents query grammar as class hierarchy")
	fmt.Println("✓ Easy to change and extend query language")
	fmt.Println("✓ Each grammar rule is a separate class")
	fmt.Println("✓ Useful for transaction query language and rule engine")
	fmt.Println("✓ JoshBank can evaluate complex financial expressions")
}
