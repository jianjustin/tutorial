package pubsub

import (
	"context"
	"sync"
)

type hub struct {
	sync.Mutex
	subs map[*subscriber]struct{}
}

func newHub() *hub {
	return &hub{
		subs: map[*subscriber]struct{}{},
	}
}

func (h *hub) subscribers() int {
	return len(h.subs)
}

func (h *hub) unsubscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	delete(h.subs, s)
	h.Unlock()
	close(s.quit)
	return nil
}

func (h *hub) subscribe(ctx context.Context, s *subscriber) error {
	h.Lock()
	h.subs[s] = struct{}{}
	h.Unlock()

	go func() {
		select {
		case <-s.quit:
		case <-ctx.Done():
			h.Lock()
			delete(h.subs, s)
			h.Unlock()
		}
	}()

	go s.run(ctx)

	return nil
}
