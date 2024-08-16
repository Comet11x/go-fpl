package events

import (
	"github.com/comet11x/go-fpl/pkg/types"
)

type EventEmitter interface {
	types.Closable
	AddEventEventListener(eventName string, l EventListener) EventEmitter
	On(eventName string, l EventListener) EventEmitter
	Once(eventName string, l EventListener) EventEmitter
	Off(eventName string, l EventListener) bool
	RemoveEventListener(eventName string, l EventListener) bool
	RemoveAllEventListeners(eventName string) bool
	Emit(e Event) EventEmitter
	AsyncEmit(e Event) EventEmitter
	Events(eventName string) []string
	ListenerCount(eventName string) int
	ToListener() EventListener
}
