package events

import (
	"testing"

	"github.com/comet11x/go-fpl/pkg/core"
)

func TestNewEventListener(t *testing.T) {
	NewEventListener(func(e Event) {
		//
	})
}

func TestGetIdOfEventListener(t *testing.T) {
	l := NewEventListener(func(e Event) {})
	_ = l.Id()
}

func TestCallOfEventListener(t *testing.T) {
	eventName := "event"
	payload := "test"
	l := NewEventListener(func(e Event) {
		if e.Name() != eventName {
			t.Fatalf("It must be '%s'", eventName)
		}
		if core.OptionFrom[string](e.From().ToTuple()).Unwrap() != payload {
			t.Fatalf("It must be '%s'", payload)
		}
	})

	l.Call(EventWithoutPayload(eventName, payload))
}
