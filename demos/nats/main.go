package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	natsURL     = nats.DefaultURL
	subjectName = "demo.subject"
	msgCount    = 5 // 要发送/接收的消息数量
)

func main() {
	ncPub, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("连接 NATS 失败(生产者): %v", err)
	}
	defer ncPub.Close()

	ncSub, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("连接 NATS 失败(消费者): %v", err)
	}
	defer ncSub.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	// 消费者 goroutine
	go func() {
		defer wg.Done()

		// 订阅主题
		sub, err := ncSub.SubscribeSync(subjectName)
		if err != nil {
			log.Fatalf("订阅失败: %v", err)
			return
		}

		// 确保订阅已经在服务端注册
		if err := ncSub.Flush(); err != nil {
			log.Fatalf("Flush 失败(消费者): %v", err)
			return
		}

		for i := 0; i < msgCount; i++ {
			// 等待下一条消息
			msg, err := sub.NextMsg(5 * time.Second)
			if err != nil {
				log.Fatalf("接收消息失败: %v", err)
				return
			}
			log.Printf("[消费者] 收到消息 #%d: %s", i+1, string(msg.Data))
		}

		log.Printf("[消费者] 收完 %d 条消息，退出", msgCount)
	}()

	// 生产者 goroutine
	go func() {
		defer wg.Done()

		// 给消费者一点时间完成订阅（简单粗暴的方式；更严谨可用额外的 ready 通知）
		time.Sleep(500 * time.Millisecond)

		for i := 0; i < msgCount; i++ {
			body := []byte(
				fmt.Sprintf("hello nats %d", i+1),
			)
			if err := ncPub.Publish(subjectName, body); err != nil {
				log.Fatalf("发送消息失败: %v", err)
				return
			}
			log.Printf("[生产者] 发送消息 #%d: %s", i+1, string(body))
			time.Sleep(200 * time.Millisecond) // 只是方便看日志节奏
		}

		// 确保所有消息真正发送到服务端
		if err := ncPub.Flush(); err != nil {
			log.Fatalf("Flush 失败(生产者): %v", err)
			return
		}

		if err := ncPub.LastError(); err != nil {
			log.Fatalf("NATS 连接错误(生产者): %v", err)
			return
		}

		log.Printf("[生产者] 已发送完 %d 条消息，退出", msgCount)
	}()

	// 等待生产者 + 消费者结束
	wg.Wait()
}
