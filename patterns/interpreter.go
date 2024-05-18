package patterns

type Expression interface {
	Interpret() int
}

type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

type Plus struct {
	left  Expression
	right Expression
}

func (p *Plus) Interpret() int {
	return p.left.Interpret() + p.right.Interpret()
}

type Minus struct {
	left  Expression
	right Expression
}

func (m *Minus) Interpret() int {
	return m.left.Interpret() - m.right.Interpret()
}
