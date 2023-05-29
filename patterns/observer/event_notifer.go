package observer

type EventNotifier interface {
	Register(obs EventObserver)
	Deregister(obs EventObserver)
	Notify(event Event)
}

type eventNotifer struct {
	observers []EventObserver
}

func NewEventNotifier() EventNotifier {
	return &eventNotifer{}
}

func (e *eventNotifer) Register(obs EventObserver) {
	e.observers = append(e.observers, obs)
}

func (e *eventNotifer) Deregister(obs EventObserver) {
	for i := 0; i < len(e.observers); i++ {
		if obs == e.observers[i] {
			e.observers = append(e.observers[:i], e.observers[i+1:]...)
		}
	}
}

func (e *eventNotifer) Notify(event Event) {
	for i := 0; i < len(e.observers); i++ {
		e.observers[i].OnNotify(event)
	}
}
