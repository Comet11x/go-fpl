package core

type result[T any, E any] struct {
	t   int
	ok  T
	err E
}

func (r *result[T, E]) IsOk() bool {
	return r.t == _OK
}

func (r *result[T, E]) IsErr() bool {
	return r.t == _ERROR
}

func (r *result[T, E]) IfOk(fn func(T)) Result[T, E] {
	if r.IsOk() {
		fn(r.ok)
	}
	return r
}

func (r *result[T, E]) IfOkAsPtr(fn func(*T)) Result[T, E] {
	if r.IsOk() {
		fn(&r.ok)
	}
	return r
}

func (r *result[T, E]) IfErr(fn func(E)) Result[T, E] {
	if r.IsErr() {
		fn(r.err)
	}
	return r
}

func (r *result[T, E]) IfErrAsPtr(fn func(*E)) Result[T, E] {
	if r.IsErr() {
		fn(&r.err)
	}
	return r
}

func (r *result[T, E]) Expect(msg string) T {
	if r.IsErr() {
		panic(msg)
	}
	return r.ok
}

func (r *result[T, E]) Unwrap() T {
	return r.Expect("called `Result.Unwrap()` on an `Err` value")
}

func (r *result[T, E]) UnwrapOr(value T) T {
	if r.IsOk() {
		return r.ok
	} else {
		return value
	}
}

func (r *result[T, E]) UnwrapOrElse(c func() T) T {
	if r.IsOk() {
		return r.ok
	} else {
		return c()
	}
}

func (r *result[T, E]) UnwrapAsPtrOr(value *T) *T {
	if r.IsOk() {
		return &r.ok
	} else {
		return value
	}
}

func (r *result[T, E]) UnwrapAsPtrOrElse(c func() *T) *T {
	if r.IsOk() {
		return &r.ok
	} else {
		return c()
	}
}

func (r *result[T, E]) UnwrapOrDefault() T {
	return r.ok
}

func (r *result[T, E]) UnwrapAsPtr() *T {
	return r.ExpectAsPtr("called `Result.UnwrapAsPtr()` on an `Err` value")
}

func (r *result[T, E]) ExpectAsPtr(msg string) *T {
	if r.IsErr() {
		panic(msg)
	}
	return &r.ok
}

func (r *result[T, E]) ExpectErr(msg string) E {
	if r.IsOk() {
		panic(msg)
	}
	return r.err
}

func (r *result[T, E]) UnwrapErr() E {
	return r.ExpectErr("called `Result.UnwrapErr()` on an `Ok` value")
}

func (r *result[T, E]) ExpectErrAsPtr(msg string) *E {
	if r.IsOk() {
		panic(msg)
	}
	return &r.err
}

func (r *result[T, E]) UnwrapAsErr() *E {
	return r.ExpectErrAsPtr("called `Result.UnwrapErr()` on an `Ok` value")
}

func (r *result[T, E]) UnwrapErrOr(err E) E {
	if r.IsErr() {
		return r.err
	} else {
		return err
	}
}

func (r *result[T, E]) UnwrapErrOrDefault() E {
	return r.err
}

func (r *result[T, E]) AsTuple() (T, E) {
	return r.ok, r.err
}

func (r *result[T, E]) AsTupleOfPtr() (*T, *E) {
	return &r.ok, &r.err
}

func (r *result[T, E]) AsEither() Either[T, E] {
	if r.IsOk() {
		return Left[T, E](r.ok)
	} else {
		return Right[T](r.err)
	}
}

func (r *result[T, E]) AsEitherPtr() Either[*T, *E] {
	if r.IsOk() {
		return Left[*T, *E](&r.ok)
	} else {
		return Right[*T](&r.err)
	}
}

func (r *result[T, E]) Ok() Option[T] {
	if r.IsOk() {
		return Some(r.ok)
	} else {
		return None[T]()
	}
}

func (r *result[T, E]) OkAsPtr() Option[*T] {
	if r.IsOk() {
		return Some(&r.ok)
	} else {
		return None[*T]()
	}
}

func (r *result[T, E]) Err() Option[E] {
	if r.IsErr() {
		return Some(r.err)
	} else {
		return None[E]()
	}
}

func (r *result[T, E]) ErrAsPtr() Option[*E] {
	if r.IsErr() {
		return Some(&r.err)
	} else {
		return None[*E]()
	}
}

/*
func (r *result[T, E]) MapOk(fn func(v T) T) Result[T, E] {
	if r.IsOk() {
		return Ok(fn(r.ok))
	} else {
		return r
	}
}

func (r *result[T, E]) MapOkFrom(fn func(v T) Result[T, E]) Result[T, E] {
	if r.IsOk() {
		return fn(r.ok)
	} else {
		return r
	}
}

func (r *result[T, E]) MapOkAsOption(fn func(v T) T) Option[T] {
	if r.IsOk() {
		return Some(fn(r.ok))
	} else {
		return None[T]()
	}
}

func (r *result[T, E]) MapOkAsOptionFrom(fn func(v T) Option[T]) Option[T] {
	if r.IsOk() {
		return fn(r.ok)
	} else {
		return None[T]()
	}
}

func (r *result[T, E]) MapErr(fn func(e error) T) Result[T, E] {
	if r.IsErr() {
		return Ok(fn(r.err))
	} else {
		return r
	}
}

func (r *result[T, E]) MapErrFrom(fn func(e error) Result[T, E]) Result[T, E] {
	if r.IsErr() {
		return fn(r.err)
	} else {
		return r
	}
}

func (r *result[T, E]) MapErrAs(fn func(e error) T) Option[T] {
	if r.IsErr() {
		return Some(fn(r.err))
	} else {
		return None[T]()
	}
}

func (r *result[T, E]) MapErrAsFrom(fn func(e E) Option[T]) Option[T] {
	if r.IsErr() {
		return fn(r.err)
	} else {
		return None[T]()
	}
}
*/
