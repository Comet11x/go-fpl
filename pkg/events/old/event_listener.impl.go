package events

type _EventListener[T comparable] struct {
	id       T
	callback func(Event)
	release  func()
}

func (e *_EventListener[T]) GetID() T {
	return e.id
}

func (self *_EventListener[T]) Call(e Event) {
	self.callback(e)
}

func (self *_EventListener[T]) __release() {
	// Default Release
}

func (self *_EventListener[T]) Release() {
	self.release()
}

// add(key T, fn func(Event), release ...func())
