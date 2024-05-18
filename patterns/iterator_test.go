package patterns

import (
	"fmt"
	"testing"
)

func TestForIterator(t *testing.T) {
	aggregate := &ConcreteAggregate{
		items: []interface{}{"a", "b", "c", "d", "e"},
	}
	iterator := aggregate.Iterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
