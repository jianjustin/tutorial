package main

import (
	"fmt"
)

func main() {
	var arr1 [6]int
	var slice1 = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The cap of arr1 is %d\n", cap(slice1))
}
