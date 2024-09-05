package async

import (
	"sync"
	"sync/atomic"

	"github.com/comet11x/go-fpl/pkg/core"
)

type promise[T any] struct {
	status         atomic.Uint32
	cond           *sync.Cond
	resultMutex    *sync.RWMutex
	handlerMutex   sync.RWMutex
	result         core.Either[T, any]
	thenHandler    []ResolveHandler[T]
	catchHandler   []RejectedHandler
	finallyHandler []FinallyHandler
}

func (p *promise[T]) call(executor func(func(T), func(any))) {
	core.ImmediateCall[any, any](func(_ any) any {
		executor(p.resolve, p.reject)
		return nil
	}, nil).
		IfFailure(func(value any) {
			if p.IsPending() {
				p.reject(value)
			}
		})
}

func (p *promise[T]) resolve(v T) {
	p.cond.L.Lock()
	p.result = core.Left[T, any](v)
	p.status.Store(FULFILLED)

	p.handlerMutex.Lock()
	for _, handler := range p.thenHandler {
		handler(v)
	}
	for _, handler := range p.finallyHandler {
		handler()
	}
	p.thenHandler = nil
	p.catchHandler = nil
	p.finallyHandler = nil
	p.handlerMutex.Unlock()
	p.cond.Broadcast()
	p.cond.L.Unlock()
}

func (p *promise[T]) reject(e any) {
	p.handlerMutex.Lock()
	p.result = core.Right[T, any](e)
	p.status.Store(REJECTED)
	p.handlerMutex.Lock()
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
	p.handlerMutex.Unlock()
}

// ----------------------------------------------------------------
//
//	PUBLIC methods
//
// ----------------------------------------------------------------
func (p *promise[T]) Then(fn func(T)) Promise[T] {
	if p.status.Load() != PENDING {
		p.result.IfLeft(fn)
	} else {
		p.resultMutex.RLock()
		if p.result == nil {
			p.resultMutex.RUnlock()
			// result is not done yet
			p.handlerMutex.Lock()
			p.thenHandler = append(p.thenHandler, fn)
			p.handlerMutex.Unlock()
		} else {
			p.resultMutex.RUnlock()
			// result is already done
			p.result.IfLeft(fn)
		}
	}
	return p
}

func (p *promise[T]) Catch(fn func(any)) Promise[T] {
	if p.status.Load() != PENDING {
		p.result.IfRight(fn)
	} else {
		p.resultMutex.RLock()
		if p.result == nil {
			p.resultMutex.RUnlock()
			// result is not done yet
			p.handlerMutex.Lock()
			p.catchHandler = append(p.catchHandler, fn)
			p.handlerMutex.Unlock()
		} else {
			p.resultMutex.RUnlock()
			// result is already done
			p.result.IfRight(fn)
		}
	}
	return p
}

func (p *promise[T]) Finally(fn func()) Promise[T] {
	if p.status.Load() != PENDING {
		fn()
	} else {
		p.resultMutex.RLock()
		if p.result == nil {
			p.resultMutex.RUnlock()
			// result is not done yet
			p.handlerMutex.Lock()
			p.finallyHandler = append(p.finallyHandler, fn)
			p.handlerMutex.Unlock()
		} else {
			p.resultMutex.RUnlock()
			// result is already done
			fn()
		}
	}
	return p
}

func (p *promise[T]) Await() core.Either[T, any] {
	if p.status.Load() != PENDING {
		return p.result
	} else {
		p.cond.L.Lock()
		p.cond.Wait()
		p.cond.L.Unlock()
		return p.result
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

func (p *promise[T]) Future() Future[T] {
	return &future[T]{}
}
