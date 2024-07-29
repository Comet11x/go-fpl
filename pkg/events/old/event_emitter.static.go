package events

import "sync"

var id uint32 = 0
var mutex sync.Mutex

func getEventEmitterID() uint32 {
	mutex.Lock()
	cid := id
	id++
	mutex.Unlock()
	return cid
}
