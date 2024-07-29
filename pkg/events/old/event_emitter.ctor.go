package events

import "fmt"

func New() EventEmitter {
	return CreateEventEmitter()
}

func CreateEventEmitter() EventEmitter {
	ee := &_EventEmitter{
		store: map[string]map[string]EventListener{},
		close: make(chan struct{}),
	}
	ee.id = fmt.Sprintf("%p", ee)
	return ee
}
