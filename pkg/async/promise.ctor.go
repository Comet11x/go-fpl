package async

import (
	"sync"

	"github.com/comet11x/go-fpl/pkg/core"
)

func Async[T any](fn func(resolve func(T), reject func(any))) Promise[T] {
	m := &sync.Mutex{}
	p := promise[T]{
		thenHandler:    make([]Resolve[T], 0),
		catchHandler:   make([]Reject, 0),
		finallyHandler: make([]func(), 0),
		resultMutex:    m,
		cond:           sync.NewCond(m),
	}
	p.status.Store(PENDING)
	go p.run(fn)
	return &p
}

func Then[T any, U any](p Promise[T], fn func(T) U) Promise[U] {
	return Async[U](func(resolve func(U), reject func(any)) {
		e := p.Await()
		if e.IsLeft() {
			try := core.Call(fn, e.UnwrapLeft())
			if try.IsSuccess() {
				resolve(try.Success().Unwrap())
			} else {
				reject(try.Failure().Unwrap())
			}
		} else {
			reject(e.UnwrapRight())
		}
	})
}

func ThenFrom[T any, U any](p Promise[T], fn func(T) Promise[U]) Promise[U] {
	return Async[U](func(resolve func(U), reject func(any)) {
		e := p.Await()
		if e.IsLeft() {
			e := fn(e.UnwrapLeft()).Await()
			if e.IsLeft() {
				resolve(e.UnwrapLeft())
			} else {
				reject(e.UnwrapRight())
			}
		} else {
			reject(e.UnwrapRight())
		}
	})
}

func PromiseAll[T any](promises ...Promise[T]) Promise[[]T] {

	return nil
}

func PromiseAny[T any](promises ...Promise[T]) Promise[[]T] {

	return nil
}
