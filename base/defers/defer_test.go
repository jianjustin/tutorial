package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestForDefer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("i = %d, address of i = %v\n", i, &i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(2 * time.Second)
			fmt.Println("111Hello, World!", i, &i)
		}
	}()
	wg.Wait()

	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println("goroutine", i)
		}(i)
	}

}

func modifyValue(x int) {
	x = x + 10
	fmt.Println("Inside function:", x) // 20
}

func main() {
	a := 10
	modifyValue(a)
	fmt.Println("Outside function:", a) // 10 (未改变)
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function3()
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function3() {
	fmt.Printf("Function3: Defered\n")
}

func function2() {
	fmt.Printf("Function2: Defered until the end of the calling function\n")
}
