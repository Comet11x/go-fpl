package events

import (
	"sync/atomic"

	"github.com/comet11x/go-fpl/pkg/sync"
)

func CreateEventEmitter() EventEmitter {
	ee := &eventEmitter{
		storage:  make(map[string]EventListenerContext),
		mu:       sync.RealRWLocker(),
		closer:   make(chan struct{}),
		listener: make(chan EventListenerContextInfo),
		empty:    atomic.Value{},
	}
	ee.empty.Store(true)

	go ee.poll()
	return ee
}
