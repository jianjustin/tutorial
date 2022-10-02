package main

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("Begin longwait")
	time.Sleep(5 * 1e9)
	fmt.Println("end longwait")
}

func shortWait() {
	fmt.Println("Begin shortwait")
	time.Sleep(2 * 1e9)
	fmt.Println("end shortwait")
}

func main() {
	fmt.Println("In main")
	go longWait()
	go shortWait()
	fmt.Println("about to sleep in main")
	time.Sleep(10 * 1e9)
	fmt.Println("at the end of main")
}
