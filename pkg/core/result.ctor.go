package core

import (
	"errors"
	"fmt"
)

func ResultFrom[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	} else {
		return Ok(value)
	}
}

func ResultFromEither[T any](either Either[T, any], errorFactory ...func(any) error) Result[T] {
	if either.IsLeft() {
		return Ok(either.UnwrapLeft())
	} else {
		var err error
		if len(errorFactory) != 0 {
			err = errorFactory[0](either.UnwrapRight())
		} else {
			err = errors.New(fmt.Sprint(either.UnwrapRight()))
		}
		return Err[T](err)
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

func MapOkFrom[T any, U any](r Result[T], fn func(v T) Result[U]) Result[U] {
	if r.IsOk() {
		return fn(r.Ok().Unwrap())
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

func MapErrorFrom[T any, U any](r Result[T], fn func(err error) Either[T, U]) Either[T, U] {
	if r.IsError() {
		return fn(r.Error().Unwrap())
	} else {
		return Left[T, U](r.Ok().Unwrap())
	}
}
