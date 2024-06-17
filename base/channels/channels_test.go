package channels

import (
	"fmt"
	"testing"
)

func TestChannels(t *testing.T) {
	message := make(chan string)
	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)
}
