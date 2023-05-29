package observer_test

import (
	"testing"

	"go.guide/patterns/observer"
)

func TestOnNotify_LogsMessage(t *testing.T) {
	event := observer.Event{Id: "something happened"}
	observer := observer.NewEventObserver("peeping tom")
	observer.OnNotify(event)
}

func TestNotify_NotifiesObservers(t *testing.T) {
	notifier := observer.NewEventNotifier()
	observers := []observer.EventObserver{
		observer.NewEventObserver("tom"),
		observer.NewEventObserver("dick"),
		observer.NewEventObserver("harry"),
	}

	for i := 0; i < len(observers); i++ {
		notifier.Register(observers[i])
	}

	notifier.Notify(observer.Event{"birthday!"})
}
