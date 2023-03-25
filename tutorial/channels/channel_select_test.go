package channels

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	c1 := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
		}
		c1 <- "hello"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	}
}
