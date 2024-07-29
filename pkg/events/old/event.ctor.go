package events

import (
	"time"
)

func CreateEvent(t string, options ...EventOptions) Event {
	return createEvent(t, false, options...)
}

func CreateCancelableEvent(t string, options ...EventOptions) Event {
	return createEvent(t, true, options...)
}

func createEvent(t string, cancelable bool, options ...EventOptions) Event {
	e := event{
		name:       t,
		cancelable: cancelable,
		canceled:   false,
		time:       time.Now(),
		from:       nil,
		payload:    nil,
	}

	if len(options) > 0 {
		e.from = options[0].From
		e.payload = options[0].Payload
	}

	return &e
}
