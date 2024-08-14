package core

func Call[A any, R any](fn func(A) R, arg A) Try[R] {
	return &try[A, R]{arg: arg, fn: fn, status: _TRY_PND}
}
