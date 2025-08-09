package pipeline

import (
	"fmt"
	"testing"
)

func TestForPipeline1(t *testing.T) {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func TestForPipeline2(t *testing.T) {
	in := gen(2, 3)

	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
