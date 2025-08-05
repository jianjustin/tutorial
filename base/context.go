package base

import (
	"context"
	"fmt"
	"time"
)

// 验证超时
func demoTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		fmt.Println("超时:", ctx.Err())
	}
}

// 验证cancel逻辑
func demoCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		cancel() // 主动取消
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("任务完成")
	case <-ctx.Done():
		fmt.Println("被取消:", ctx.Err())
	}
}
