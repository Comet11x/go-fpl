package core

// This constructor creates an instance of Try[R] which calls the function lazily
func Call[A any, R any](fn func(A) R, arg A) Try[R] {
	return &try[A, R]{arg: arg, fn: fn, status: _TRY_PND}
}

// This constructor creates an instance of Try[R] which calls the function immediately
func ImmediateCall[A any, R any](fn func(A) R, arg A) Try[R] {
	t := &try[A, R]{arg: arg, fn: fn, status: _TRY_PND}
	t.call()
	return t
}
