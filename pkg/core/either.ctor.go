package core

func EitherFromResult[T any](r Result[T]) Either[T, error] {
	var e either[T, error]
	if r.IsOk() {
		e = either[T, error]{t: _LEFT, left: r.Ok().Unwrap()}
	} else {
		e = either[T, error]{t: _RIGHT, right: r.Error().Unwrap()}
	}
	return &e
}

func Left[A any, B any](value A) Either[A, B] {
	e := either[A, B]{t: _LEFT, left: value}
	return &e
}

func Right[A any, B any](value B) Either[A, B] {
	e := either[A, B]{t: _RIGHT, right: value}
	return &e
}

func MapRight[A any, B any, R any](e either[A, B], fn func(ptr B) R) Either[A, R] {
	var other either[A, R]
	if e.IsRight() {
		other = either[A, R]{t: _RIGHT, right: fn(e.right)}
	} else {
		other = either[A, R]{t: _LEFT, left: e.left}
	}
	return &other
}

func MapRightFrom[A any, B any, R any](e either[A, B], fn func(ptr B) Either[A, R]) Either[A, R] {
	if e.IsRight() {
		return fn(e.right)
	} else {
		return &either[A, R]{t: _LEFT, left: e.left}
	}
}

func MapLeft[A any, B any, L any](e either[A, B], fn func(ptr A) L) Either[L, B] {
	var other either[L, B]
	if e.IsLeft() {
		other = either[L, B]{t: _LEFT, left: fn(e.left)}
	} else {
		other = either[L, B]{t: _RIGHT, right: e.right}
	}
	return &other
}

func MapLeftFrom[A any, B any, L any](e either[A, B], fn func(ptr A) Either[L, B]) Either[L, B] {
	if e.IsLeft() {
		return fn(e.left)
	} else {
		return &either[L, B]{t: _RIGHT, right: e.right}
	}
}
