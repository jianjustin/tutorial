package patterns

import "fmt"

// Mediator interface
type Mediator interface {
	Notify(sender Colleague, event string)
}

// ConcreteMediator struct
type ConcreteMediator struct {
	colleague1 *Colleague1
	colleague2 *Colleague2
}

// Notify method of ConcreteMediator
func (m *ConcreteMediator) Notify(sender Colleague, event string) {
	switch sender.(type) {
	case *Colleague1:
		if event == "A" {
			fmt.Println("Mediator reacts on A and triggers following operations:")
			m.colleague2.DoB()
		}
	case *Colleague2:
		if event == "D" {
			fmt.Println("Mediator reacts on D and triggers following operations:")
			m.colleague1.DoC()
			m.colleague2.DoB()
		}
	}
}

// Colleague interface
type Colleague interface {
	SetMediator(mediator Mediator)
}

// Colleague1 struct
type Colleague1 struct {
	mediator Mediator
}

// SetMediator method of Colleague1
func (c *Colleague1) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

// DoA method of Colleague1
func (c *Colleague1) DoA() {
	fmt.Println("Colleague1 does A.")
	c.mediator.Notify(c, "A")
}

// DoC method of Colleague1
func (c *Colleague1) DoC() {
	fmt.Println("Colleague1 does C.")
}

// Colleague2 struct
type Colleague2 struct {
	mediator Mediator
}

// SetMediator method of Colleague2
func (c *Colleague2) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

// DoB method of Colleague2
func (c *Colleague2) DoB() {
	fmt.Println("Colleague2 does B.")
	c.mediator.Notify(c, "B")
}

// DoD method of Colleague2
func (c *Colleague2) DoD() {
	fmt.Println("Colleague2 does D.")
	c.mediator.Notify(c, "D")
}
