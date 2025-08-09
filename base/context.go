package base

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 定义context中使用的key类型
type contextKey string

const (
	requestIDKey contextKey = "requestID"
	traceIDKey   contextKey = "traceID"
)

// worker函数演示取消信号处理
func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Printf("worker %d: 开始工作\n", id)

	for {
		select {
		case <-time.After(time.Duration(rand.Intn(500)) * time.Millisecond):
			fmt.Printf("worker %d: 处理中...\n", id)
		case <-ctx.Done():
			fmt.Printf("worker %d: 收到取消信号，原因: %v\n", id, ctx.Err())
			return
		}
	}
}
