package events

func init() {
	broadcastBus = CreateEventEmitter()
}
