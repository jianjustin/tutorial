package statenew

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
