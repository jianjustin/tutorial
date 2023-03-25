package channels

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func worker(id int, tasks chan int, quit chan bool) {
	for {
		select {
		case task := <-tasks:
			n := time.Duration(rand.Intn(10))
			fmt.Printf("Worker %d: started task %d\n", id, task)
			time.Sleep(n * time.Second) // 模拟处理任务
			fmt.Printf("Worker %d: finished task %d; time is %d\n", id, task, n)
		case <-quit:
			fmt.Printf("Worker %d: quitting\n", id)
			return
		}
	}
}

func TestWorkerPool(t *testing.T) {
	// 创建任务和退出channel
	tasks := make(chan int)
	quit := make(chan bool)

	// 启动3个worker goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, quit)
	}

	// 发送10个任务
	for i := 1; i <= 10; i++ {
		tasks <- i
	}

	// 关闭quit channel等待所有worker goroutine退出
	close(quit)
	time.Sleep(time.Second)
	fmt.Println("All workers have quit")
}
