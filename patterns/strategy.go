package patterns

import "fmt"

// Strategy defines the interface for the strategy to execute.
type Strategy interface {
	Execute()
}

// strategyA defines an implementation of a Strategy to execute.
type strategyA struct {
}

// NewStrategyA creates a new instance of strategy A.
func NewStrategyA() Strategy {
	return &strategyA{}
}

// Execute executes strategy A.
func (s *strategyA) Execute() {
	fmt.Printf("executing strategy A\n")
}

// strategyB defines an implementation of a Strategy to execute.
type strategyB struct {
}

// NewStrategyB creates a new instance of strategy B.
func NewStrategyB() Strategy {
	return &strategyB{}
}

// Execute executes strategy B.
func (s *strategyB) Execute() {
	fmt.Printf("executing strategy B\n")
}

// Context defines a context for executing a strategy.
type StrategyContext struct {
	S Strategy
}

// NewContext creates a new instance of a context.
func NewContext() *StrategyContext {
	return &StrategyContext{}
}

// SetStrategy sets the strategy to execute for this context.
func (c *StrategyContext) SetStrategy(strategy Strategy) {
	c.S = strategy
}

// Execute executes the strategy.
func (c *StrategyContext) Execute() {
	c.S.Execute()
}
