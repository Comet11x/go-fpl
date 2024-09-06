package events

import "time"

func NewEventListener(eventHandler func(Event)) EventListener {
	l := eventListener{fn: eventHandler, t: time.Now().UnixNano()}
	return &l
}
