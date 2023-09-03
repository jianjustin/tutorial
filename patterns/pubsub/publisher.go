package pubsub

import "context"

func (h *hub) publish(ctx context.Context, msg *message) error {
	h.Lock()
	for s := range h.subs {
		s.publish(ctx, msg)
	}
	h.Unlock()

	return nil
}

func (s *subscriber) publish(ctx context.Context, msg *message) {
	select {
	case <-ctx.Done():
		return
	case s.handler <- msg:
	default:
	}
}
