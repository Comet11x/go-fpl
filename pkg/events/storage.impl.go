package events

import (
	"sync"

	"github.com/comet11x/go-fpl/pkg/core"
)

type storage[K comparable, T any] struct {
	mu sync.RWMutex
	sm map[K]T
	c  chan struct{}
}

func (s *storage[K, T]) Get(k K) core.Option[T] {
	s.mu.RLock()
	v, ok := s.sm[k]
	s.mu.RUnlock()
	return core.OptionFrom[T](v, ok)
}

func (s *storage[K, T]) Insert(k K, v T) {
	s.mu.Lock()
	s.sm[k] = v
	s.mu.Unlock()
}

func (s *storage[K, T]) Length() int {
	s.mu.RLock()
	l := len(s.sm)
	s.mu.RUnlock()
	return l
}

func (s *storage[K, T]) check() {
	if s.Length() == 0 && s.c != nil {
		s.c <- struct{}{}
	}
}

func (s *storage[K, T]) Remove(k K) core.Option[T] {
	s.mu.Lock()
	v, ok := s.sm[k]
	o := core.OptionFrom[T](v, ok)
	if o.IsSome() {
		delete(s.sm, k)
	}
	s.mu.RUnlock()
	go s.check()
	return o
}

func (s *storage[K, T]) Exists(k K) bool {
	s.mu.RLock()
	_, ok := s.sm[k]
	s.mu.RUnlock()
	return ok
}

func (s *storage[K, T]) Await() {
	if s.c != nil {
		s.c = make(chan struct{})
	}
	<- s.c
	s.c = nil
}