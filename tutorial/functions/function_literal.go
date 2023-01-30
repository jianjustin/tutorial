package functions

import (
	"fmt"
	"testing"
)

func TestingForFunctionLiteral(t *testing.T) {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
