package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	natsURL = nats.DefaultURL
)

func main() {
	ComsumerAndProducerWithDelay()

}

func ComsumerAndProducerWithDelay() {
	// 1. 连接 NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("connect error:", err)
	}
	defer nc.Drain()

	// 2. 获取 JetStream 上下文
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("jetstream error:", err)
	}

	// 3. 创建 Stream（如果不存在）
	//    这里我们用 subject: orders.cancel
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.cancel"},
	})
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		log.Fatal("add stream error:", err)
	}

	// 4. 启动消费者 goroutine：订阅 orders.cancel
	go func() {
		// 使用 Push 模式订阅
		sub, err := js.Subscribe(
			"orders.cancel",
			func(m *nats.Msg) {
				md, _ := m.Metadata()

				if md.NumDelivered == 1 {
					// 第一次投递：不真正处理，只告诉服务器 30s 后再投递一次
					_ = m.NakWithDelay(30 * time.Second)
					log.Printf("[消费者] 第一次收到取消通知，30秒后开始执行 (NumDelivered=%d)", md.NumDelivered)
					return
				}

				// 第二次（以及之后）才是真正处理
				log.Printf("[%s] 收到取消通知：%s (NumDelivered=%d)",
					time.Now().Format(time.RFC3339),
					string(m.Data),
					md.NumDelivered,
				)
				_ = m.Ack()
			},
			nats.Durable("ORDERS_CANCEL_CONSUMER"),
			nats.ManualAck(),
			nats.DeliverNew(),
		)
		if err != nil {
			log.Fatal("subscribe error:", err)
		}
		log.Println("[消费者] 已订阅 orders.cancel，等待消息...")
		// 防止 goroutine 退出
		_ = sub
		select {}
	}()

	// 给消费者一点时间完成订阅
	time.Sleep(1 * time.Second)

	// 5. 生产者发送 “30 秒后取消订单” 的消息
	orderID := "order-123"
	delay := 30 * time.Second

	log.Printf("[生产者] 当前时间: %s\n", time.Now().Format(time.RFC3339))
	log.Printf("[生产者] 发送 30 秒后取消订单的通知，订单ID=%s\n", orderID)

	if err := publishDelayCancel(js, orderID, delay); err != nil {
		log.Fatal("publish delay msg error:", err)
	}

	// 主 goroutine 等待一会儿，观察消费者输出
	time.Sleep(40 * time.Second)
	log.Println("Demo 结束")
}

// publishDelayCancel 使用 JetStream 的延迟机制：Nats-Msg-Not-Before
func publishDelayCancel(js nats.JetStreamContext, orderID string, delay time.Duration) error {
	subject := "orders.cancel"

	// 构造消息
	msg := nats.NewMsg(subject)
	// 设置不早于的投递时间（Nats-Msg-Not-Before）
	notBefore := time.Now().Add(delay).Format(time.RFC3339Nano)
	msg.Header.Set("Nats-Msg-Not-Before", notBefore)

	msg.Data = []byte(fmt.Sprintf("订单 %s 未支付，执行关闭操作（延迟 %s）", orderID, delay))

	_, err := js.PublishMsg(msg)
	return err
}

func ComsumerAndProducer() {
	subjectName := "demo.subject"
	msgCount := 5 // 要发送/接收的消息数量
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

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
