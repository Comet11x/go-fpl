package events

import (
	"time"

	"github.com/comet11x/go-fpl/pkg/algorithm/slice"
	"github.com/comet11x/go-fpl/pkg/core"
)

func CreateEvent(name string, payload core.Option[any], from core.Option[any]) Event {
	e := event{name: name, t: time.Now(), payload: payload, from: from}
	return &e
}

func CreateEventWithoutPayload(name string, from ...any) Event {
	return CreateEvent(name, core.None[any](), slice.Head(from))
}

func CreateEventWithPayload(name string, payload any, from ...any) Event {
	return CreateEvent(name, core.Some[any](payload), slice.Head(from))
}
