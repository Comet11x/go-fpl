package core

type result[T any] struct {
	t   int
	ok  T
	err error
}

func ResultFrom[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	} else {
		return Ok(value)
	}
}

func Ok[T any](value T) Result[T] {
	r := result[T]{t: _OK, ok: value}
	return &r
}

func Err[T any](err error) Result[T] {
	r := result[T]{t: _ERROR, err: err}
	return &r
}

func (r *result[T]) IsOk() bool {
	return r.t == _OK
}

func (r *result[T]) IsError() bool {
	return r.t == _ERROR
}

func (r *result[T]) UnwrapOr(value T) T {
	if r.IsOk() {
		return r.ok
	} else {
		return value
	}
}

func (r *result[T]) Unwrap() T {
	return r.ok
}

func (r *result[T]) UnwrapPtr() *T {
	return &r.ok
}

func (r *result[T]) ToTuple() (T, error) {
	return r.ok, r.err
}

func (r *result[T]) ToEither() Either[T, error] {
	if r.IsOk() {
		return Left[T, error](r.ok)
	} else {
		return Right[T, error](r.err)
	}
}

func (r *result[T]) Ok() Option[T] {
	if r.IsOk() {
		return Some[T](r.ok)
	} else {
		return None[T]()
	}
}

func (r *result[T]) Error() Option[error] {
	if r.IsError() {
		return Some[error](r.err)
	} else {
		return None[error]()
	}
}

func MapOk[T any, U any](r Result[T], fn func(v T) U) Result[U] {
	if r.IsOk() {
		return Ok[U](fn(r.Ok().Unwrap()))
	} else {
		return Err[U](r.Error().Unwrap())
	}
}

func MapError[T any, U any](r Result[T], fn func(err error) U) Either[T, U] {
	if r.IsError() {
		return Right[T, U](fn(r.Error().Unwrap()))
	} else {
		return Left[T, U](r.Ok().Unwrap())
	}
}
