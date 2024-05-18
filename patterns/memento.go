package patterns

type Memento struct {
	state string
}

func (m *Memento) GetSavedState() string {
	return m.state
}

type Originator struct {
	state string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.state = m.GetSavedState()
}

type Caretaker struct {
	memento *Memento
}
