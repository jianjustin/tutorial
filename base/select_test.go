package base

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestForSelect(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch1 <- 42
		}
	}()
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			ch2 <- "hello"
		}
	}()

	timeout := time.After(2000 * time.Millisecond)
	for {
		select {
		case v := <-ch1:
			fmt.Println("Received from ch1:", v)
		case v := <-ch2:
			fmt.Println("Received from ch2:", v)
		case <-timeout:
			fmt.Println("Timeout!")
			timeout = time.After(2000 * time.Millisecond)
			//return
		}
	}
}

func TestForNoCase(t *testing.T) {
	go func() {
		for {
			fmt.Println("Running in goroutine...")
			time.Sleep(1 * time.Second)
		}
	}()
	select {}
}

func TestForLogFull(t *testing.T) {
	logCh := make(chan string, 5)

	// 模拟日志生产
	for i := 0; i < 10; i++ {
		select {
		case logCh <- fmt.Sprintf("log-%d", i):
			fmt.Println("log queued:", i)
		default:
			fmt.Println("log dropped:", i) // 队列满就丢弃
		}
	}
}

func TestForConsumer(t *testing.T) {
	consumer := func(ch <-chan int) {
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					fmt.Println("Channel closed, consumer exiting")
					return
				}
				fmt.Println("Consumed", v)
			}
		}
	}
	producer := func(ch chan<- int) {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println("Produced", i)
		}
		close(ch)
		fmt.Println("Producer finished")
	}

	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
