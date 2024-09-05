package async

import (
	"sync"

	"github.com/comet11x/go-fpl/pkg/core"
)

type future[T any] struct {
	cond sync.Cond
	data core.Either[T, any]
}

func (f *future[T]) Await() core.Either[T, any] {
	f.cond.L.Lock()
	f.cond.Wait()
	e := f.data
	f.cond.L.Unlock()
	return e
}

func (f *future[T]) Promise() Promise[T] {
	return Async((func(resolve func(T), reject func(any)) {
		f.cond.L.Lock()
		f.cond.Wait()
		e := f.data
		f.cond.L.Unlock()
		e.IfLeft(resolve).IfRight(reject)
	}))
}
