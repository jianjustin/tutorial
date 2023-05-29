package observer

import "fmt"

type EventObserver interface {
	OnNotify(event Event)
}

type observer struct {
	name string
}

func NewEventObserver(name string) EventObserver {
	return &observer{name}
}

// OnNotify logs the event being notified on.
func (o *observer) OnNotify(event Event) {
	fmt.Printf("observer '%s' received event '%s'\n", o.name, event.Id)
}
