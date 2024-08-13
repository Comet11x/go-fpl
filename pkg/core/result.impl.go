package core

type result[T any] struct {
	t   int
	ok  T
	err error
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

func (r *result[T]) UnwrapOrValueFrom(c func() T) T {
	if r.IsOk() {
		return r.ok
	} else {
		return c()
	}
}

func (r *result[T]) UnwrapAsPtrOr(value *T) *T {
	if r.IsOk() {
		return &r.ok
	} else {
		return value
	}
}

func (r *result[T]) UnwrapAsPtrOrPtrFrom(c func() *T) *T {
	if r.IsOk() {
		return &r.ok
	} else {
		return c()
	}
}

func (r *result[T]) UnwrapOrDefault() T {
	return r.ok
}

func (r *result[T]) Unwrap() T {
	if r.IsError() {
		panic("called `Result.Unwrap()` on an `Err` value")
	}
	return r.ok
}

func (r *result[T]) UnwrapAsPtr() *T {
	if r.IsError() {
		panic("called `Result.UnwrapAsPtr()` on an `Err` value")
	}
	return &r.ok
}

func (r *result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("called `Result.UnwrapErr()` on an `Ok` value")
	}
	return r.err
}

func (r *result[T]) UnwrapErrOr(err error) error {
	if r.IsError() {
		return r.err
	} else {
		return err
	}
}

func (r *result[T]) ToTuple() (T, error) {
	return r.ok, r.err
}

func (r *result[T]) ToTupleAsPtr() (*T, error) {
	return &r.ok, r.err
}

func (r *result[T]) ToEither() Either[T, error] {
	if r.IsOk() {
		return Left[T, error](r.ok)
	} else {
		return Right[T, error](r.err)
	}
}

func (r *result[T]) ToEitherPtr() Either[*T, error] {
	if r.IsOk() {
		return Left[*T, error](&r.ok)
	} else {
		return Right[*T, error](r.err)
	}
}

func (r *result[T]) Ok() Option[T] {
	if r.IsOk() {
		return Some[T](r.ok)
	} else {
		return None[T]()
	}
}

func (r *result[T]) OkPtr() Option[*T] {
	if r.IsOk() {
		return Some[*T](&r.ok)
	} else {
		return None[*T]()
	}
}

func (r *result[T]) Error() Option[error] {
	if r.IsError() {
		return Some[error](r.err)
	} else {
		return None[error]()
	}
}
