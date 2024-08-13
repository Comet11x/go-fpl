package events

import "time"

func CreateEventListener(eventHandler func(Event)) EventListener {
	l := eventListener{fn: eventHandler, t: time.Now().UnixNano()}
	return &l
}
