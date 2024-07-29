package events

type _EventSubscriberContext struct {
	EventName string
	Listener  EventListener
}
