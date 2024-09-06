package core

func EitherFromResult[T any](r Result[T]) Either[T, error] {
	var e either[T, error]
	if r.IsOk() {
		e = either[T, error]{t: _LEFT, left: r.Ok().Unwrap()}
	} else {
		e = either[T, error]{t: _RIGHT, right: r.Err().Unwrap()}
	}
	return &e
}

func Left[L any, R any](value L) Either[L, R] {
	e := either[L, R]{t: _LEFT, left: value}
	return &e
}

func Right[L any, R any](value R) Either[L, R] {
	e := either[L, R]{t: _RIGHT, right: value}
	return &e
}

func MapRight[L any, R any, R2 any](e Either[L, R], fn func(value R) R2) Either[L, R2] {
	var other either[L, R2]
	if e.IsRight() {
		other = either[L, R2]{t: _RIGHT, right: fn(e.Right().Unwrap())}
	} else {
		other = either[L, R2]{t: _LEFT, left: e.Left().Unwrap()}
	}
	return &other
}

func MapRightFrom[L any, R any, R2 any](e Either[L, R], fn func(value R) Either[L, R2]) Either[L, R2] {
	if e.IsRight() {
		return fn(e.Right().Unwrap())
	} else {
		return &either[L, R2]{t: _LEFT, left: e.Left().Unwrap()}
	}
}

func MapLeft[L any, R any, L2 any](e Either[L, R], fn func(value L) L2) Either[L2, R] {
	var other either[L2, R]
	if e.IsLeft() {
		other = either[L2, R]{t: _LEFT, left: fn(e.Left().Unwrap())}
	} else {
		other = either[L2, R]{t: _RIGHT, right: e.Right().Unwrap()}
	}
	return &other
}

func MapLeftFrom[L any, R any, L2 any](e Either[L, R], fn func(value L) Either[L2, R]) Either[L2, R] {
	if e.IsLeft() {
		return fn(e.Left().Unwrap())
	} else {
		return &either[L2, R]{t: _RIGHT, right: e.Right().Unwrap()}
	}
}
