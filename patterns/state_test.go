package patterns_test

import (
	statenew "go.guide/patterns"
	"testing"
)

func TestState(t *testing.T) {
	context := statenew.ContextA{}

	context.SetState(&statenew.StateA{})
	context.SetState(&statenew.StateB{})
}
