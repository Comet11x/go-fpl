package events

import "time"

type Event interface {

	// Returns an event name
	Name() string

	// Return true if the event is cancelable
	IsCancelable() bool

	// Returns true if an event was canceled
	IsCanceled() bool

	// It cancels an event if the event is cancelable
	Cancel()

	// Returns data about a publisher (optional)
	From() any

	// Returns time when an event was emitted
	Time() time.Time

	// Returns a payload (optional)
	Payload() any
}
