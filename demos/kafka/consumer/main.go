package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	// 配置Kafka reader
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "test-topic",
		GroupID:  "test-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	defer r.Close()

	// 设置优雅退出
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Consumer started. Waiting for messages...")

	// 消费消息
	ctx := context.Background()
	for {
		select {
		case <-sigchan:
			fmt.Println("Shutting down consumer...")
			return
		default:
			m, err := r.ReadMessage(ctx)
			if err != nil {
				fmt.Printf("Error while reading message: %v\n", err)
				continue
			}
			if err = r.CommitMessages(ctx, m); err != nil {
				fmt.Printf("Failed to commit message: %v\n", err)
			}
			fmt.Printf("Consumed: topic=%s partition=%d offset=%d key=%s value=%s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}
}
