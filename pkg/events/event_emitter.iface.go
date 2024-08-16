package events

import (
	"github.com/comet11x/go-fpl/pkg/types"
)

type EventEmitter interface {
	types.Closable
	AddEventEventListener(eventName string, l EventListener)
	On(eventName string, l EventListener)
	Once(eventName string, l EventListener)
	Off(eventName string, l EventListener) bool
	RemoveEventListener(eventName string, l EventListener) bool
	RemoveAllEventListeners(eventName string) bool
	Emit(e Event)
	AsyncEmit(e Event)
	Events() []string
	ListenerCount(eventName string) int
	ToListener() EventListener
}
