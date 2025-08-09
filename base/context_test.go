package base

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestForDemoTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 22*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(5 * time.Second):
				fmt.Println("长时间任务完成")
			case <-ctx.Done():
				fmt.Printf("任务被取消: %v\n", ctx.Err())
				return
			}
		}
	}()

	wg.Wait()
}

func TestForDemoCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("main: 发出取消信号")
	cancel() // 取消所有worker

	// 等待所有worker退出
	wg.Wait()
	fmt.Println("main: 所有worker已退出")
}

func TestForDemoContext(t *testing.T) {
	ctx := context.WithValue(context.Background(), traceIDKey, "trace-67890")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			traceID, ok := ctx.Value(traceIDKey).(string)
			if !ok {
				traceID = "unknown"
			}
			fmt.Printf("goroutine %d: traceID=%s\n", id, traceID)
		}(i)
	}

	wg.Wait()
}
