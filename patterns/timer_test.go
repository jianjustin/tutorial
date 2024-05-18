package patterns

import (
	"sync"
	"testing"
	"time"
)

func TestForTimer(t *testing.T) {
	wg := sync.WaitGroup{}
	timer := time.NewTimer(3 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		isDone := make(chan struct{})
		count := int64(0)

		for {
			select {
			case <-isDone:
				timer.Stop()
				t.Log("timer is done")
				return
			case <-timer.C:
				count++
				t.Log("count:", count)
				if count == 5 {
					close(isDone)
				}
				timer.Reset(3 * time.Second)
			}
		}
	}()
	wg.Wait()
}
