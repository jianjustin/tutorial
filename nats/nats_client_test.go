package nats

import (
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"
)

func TestCreateSubject(t *testing.T) {
	// 连接到NATS服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 创建主题
	subject := "my1.subject"
	if err := nc.Publish(subject, []byte("Hello NATS!")); err != nil {
		log.Fatal(err)
	}
}

func TestSubscribeSync(t *testing.T) {
	// 连接到NATS服务器
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// 订阅主题
	sub, err := nc.SubscribeSync("my.subject")
	if err != nil {
		log.Fatal(err)
	}

	// 等待消息
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// 打印接收到的消息
	log.Printf("Received message: %s", msg.Data)

	// 捕获Ctrl+C信号以优雅地关闭连接
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
