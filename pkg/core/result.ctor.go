package core

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
