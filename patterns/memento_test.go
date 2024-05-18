package patterns

import "testing"

func TestForMemento(t *testing.T) {
	originator := &Originator{state: "state1"}
	caretaker := &Caretaker{memento: originator.CreateMemento()}

	if originator.state != "state1" {
		t.Errorf("originator.state is not state1: %s", originator.state)
	}

	originator.state = "state2"
	if originator.state != "state2" {
		t.Errorf("originator.state is not state2: %s", originator.state)
	}

	originator.RestoreMemento(caretaker.memento)
	if originator.state != "state1" {
		t.Errorf("originator.state is not state1: %s", originator.state)
	}
}
