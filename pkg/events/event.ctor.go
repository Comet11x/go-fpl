package events

import (
	"time"

	"github.com/comet11x/go-fpl/pkg/algorithm/slice"
	"github.com/comet11x/go-fpl/pkg/core"
)

func createEvent(name string, cancelable bool, payload core.Option[any], from core.Option[any]) Event {
	e := event{name: name, cancelable: cancelable, canceled: false, t: time.Now(), payload: payload, from: from}
	return &e
}

func CreateAsyncEvent(name string, payload core.Option[any], from ...any) Event {
	return createEvent(name, false, core.None[any](), slice.Head(from))
}

func CreateAsyncEventWithPayload(name string, payload any, from ...any) Event {
	return createEvent(name, false, core.Some(payload), slice.Head(from))
}

func CreateEventPayload(name string, from ...any) Event {
	return createEvent(name, true, core.None[any](), slice.Head(from))
}

func CreateEventWitPayload(name string, payload any, from ...any) Event {
	return createEvent(name, true, core.Some(payload), slice.Head(from))
}
