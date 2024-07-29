package events

import (
	"reflect"
	"sync"
)

type _EventEmitter struct {
	mutex   sync.RWMutex
	id      string
	store   map[string]map[string]EventListener
	close   chan struct{}
	waiting *chan struct{}
	running bool
}

func (self *_EventEmitter) ID() string {
	return self.id
}

func (self *_EventEmitter) Close() {
	self.mutex.RLock()
	if self.running {
		self.close <- struct{}{}
	}
	self.mutex.RUnlock()
}

func (self *_EventEmitter) __loop() {
	self.mutex.RLock()
	running := self.running
	self.mutex.RUnlock()
	if !running {
		self.mutex.Lock()
		self.running = true
		self.mutex.Unlock()
		go func() {
			for {
				select {
				case <-self.close:
					self.mutex.Lock()
					self.running = false
					self.mutex.Unlock()
					return
				default:
					self.mutex.RLock()
					if len(self.store) == 0 {
						if self.waiting != nil {
							(*self.waiting) <- struct{}{}
						}
						self.mutex.RUnlock()
						self.mutex.Lock()
						self.running = false
						self.mutex.Unlock()
						return
					}
					self.mutex.RUnlock()
				}
			}
		}()
	}
}

func (self *_EventEmitter) release() {
	self.mutex.Lock()
	rft := reflect.TypeOf(&_EventListener{})
	for eventName, listeners := range self.store {
		delete(self.store, eventName)
		for _, listener := range listeners {
			if reflect.TypeOf(listener) == rft {
				listener.(*_EventListener).release()
			}
		}
	}
	self.mutex.Unlock()
}

func (self *_EventEmitter) GetID() string {
	return self.id
}

func (self *_EventEmitter) _on(eventName string, listener EventListener) {
	self.mutex.Lock()
	if _, ok := self.store[eventName]; !ok {
		self.store[eventName] = map[string]EventListener{}
	}
	self.store[eventName][listener.GetID()] = listener
	self.mutex.Unlock()
	self.__loop()
}

func (self *_EventEmitter) On(eventName string, listener EventListener) {
	go self._on(eventName, listener)
}

func (self *_EventEmitter) _off(eventName string, listener EventListener) {
	self.mutex.RLock()
	exists := self.store[eventName]
	self.mutex.RUnlock()
	if exists != nil {
		if listener != nil {
			self.mutex.RLock()
			_, ok := self.store[eventName][listener.GetID()]
			self.mutex.RUnlock()
			if ok {
				self.mutex.Lock()
				delete(self.store[eventName], listener.GetID())
				ok = len(self.store[eventName]) == 0
				if ok {
					delete(self.store, eventName)
				}
				self.mutex.Unlock()
				if reflect.TypeOf(listener) == reflect.TypeOf(&_EventListener{}) {
					go listener.(*_EventListener).release()
				}
			}
		} else {
			// clean up all
			self.mutex.Lock()
			__map := self.store[eventName]
			delete(self.store, eventName)
			self.mutex.Unlock()
			rft := reflect.TypeOf(&_EventListener{})
			for _, listener := range __map {
				if reflect.TypeOf(listener) == rft {
					go listener.(*_EventListener).release()
				}
			}
		}
	}
}

func (self *_EventEmitter) Off(eventName string, listener EventListener) {
	go self._off(eventName, listener)
}

func (self *_EventEmitter) _emit(e Event) {
	self.mutex.RLock()
	_, ok := self.store[e.Name()]
	self.mutex.RUnlock()
	if ok {
		self.mutex.RLock()
		for _, listener := range self.store[e.Name()] {
			if e.IsAsync() {
				go listener.Callback(e)
			} else {
				listener.Callback(e)
				if e.IsPropagationStopped() {
					break
				}
			}
		}
		self.mutex.RUnlock()
	}
}

func (self *_EventEmitter) Emit(e Event) {
	if e.IsAsync() {
		go self._emit(e)
	} else {
		self._emit(e)
	}
}

func (self *_EventEmitter) RemoveAllListenersByEventName(eventName string) {
	go self._off(eventName, nil)
}

func (self *_EventEmitter) DoesEventExist(eventName string) bool {
	self.mutex.RLock()
	_, ok := self.store[eventName]
	self.mutex.RUnlock()
	return ok
}

func (self *_EventEmitter) CreateEvent(name string, options ...EventOptionsType) Event {
	return CreateEvent(name, options...)
}

func (self *_EventEmitter) CreateListener(callback func(Event)) EventListener {
	return CreateEventListener(callback)
}

func (self *_EventEmitter) Wait() {
	self.mutex.RLock()
	pc := self.waiting
	self.mutex.RUnlock()
	if pc == nil {
		w := make(chan struct{})
		self.mutex.Lock()
		self.waiting = &w
		self.mutex.Unlock()
	}
	<-*self.waiting
}
