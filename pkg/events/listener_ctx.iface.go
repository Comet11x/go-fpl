package events

type EventListenerContext interface {
	Id() int64

	IsSame(EventListenerContext) bool

	IsLock() bool

	// Returns an event name
	EventName() string

	// Sends an event to listeners.
	//
	// event is an item of Event, which will be sent to listener.
	// mode is a one of modes: SYNC_MODE
	Send(event Event, mode ModeEventPropagation)

	// Adds a persistent event listener
	AddPersistentEventListener(l EventListener)

	// Adds a temporary event listener
	AddTemporaryEventListener(l EventListener)

	// Returns true if an event listener exists
	EventListenerExists(l EventListener) bool

	// Removes an event listener
	RemoveEventListener(l EventListener) bool

	// Removes all listeners
	RemoveAllEventListeners()
}
