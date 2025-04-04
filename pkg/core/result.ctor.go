package core

func ResultFrom[T any](value T, err error) Result[T, error] {
	if err != nil {
		return Err[T](err)
	} else {
		return Ok[T, error](value)
	}
}

func ResultFromEither[L any, R any](either Either[L, R]) Result[L, R] {
	if either.IsLeft() {
		return Ok[L, R](either.UnwrapLeft())
	} else {
		return Err[L](either.UnwrapRight())
	}
}

func Ok[T any, E any](value T) Result[T, E] {
	r := result[T, E]{t: _OK, ok: value}
	return &r
}

func Err[T any, E any](err E) Result[T, E] {
	r := result[T, E]{t: _ERROR, err: err}
	return &r
}

func MapOk[T any, E any, U any](r Result[T, E], fn func(v T) U) Result[U, E] {
	if r.IsOk() {
		return Ok[U, E](fn(r.Ok().Unwrap()))
	} else {
		return Err[U](r.Err().Unwrap())
	}
}

func MapOkOr[T any, E any, U any](r Result[T, E], fn func(v T) U, value T) U {
	if r.IsOk() {
		return fn(r.Ok().Unwrap())
	} else {
		return fn(value)
	}
}

func MapErr[T any, E any, U any](r Result[T, E], fn func(err E) U) Result[T, U] {
	if r.IsErr() {
		return Err[T](fn(r.Err().Unwrap()))
	} else {
		return Ok[T, U](r.Ok().Unwrap())
	}
}

func MapErrOr[T any, E any, U any](r Result[T, E], fn func(v E) U, value E) U {
	if r.IsErr() {
		return fn(r.Err().Unwrap())
	} else {
		return fn(value)
	}
}

/*
func MapOkFrom[T any, U any](r Result[T], fn func(v T) Result[U]) Result[U] {
	if r.IsOk() {
		return fn(r.Ok().Unwrap())
	} else {
		return Err[U](r.Err().Unwrap())
	}
}

func MapErrFrom[T any, U any](r Result[T], fn func(err error) Either[T, U]) Either[T, U] {
	if r.IsErr() {
		return fn(r.Err().Unwrap())
	} else {
		return Left[T, U](r.Ok().Unwrap())
	}
}
*/
