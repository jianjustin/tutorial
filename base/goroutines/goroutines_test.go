package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestF(t *testing.T) {
	go f("goroutines1")
	go f("goroutines2")

	time.Sleep(time.Second)
	fmt.Println("done")
}
