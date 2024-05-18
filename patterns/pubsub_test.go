package patterns

import (
	"sync"
	"testing"
)

func TestPublisher(t *testing.T) {
	ch := make(chan Message)
	publisher := &Publisher{Data: ch}

	go func() {
		publisher.Publish("test", "Hello")
		close(ch)
	}()

	msg := <-ch
	if msg.Topic != "test" || msg.Data != "Hello" {
		t.Errorf("Publisher.Publish() = %v; want Topic=test, Data=Hello", msg)
	}
}

func TestSubscriber(t *testing.T) {
	ch := make(chan Message)
	subscriber := &Subscriber{Topic: "test", Data: ch}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		subscriber.Subscribe(&wg)
	}()

	ch <- Message{Topic: "test", Data: "Hello"}
	close(ch)
	wg.Wait()
	// 这个测试用例主要是为了检查是否有panic，所以没有具体的断言
}

func TestPubHub(t *testing.T) {
	pubHub := NewPubHub()
	publisher := pubHub.RegisterPublisher()
	if publisher == nil {
		t.Error("PubHub.RegisterPublisher() = nil; want non-nil")
	}

	subscriber := pubHub.RegisterSubscriber("test")
	if subscriber == nil {
		t.Error("PubHub.RegisterSubscriber() = nil; want non-nil")
	}

	pubHub.Close()
	// 这个测试用例主要是为了检查是否有panic，所以没有具体的断言
}

// TestMultiSubscriber 测试多个订阅者
func TestMultiSubscriber(t *testing.T) {
	pubHub := NewPubHub()
	publisher := pubHub.RegisterPublisher()
	if publisher == nil {
		t.Error("PubHub.RegisterPublisher() = nil; want non-nil")
	}

	subscriber01 := pubHub.RegisterSubscriber("test")
	if subscriber01 == nil {
		t.Error("PubHub.RegisterSubscriber() = nil; want non-nil")
	}

	subscriber02 := pubHub.RegisterSubscriber("test")
	if subscriber02 == nil {
		t.Error("PubHub.RegisterSubscriber() = nil; want non-nil")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		subscriber01.Subscribe(&wg)
	}()
	go func() {
		defer wg.Done()
		subscriber02.Subscribe(&wg)
	}()

	publisher.Publish("topic1", "Hello")
	pubHub.Close()
	wg.Wait()
	// 这个测试用例主要是为了检查是否有panic，所以没有具体的断言
}

// TestMultiTopic
func TestMultiTopic(t *testing.T) {
	pubHub := NewPubHub()
	publisher01 := pubHub.RegisterPublisher()
	if publisher01 == nil {
		t.Error("PubHub.RegisterPublisher() = nil; want non-nil")
	}

	subscriber01 := pubHub.RegisterSubscriber("topic01")
	if subscriber01 == nil {
		t.Error("PubHub.RegisterSubscriber() = nil; want non-nil")
	}

	subscriber02 := pubHub.RegisterSubscriber("topic02")
	if subscriber02 == nil {
		t.Error("PubHub.RegisterSubscriber() = nil; want non-nil")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		subscriber01.Subscribe(&wg)
	}()
	go func() {
		defer wg.Done()
		subscriber02.Subscribe(&wg)
	}()

	publisher01.Publish("topic01", "Hello topic01")
	publisher01.Publish("topic02", "Hello topic02")
	pubHub.Close()
	wg.Wait()
	// 这个测试用例主要是为了检查是否有panic，所以没有具体的断言
}
