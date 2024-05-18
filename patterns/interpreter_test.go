package patterns

import (
	"fmt"
	"strconv"
	"testing"
)

func TestForInterpreter(t *testing.T) {
	// represents the expression -1 + 2
	expression := &Plus{
		left:  &Minus{left: &Number{value: 0}, right: &Number{value: 1}},
		right: &Number{value: 2},
	}

	fmt.Println(strconv.Itoa(expression.Interpret())) // prints: 1
}
