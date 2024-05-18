package patterns

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Aggregate interface
type Aggregate interface {
	Iterator() Iterator
}

// ConcreteAggregate struct
type ConcreteAggregate struct {
	items []interface{}
}

// Iterator method of ConcreteAggregate
func (c *ConcreteAggregate) Iterator() Iterator {
	return &ConcreteIterator{
		aggregate: c,
		index:     0,
	}
}

// ConcreteIterator struct
type ConcreteIterator struct {
	aggregate *ConcreteAggregate
	index     int
}

// HasNext method of ConcreteIterator
func (c *ConcreteIterator) HasNext() bool {
	if c.index < len(c.aggregate.items) {
		return true
	}
	return false
}

// Next method of ConcreteIterator
func (c *ConcreteIterator) Next() interface{} {
	if c.HasNext() {
		item := c.aggregate.items[c.index]
		c.index++
		return item
	}
	return nil
}
