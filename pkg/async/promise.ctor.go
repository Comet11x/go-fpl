package async

import (
	"sync"

	"github.com/comet11x/go-fpl/pkg/core"
)

func Async[T any](executor func(resolve func(T), reject func(any))) Promise[T] {
	m := &sync.RWMutex{}
	p := promise[T]{
		thenHandler:    make([]ResolveHandler[T], 0),
		catchHandler:   make([]RejectedHandler, 0),
		finallyHandler: make([]FinallyHandler, 0),
		resultMutex:    m,
		cond:           sync.NewCond(m),
	}
	p.status.Store(PENDING)

	// call the executor
	go p.call(executor)
	return &p
}

func Then[T any, U any](p Promise[T], fn func(T) U) Promise[U] {
	return Async[U](func(resolve func(U), reject func(any)) {
		p.Await().
			IfLeft(func(value T) {
				core.Call(fn, value).
					IfSuccess(resolve).
					IfFailure(reject)
			}).
			IfRight(reject)
	})
}

func ThenFrom[T any, U any](p Promise[T], fn func(T) Promise[U]) Promise[U] {
	return Async[U](func(resolve func(U), reject func(any)) {
		p.Await().
			IfLeft(func(value T) {
				fn(value).Await().
					IfLeft(resolve).
					IfRight(reject)
			}).
			IfRight(reject)
	})
}

func PromiseAllSettled[T any](promises ...Promise[T]) Promise[[]core.Either[T, any]] {
	return Async(func(resolve func([]core.Either[T, any]), reject func(any)) {
		out := make([]core.Either[T, any], len(promises))
		for i, p := range promises {
			out[i] = p.Await()
		}
		resolve(out)
	})
}

func PromiseAll[T any](promises ...Promise[T]) Promise[[]T] {
	return Async(func(resolve func([]T), reject func(any)) {
		if len(promises) == 0 {
			resolve(make([]T, 0))
			return
		}
		out := make([]T, 0, len(promises))
		var e core.Either[T, any]
		for _, p := range promises {
			e = p.Await()
			if e.IsRight() {
				break
			} else {
				out = append(out, e.UnwrapLeft())
			}
		}
		e.IfLeft(func(_ T) {
			resolve(out)
		}).IfRight(reject)
	})
}

func PromiseAny[T any](promises ...Promise[T]) Promise[core.Option[T]] {
	return Async(func(resolve func(core.Option[T]), reject func(any)) {
		if len(promises) == 0 {
			resolve(core.None[T]())
			return
		}
		var e core.Either[T, any]
		for _, p := range promises {
			e = p.Await()
			if e.IsLeft() {
				break
			}
		}
		e.IfLeft(func(v T) { resolve(core.Some(v)) }).IfRight(reject)
	})
}
