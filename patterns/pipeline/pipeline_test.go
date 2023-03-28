package pipeline

import (
	"fmt"
	"testing"
)

func TestPipeline1(t *testing.T) {
	// Set up the pipeline.
	c := Gen(2, 3)
	out := Sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func TestPipeline2(t *testing.T) {
	// Set up the pipeline and consume the output.
	for n := range Sq(Sq(Gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}

func TestMerge(t *testing.T) {
	in := Gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := Sq(in)
	c2 := Sq(in)

	done := make(chan struct{}, 2)

	// Consume the merged output from c1 and c2.
	for n := range Merge(done, c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
	done <- struct{}{}
	done <- struct{}{}
}
