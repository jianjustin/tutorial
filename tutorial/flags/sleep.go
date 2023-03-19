package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	period := flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("%v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
