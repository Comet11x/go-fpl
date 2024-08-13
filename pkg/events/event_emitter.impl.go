package events

import (
	"sync/atomic"

	"github.com/comet11x/go-fpl/pkg/core"
	"github.com/comet11x/go-fpl/pkg/sync"
	"github.com/comet11x/go-fpl/pkg/types"
)

type eventEmitter struct {
	storage  map[string]EventListenerContext
	mu       sync.RWLocker
	closer   chan struct{}
	listener chan EventListenerContextInfo
	empty    atomic.Value
}

func (ee *eventEmitter) Close() {
	ee.closer <- struct{}{}
}

func (ee *eventEmitter) AddEventEventListener(eventName string, l EventListener) EventEmitter {
	ee.On(eventName, l)
	return ee
}

func (ee *eventEmitter) poll() {
	for {
		select {
		case info := <-ee.listener:
			ee.mu.Lock()
			lc, ok := ee.storage[info.string]
			if ok && lc.Id() == info.int64 {
				delete(ee.storage, info.string)
			}
			l := len(ee.storage)
			ee.mu.Unlock()
			if l == 0 {
				ee.empty.Store(true)
			}

		case <-ee.closer:
			ee.mu.Lock()
			for _, l := range ee.storage {
				l.RemoveAllEventListeners()
			}
			ee.mu.Unlock()
		}
	}
}

func (ee *eventEmitter) On(eventName string, l EventListener) EventEmitter {
	ee.mu.RLock()
	lc, ok := ee.storage[eventName]
	ee.mu.RUnlock()

	if !ok {
		lc = CreateEventListenerContext(eventName, ee.listener)
		ee.mu.Lock()
		ee.storage[eventName] = lc
		ee.mu.Unlock()
	}
	lc.AddPersistentEventListener(l)
	ee.empty.Store(false)
	return ee
}

func (ee *eventEmitter) Once(eventName string, l EventListener) EventEmitter {
	ee.mu.RLock()
	lc, ok := ee.storage[eventName]
	ee.mu.RUnlock()

	if !ok {
		lc = CreateEventListenerContext(eventName, ee.listener)
		ee.mu.Lock()
		ee.storage[eventName] = lc
		ee.mu.Unlock()
	}
	lc.AddTemporaryEventListener(l)
	ee.empty.Store(false)
	return ee
}

func (ee *eventEmitter) Off(eventName string, l EventListener) bool {
	ee.mu.RLock()
	lc, ok := ee.storage[eventName]
	ee.mu.RUnlock()

	if ok {
		lc.RemoveEventListener(l)
	}
	return ok
}

func (ee *eventEmitter) RemoveEventListener(eventName string, l EventListener) bool {
	return ee.Off(eventName, l)
}

func (ee *eventEmitter) RemoveAllEventListeners(eventName string) bool {
	ee.mu.Lock()
	elc, ok := ee.storage[eventName]
	if ok {
		delete(ee.storage, eventName)
	}
	ee.mu.Unlock()
	if ok {
		elc.RemoveAllEventListeners()
	}
	return ok
}

func (ee *eventEmitter) Emit(e Event) EventEmitter {
	ee.mu.RLock()
	lc, ok := ee.storage[e.Name()]
	ee.mu.RUnlock()
	if ok {
		lc.Send(e, CreateSyncModeEventPropagation())
	}
	return ee
}

func (ee *eventEmitter) listen(e Event) {
	ee.Emit(e)
}

func (ee *eventEmitter) AsyncEmit(e Event) EventEmitter {
	ee.mu.RLock()
	lc, ok := ee.storage[e.Name()]
	ee.mu.Unlock()
	if ok {
		lc.Send(e, CreateAsyncModeEventPropagation())
	}
	return ee
}

func (ee *eventEmitter) Events(eventName string) []string {
	ee.mu.RLock()
	out := make([]string, len(ee.storage))
	for k := range ee.storage {
		out = append(out, k)
	}
	ee.mu.RUnlock()
	return out
}

func (ee *eventEmitter) ListenerCount(eventName string) int {
	ee.mu.RLock()
	c := len(ee.storage)
	ee.mu.RUnlock()
	return c
}

func (em *eventEmitter) Await() core.Result[types.Void] {
	for !em.empty.Load().(bool) {
		// await
	}
	return core.Ok[types.Void](core.Void())
}

func (ee *eventEmitter) ToListener() EventListener {
	return CreateEventListener(ee.listen)
}
