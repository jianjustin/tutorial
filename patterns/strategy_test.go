package patterns_test

import (
	"go.guide/patterns"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrategyAExecute_ExecutesStrategyA(t *testing.T) {

	strategy := patterns.NewStrategyA()
	strategy.Execute()

}

func TestNewStrategyB_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	strategy := patterns.NewStrategyB()
	assert.NotNil(t, strategy)
}

func TestStrategyBExecute_ExecutesStrategyB(t *testing.T) {

	strategy := patterns.NewStrategyB()
	strategy.Execute()

}

func TestNewContext_ReturnsNonNil(t *testing.T) {
	t.Parallel()
	context := patterns.NewContext()
	assert.NotNil(t, context)
}

func TestSetStrategy_SetsStrategy(t *testing.T) {
	t.Parallel()
	s := patterns.NewStrategyB()
	context := patterns.NewContext()
	context.SetStrategy(s)
	assert.Equal(t, s, context.S)
}

func TestContextExecute_ExecutesSetStrategy(t *testing.T) {

	s := patterns.NewStrategyB()
	context := patterns.NewContext()
	context.SetStrategy(s)
	context.Execute()

}
