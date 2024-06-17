package channels

import (
	"fmt"
	"testing"
	"time"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		if i < 10 {
			time.Sleep(time.Duration(100))
			fmt.Sprintf("waiting for %d\n", time.Duration(100))
		}
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func TestForChannelBlock(t *testing.T) {
	ch1 := make(chan int)
	go pump(ch1)
	//fmt.Println(<-ch1)
	go suck(ch1)
	time.Sleep(1e9)
}
