package async

import (
	"sync"
	"sync/atomic"

	"github.com/comet11x/go-fpl/pkg/core"
)

type promise[T any] struct {
	status          atomic.Uint32
	cond            *sync.Cond
	resultMutex     *sync.Mutex
	successfulValue core.Option[T]
	failedValue     core.Option[any]
	handlerMutex    sync.RWMutex
	thenHandler     []Resolve[T]
	catchHandler    []func(any)
	finallyHandler  []func()
}

func (p *promise[T]) run(executor func(func(T), func(any))) {
	t := core.Call[any, any](func(_ any) any {
		executor(p.resolve, p.reject)
		return nil
	}, nil,
	)
	if t.IsFailure() && p.IsPending() {
		p.reject(t.Failure().Unwrap())
	}
}

func (p *promise[T]) resolve(v T) {
	p.cond.L.Lock()
	p.status.Store(FULFILLED)
	p.successfulValue = core.Some(v)
	p.failedValue = core.None[any]()
	for _, handler := range p.thenHandler {
		go handler(v)
	}
	for _, handler := range p.finallyHandler {
		go handler()
	}
	p.thenHandler = nil
	p.catchHandler = nil
	p.finallyHandler = nil
	p.cond.Broadcast()
	p.cond.L.Unlock()
}

func (p *promise[T]) reject(e any) {
	p.handlerMutex.Lock()
	p.status.Store(REJECTED)
	p.successfulValue = core.None[T]()
	p.failedValue = core.Some(e)
	for _, handler := range p.catchHandler {
		go handler(e)
	}
	for _, handler := range p.finallyHandler {
		go handler()
	}
	p.thenHandler = nil
	p.catchHandler = nil
	p.finallyHandler = nil
	p.handlerMutex.Unlock()
}

// ----------------------------------------------------------------
//  PUBLIC methods
// ----------------------------------------------------------------

func (p *promise[T]) Then(fn func(T)) Promise[T] {
	p.handlerMutex.RLock()
	isDone := p.thenHandler == nil
	switch true {
	case isDone && p.status.Load() == FULFILLED:
		go fn(p.successfulValue.Unwrap())
	case !isDone:
		p.thenHandler = append(p.thenHandler, fn)
	}
	p.handlerMutex.RUnlock()
	return p
}

func (p *promise[T]) Catch(fn func(any)) Promise[T] {
	p.handlerMutex.RLock()
	isDone := p.catchHandler == nil
	switch true {
	case isDone && p.status.Load() == REJECTED:
		// call
		go fn(p.failedValue.Unwrap())
	case !isDone:
		p.catchHandler = append(p.catchHandler, fn)
	}
	p.handlerMutex.RUnlock()
	return p
}

func (p *promise[T]) Finally(func()) Promise[T] {
	return p
}

func (p *promise[T]) Await() core.Either[T, any] {

	if p.status.Load() == PENDING {
		p.cond.L.Lock()
		p.cond.Wait()
		p.cond.L.Unlock()
	}

	if p.successfulValue.IsSome() {
		return core.Left[T, any](p.successfulValue.Unwrap())
	} else {
		return core.Right[T, any](p.failedValue.Unwrap())
	}
}

func (p *promise[T]) IsPending() bool {
	return p.status.Load() == PENDING
}

func (p *promise[T]) IsResolve() bool {
	return p.status.Load() == FULFILLED
}

func (p *promise[T]) IsRejected() bool {
	return p.status.Load() == REJECTED
}

func (p *promise[T]) Status() string {
	switch p.status.Load() {
	case FULFILLED:
		return "fulfilled"
	case REJECTED:
		return "rejected"
	default:
		return "pending"
	}
}
