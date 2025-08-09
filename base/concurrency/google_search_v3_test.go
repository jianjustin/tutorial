package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestForGoogleSearch(t *testing.T) {
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
