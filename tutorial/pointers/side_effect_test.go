package pointers

import (
	"fmt"
	"testing"
)

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func TestForSideEffect(t *testing.T) {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply: ", *reply)
}
