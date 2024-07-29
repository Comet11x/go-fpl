package events

type EventEmitter[T comparable] interface {

	//
	// This calls each of the listeners registered
	// A
	Emit(event Event)

	//
	//
	//
	AsyncEmit(event Event)

	// EventListenerContext()
	// EventListenerParams{Id: T, callback: func(Event), release: []func()}
	// Subscribes
	On(eventName string, id T, listener EventListener) bool

	//
	//
	Once(eventName string, listener EventListener) bool

	//
	//
	Off(eventName string, listener EventListener) bool

	// Removes all listeners and stops itself
	//
	RemoveAllListeners(eventName string) bool

	RemoveListener(eventName string) EventEmitter[T]

	Listeners(eventName string) []EventListener

	//
	EventExists(eventName string) bool

	ListenerExists(eventName string, listener EventListener) bool

	//
	ListenerCount(eventName string) int

	//
	CreateEvent(eventName string, options ...EventOptions) Event

	//
	CreateListener(callback func(event Event)) EventListener

	// Awaits until an emitter will be empty
	Await()

	GetMaxListeners() int

	SetMaxListeners(int) EventEmitter[T]
}
