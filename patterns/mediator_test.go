package patterns

import (
	"testing"
)

func TestForMediator(t *testing.T) {
	t.Parallel()

	c1 := &Colleague1{}
	c2 := &Colleague2{}
	m := &ConcreteMediator{
		colleague1: c1,
		colleague2: c2,
	}

	c1.SetMediator(m)
	c2.SetMediator(m)

	c1.DoA()
	c2.DoD()
}
