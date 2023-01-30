package functions

import (
	"fmt"
	"testing"
)

func TestForFunctionParamter(t *testing.T) {
	callback(1, 2, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, x int, f func(int, int)) {
	f(y, x)
}
