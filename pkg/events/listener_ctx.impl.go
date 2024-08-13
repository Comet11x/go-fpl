package events

import (
	"sync/atomic"

	"github.com/comet11x/go-fpl/pkg/algorithm/hashmap"
	"github.com/comet11x/go-fpl/pkg/sync"
)

type EventListenerContextInfo struct {
	string
	int64
}

type eventListenerContext struct {
	id         int64
	eventName  string
	mu         sync.RWLocker
	och        chan<- EventListenerContextInfo
	temporary  map[int64]EventListener
	persistent map[int64]EventListener
	latch      atomic.Value
}

func (lc *eventListenerContext) IsLock() bool {
	return lc.latch.Load().(bool)
}

func (lc *eventListenerContext) EventName() string {
	return lc.eventName
}

func (lc *eventListenerContext) sendToPersistentListener(e Event, mode ModeEventPropagation) {
	if lc.latch.Load().(bool) {
		listeners := hashmap.Values(lc.persistent, lc.mu)
		for _, l := range listeners {
			if mode.IsAsync() {
				go l.Call(e)
			} else {
				l.Call(e)
			}
		}
	}
}

func (lc *eventListenerContext) sendToTemporaryListener(e Event, mode ModeEventPropagation) {
	keys := hashmap.Keys(lc.temporary, lc.mu)
	for _, k := range keys {
		lc.mu.Lock()
		l := lc.temporary[k]
		delete(lc.temporary, k)
		lc.mu.Unlock()
		if mode.IsAsync() {
			go l.Call(e)
		} else {
			l.Call(e)
		}
	}
	go lc.check()
}

func (lc *eventListenerContext) Send(e Event, mode ModeEventPropagation) {
	if mode.IsAsync() {
		go lc.sendToPersistentListener(e, mode)
		go lc.sendToTemporaryListener(e, mode)
	} else {
		lc.sendToPersistentListener(e, mode)
		lc.sendToTemporaryListener(e, mode)
	}
}

func (lc *eventListenerContext) AddPersistentEventListener(l EventListener) {
	if lc.latch.Load().(bool) {
		lc.mu.Lock()
		defer lc.mu.Unlock()
		lc.persistent[l.Id()] = l
	}
}

func (lc *eventListenerContext) AddTemporaryEventListener(l EventListener) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.temporary[l.Id()] = l
}

func (lc *eventListenerContext) removeEventListener(id int64, m map[int64]EventListener) bool {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	_, ok := m[id]
	if ok {
		delete(m, id)
	}
	return ok
}

func (lc *eventListenerContext) RemoveAllEventListeners() {
	lc.mu.Lock()
	lc.persistent = make(map[int64]EventListener)
	lc.temporary = make(map[int64]EventListener)
	lc.mu.Unlock()
	go lc.check()
}

func (lc *eventListenerContext) RemoveEventListener(l EventListener) bool {
	id := l.Id()
	ok := lc.removeEventListener(id, lc.temporary) || lc.removeEventListener(id, lc.persistent)
	go lc.check()
	return ok
}

func (lc *eventListenerContext) EventListenerExists(l EventListener) bool {
	lc.mu.RLock()
	_, ok := lc.persistent[l.Id()]
	if !ok {
		_, ok = lc.persistent[l.Id()]
	}
	lc.mu.RUnlock()
	return ok
}

func (lc *eventListenerContext) check() {
	lc.mu.RLock()
	l1 := len(lc.persistent)
	l2 := len(lc.temporary)
	lc.mu.RUnlock()
	if l1 == 0 && l2 == 0 {
		lc.latch.Store(false)
		lc.och <- EventListenerContextInfo{lc.eventName, lc.id}
	}
}

func (lc *eventListenerContext) Id() int64 {
	return lc.id
}

func (lc *eventListenerContext) IsSame(other EventListenerContext) bool {
	return lc.id == other.Id()
}
