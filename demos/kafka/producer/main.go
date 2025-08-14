package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 配置Kafka writer
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	}

	defer w.Close()

	// 设置优雅退出
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// 生产消息
	ctx := context.Background()
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 生成随机消息
			msg := fmt.Sprintf("Message-%d", rand.Intn(1000))
			err := w.WriteMessages(ctx,
				kafka.Message{
					Key:   []byte("key"),
					Value: []byte(msg),
				},
			)
			if err != nil {
				fmt.Printf("Failed to write message: %v\n", err)
				continue
			}
			fmt.Printf("Produced: %s\n", msg)
		case <-sigchan:
			fmt.Println("Shutting down producer...")
			return
		}
	}
}
