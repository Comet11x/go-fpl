package events

import (
	"time"

	"github.com/comet11x/go-fpl/pkg/core"
)

type Event interface {

	// Returns an event name
	Name() string

	// Return true if the event is cancelable
	IsCancelable() bool

	// Returns true if an event was canceled
	IsCanceled() bool

	// Cancels the event if it is cancelable
	Cancel()

	// Returns data about a publisher (optional)
	From() core.Option[any]

	// Returns time when an event was emitted
	Time() time.Time

	// Returns a payload (optional)
	Payload() core.Option[any]
}
