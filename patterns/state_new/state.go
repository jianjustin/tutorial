package statenew

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
