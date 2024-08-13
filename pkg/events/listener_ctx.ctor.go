package events

import (
	"sync/atomic"

	"github.com/comet11x/go-fpl/pkg/sync"
)

func CreateEventListenerContext(eventName string, ch chan<- EventListenerContextInfo) EventListenerContext {
	latch := atomic.Value{}
	latch.Store(true)
	return &eventListenerContext{
		eventName:  eventName,
		och:        ch,
		temporary:  make(map[int64]EventListener),
		persistent: make(map[int64]EventListener),
		mu:         sync.RealRWLocker(),
		latch:      latch,
	}
}
