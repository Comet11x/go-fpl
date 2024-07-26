package core

type Result[T any] struct {
	ok  *T
	err error
}

func ResultFromPtr[T any](ptr *T, err error) Result[T] {
	if err == nil {
		return Result[T]{err: err}
	} else {
		return Result[T]{ok: ptr}
	}
}

func ResultFrom[T any](value T, err error) Result[T] {
	return ResultFromPtr(&value, err)
}

func Ok[T any](data T) Result[T] {
	return Result[T]{ok: &data}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

func (r *Result[T]) IsError() bool {
	return r.err != nil
}

func (r *Result[T]) UnwrapOr(value T) *T {
	if r.IsOk() {
		return r.ok
	} else {
		return &value
	}
}

func (r *Result[T]) Unwrap() *T {
	return r.ok
}

func (r *Result[T]) UnwrapClone() T {
	return *r.ok
}

func (r *Result[T]) ToTuple() (*T, error) {
	return r.ok, r.err
}

func (r *Result[T]) ToEither() Either[T, error] {
	return Either[T, error]{left: r.ok, right: &r.err}
}

func (r *Result[T]) Ok() Option[T] {
	if r.IsOk() {
		return Option[T]{value: r.ok}
	} else {
		return None[T]()
	}
}

func (r *Result[T]) Error() Option[error] {
	if r.IsError() {
		return Some[error](r.err)
	} else {
		return None[error]()
	}
}

func MapOk[T any, U any](r Result[T], fn func(v *T) U) Result[U] {
	if r.IsOk() {
		return Ok[U](fn(r.ok))
	} else {
		return Err[U](r.err)
	}
}

func MapError[T any, U any](r Result[T], fn func(err error) U) Either[T, U] {
	if r.IsError() {
		return Right[T, U](fn(r.err))
	} else {
		return Either[T, U]{left: r.ok}
	}
}
