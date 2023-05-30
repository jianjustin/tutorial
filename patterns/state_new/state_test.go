package statenew_test

import (
	"testing"

	statenew "go.guide/patterns/state_new"
)

func TestState(t *testing.T) {
	context := statenew.ContextA{}

	context.SetState(&statenew.StateA{})
	context.SetState(&statenew.StateB{})
}
