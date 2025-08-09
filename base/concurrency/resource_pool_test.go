package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func TestForResourcePool(t *testing.T) {
	pool := NewPool(3)

	for i := 0; i < 5; i++ {
		go func(id int) {
			resCh := pool.Acquire() // 获取资源 channel
			res := <-resCh          // 获取资源
			fmt.Printf("Goroutine %d 获得资源 %d\n", id, res.ID)
			time.Sleep(time.Second) // 模拟使用资源
			resCh <- res            // 归还资源
			pool.Release(resCh)     // 归还资源 channel
			fmt.Printf("Goroutine %d 归还资源 %d\n", id, res.ID)
		}(i + 1)
	}

	time.Sleep(5 * time.Second)
}
