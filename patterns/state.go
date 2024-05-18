package patterns

import "fmt"

type State interface {
	Handle()
}

type StateA struct{}

func (s *StateA) Handle() {
	fmt.Println("state A Handle")
}

type StateB struct{}

func (s *StateB) Handle() {
	fmt.Println("state B Handle")
}

type Context interface {
	SetState(s State)
	GetCurrentState() State
}

type ContextA struct {
	currentState State
}

func (c *ContextA) SetState(s State) {
	c.currentState = s
	c.currentState.Handle()
}

func (c *ContextA) GetCurrentState() State {
	return c.currentState
}
