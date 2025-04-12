package core

// A From construct of `Either[T, E]` which creates a new instance that of `Either[T, E]` from `Result[T, E]`.
func EitherFromResult[T any, E any](r Result[T, E]) Either[T, E] {
	var e either[T, E]
	if r.IsOk() {
		e = either[T, E]{t: _LEFT, left: r.Ok().Unwrap()}
	} else {
		e = either[T, E]{t: _RIGHT, right: r.Err().Unwrap()}
	}
	return &e
}

// Constructs a new instance which implements `Either[L, R]` with the `Left` value.
func Left[L any, R any](value L) Either[L, R] {
	e := either[L, R]{t: _LEFT, left: value}
	return &e
}

// Constructs a new instance which implements `Either[L, R]` with the `Right` value.
func Right[L any, R any](value R) Either[L, R] {
	e := either[L, R]{t: _RIGHT, right: value}
	return &e
}

// Maps `Either[L, R]` to `Either[L, R2]` by applying a function to a contained `Right`value,
// leaving an `Left` value untouched.
func MapRight[L any, R any, R2 any](e Either[L, R], fn func(value R) R2) Either[L, R2] {
	var other either[L, R2]
	if e.IsRight() {
		other = either[L, R2]{t: _RIGHT, right: fn(e.Right().Unwrap())}
	} else {
		other = either[L, R2]{t: _LEFT, left: e.Left().Unwrap()}
	}
	return &other
}

// Maps `Either[L, R]` to `Either[L, R2]` by applying a function to a pointer to a contained `Right`value,
// leaving an `Left` value untouched.
func MapRightPtr[L any, R any, R2 any](e Either[L, R], fn func(value *R) R2) Either[L, R2] {
	var other either[L, R2]
	if e.IsRight() {
		other = either[L, R2]{t: _RIGHT, right: fn(e.Right().UnwrapAsPtr())}
	} else {
		other = either[L, R2]{t: _LEFT, left: e.Left().Unwrap()}
	}
	return &other
}

// Maps `Either[L, R]` to `Either[L, R2]` by applying a function to a contained `Right`value,
// leaving an `Left` value untouched.
// The function returns a new `Either[L, R2]`.
func MapRightFrom[L any, R any, R2 any](e Either[L, R], fn func(value R) Either[L, R2]) Either[L, R2] {
	if e.IsRight() {
		return fn(e.Right().Unwrap())
	} else {
		return &either[L, R2]{t: _LEFT, left: e.Left().Unwrap()}
	}
}

// Maps `Either[L, R]` to `Either[L2, R]` by applying a function to a contained `Left`value,
// leaving an `Right` value untouched.
func MapLeft[L any, R any, L2 any](e Either[L, R], fn func(value L) L2) Either[L2, R] {
	var other either[L2, R]
	if e.IsLeft() {
		other = either[L2, R]{t: _LEFT, left: fn(e.Left().Unwrap())}
	} else {
		other = either[L2, R]{t: _RIGHT, right: e.Right().Unwrap()}
	}
	return &other
}

// Map `Either[L, R]` to `Either[L2, R]` by applying a function to a contained `Left`value,
// leaving an `Right` value untouched.
// The function returns a new `Either[L2, R]`.
func MapLeftFrom[L any, R any, L2 any](e Either[L, R], fn func(value L) Either[L2, R]) Either[L2, R] {
	if e.IsLeft() {
		return fn(e.Left().Unwrap())
	} else {
		return &either[L2, R]{t: _RIGHT, right: e.Right().Unwrap()}
	}
}
