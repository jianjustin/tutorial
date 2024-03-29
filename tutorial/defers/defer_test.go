package main

import (
	"fmt"
	"testing"
)

func TestForDefer(t *testing.T) {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2: Defered until the end of the calling function")
}
