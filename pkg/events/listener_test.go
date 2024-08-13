package events

import (
	"testing"

	"github.com/comet11x/go-fpl/pkg/core"
)

func TestCreateEventListener(t *testing.T) {
	CreateEventListener(func(e Event) {
		//
	})
}

func TestGetIdOfEventListener(t *testing.T) {
	l := CreateEventListener(func(e Event) {})
	_ = l.Id()
}

func TestCallOfEventListener(t *testing.T) {
	eventName := "event"
	payload := "test"
	l := CreateEventListener(func(e Event) {
		if e.Name() != eventName {
			t.Fatalf("It must be '%s'", eventName)
		}
		if core.OptionFrom[string](e.From().ToTuple()).Unwrap() != payload {
			t.Fatalf("It must be '%s'", payload)
		}
	})

	l.Call(CreateEventWithoutPayload(eventName, payload))
}
