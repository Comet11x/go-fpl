package async

import "github.com/comet11x/go-fpl/pkg/core"

// This constructor creates an instance of Try[R] which calls the function lazily
func Call[A any, R any](fn func(A) R, arg A) Try[R] {
	promise := Async(func(resolve func(R), reject func(any)) {
		core.ImmediateCall(fn, arg).
			IfSuccess(resolve).
			IfFailure(reject)
	})
	return &try[R]{p: promise}
}
