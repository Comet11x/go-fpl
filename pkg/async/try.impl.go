package async

import (
	"errors"
	"fmt"

	"github.com/comet11x/go-fpl/pkg/algorithm/slice"
	"github.com/comet11x/go-fpl/pkg/core"
)

type try[T any] struct {
	p Promise[T]
}

func (t *try[T]) IsSuccess() Promise[bool] {
	return Async(func(resolve func(bool), reject func(any)) {
		t.p.Await()
		resolve(t.p.IsFulfilled())
	})
}

func (t *try[T]) IsFailure() Promise[bool] {
	return Async(func(resolve func(bool), reject func(any)) {
		t.p.Await()
		resolve(t.p.IsRejected())
	})
}

func (t *try[T]) IfSuccess(fn func(value T)) {
	t.p.Await().IfLeft(fn)
}

func (t *try[T]) IfFailure(fn func(value any)) {
	t.p.Await().IfRight(fn)
}

func (t *try[T]) Success() Promise[core.Option[T]] {
	return Async(func(resolve func(core.Option[T]), reject func(any)) {
		t.p.Await().
			IfLeft(func(value T) {
				resolve(core.Some(value))
			}).
			IfRight(func(value any) {
				resolve(core.None[T]())
			})
	})
}

func (t *try[T]) Failure() Promise[core.Option[any]] {
	return Async(func(resolve func(core.Option[any]), reject func(any)) {
		t.p.Await().
			IfRight(func(value any) {
				resolve(core.Some(value))
			}).
			IfLeft(func(_ T) {
				resolve(core.None[any]())
			})
	})
}

func (t *try[T]) AsResult(errorFactory ...func(any) error) Promise[core.Result[T]] {
	return Async(func(resolve func(core.Result[T]), reject func(any)) {
		t.p.Await().
			IfLeft(func(value T) {
				resolve(core.Ok(value))
			}).
			IfRight(func(value any) {
				ctor := slice.
					Head(errorFactory).
					UnwrapOr(func(a any) error { return errors.New(fmt.Sprint(a)) })
				reject(core.Err[T](ctor(value)))
			})
	})
}
