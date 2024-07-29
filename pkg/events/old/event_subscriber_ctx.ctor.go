package events

func __createEventSubscriberContext(eventName string, listener EventListener) *_EventSubscriberContext {
	return &_EventSubscriberContext{EventName: eventName, Listener: listener}
}
