package events

var broadcastBus EventEmitter

func GetBroadcastBus() EventEmitter {
	return broadcastBus
}
