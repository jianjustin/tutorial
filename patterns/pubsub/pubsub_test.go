package pubsub

import (
	"context"
	"testing"
	"time"
)

func TestForPublish(t *testing.T) {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	h.subscribe(ctx, sub03)

	_ = h.publish(ctx, &message{data: []byte("test01")})
	_ = h.publish(ctx, &message{data: []byte("test02")})
	_ = h.publish(ctx, &message{data: []byte("test03")})
	time.Sleep(1 * time.Second)

	h.unsubscribe(ctx, sub03)
	_ = h.publish(ctx, &message{data: []byte("test04")})
	_ = h.publish(ctx, &message{data: []byte("test05")})

	time.Sleep(1 * time.Second)
}

func TestSubscriber(t *testing.T) {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	h.subscribe(ctx, sub03)

	if h.subscribers() != 3 {
		t.Errorf("subscribers error, want 3, get %d", h.subscribers())
	}

	h.unsubscribe(ctx, sub01)
	h.unsubscribe(ctx, sub02)
	h.unsubscribe(ctx, sub03)

	if h.subscribers() != 0 {
		t.Errorf("subscribers error, want 0, get %d", h.subscribers())
	}
}

func TestCancelSubscriber(t *testing.T) {
	ctx := context.Background()
	h := newHub()
	sub01 := newSubscriber("sub01")
	sub02 := newSubscriber("sub02")
	sub03 := newSubscriber("sub03")

	h.subscribe(ctx, sub01)
	h.subscribe(ctx, sub02)
	ctx03, cancel := context.WithCancel(ctx)
	h.subscribe(ctx03, sub03)

	if h.subscribers() != 3 {
		t.Errorf("subscribers error, want 3, get %d", h.subscribers())
	}

	// cancel subscriber 03
	cancel()
	time.Sleep(100 * time.Millisecond)
	if h.subscribers() != 2 {
		t.Errorf("subscribers error, want 2, get %d", h.subscribers())
	}

	h.unsubscribe(ctx, sub01)
	h.unsubscribe(ctx, sub02)

	if h.subscribers() != 0 {
		t.Errorf("subscribers error, want 0, get %d", h.subscribers())
	}
}
