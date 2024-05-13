package observer

import "fmt"

type EventObserver interface {
	OnNotify(event Event)
	GetName() string
}

type observer struct {
	name string
}

func NewEventObserver(name string) EventObserver {
	return &observer{name}
}

func (o *observer) OnNotify(event Event) {
	fmt.Printf("observer '%s' received event '%s'\n", o.name, event.Id)
}

func (o *observer) GetName() string {
	return o.name
}
