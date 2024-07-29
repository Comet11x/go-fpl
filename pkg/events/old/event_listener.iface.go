package events

type EventListener[T comparable] interface {
	GetID() T

	//
	// It will be called when
	Release()

	//
	//
	Call(event Event)
}
