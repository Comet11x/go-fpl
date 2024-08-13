package events

type EventListener interface {

	// Returns the listener id
	Id() int64

	// Calls an event handler
	Call(Event)
}
