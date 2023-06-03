package strategy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.guide/patterns/strategy"
)

func TestStrategyAExecute_ExecutesStrategyA(t *testing.T) {

	strategy := strategy.NewStrategyA()
	strategy.Execute()

}

func TestNewStrategyB_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	strategy := strategy.NewStrategyB()
	assert.NotNil(t, strategy)
}

func TestStrategyBExecute_ExecutesStrategyB(t *testing.T) {

	strategy := strategy.NewStrategyB()
	strategy.Execute()

}

func TestNewContext_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	context := strategy.NewContext()
	assert.NotNil(t, context)
}

func TestSetStrategy_SetsStrategy(t *testing.T) {
	t.Parallel()
	s := strategy.NewStrategyB()
	context := strategy.NewContext()
	context.SetStrategy(s)
	assert.Equal(t, s, context.S)
}

func TestContextExecute_ExecutesSetStrategy(t *testing.T) {

	s := strategy.NewStrategyB()
	context := strategy.NewContext()
	context.SetStrategy(s)
	context.Execute()

}
