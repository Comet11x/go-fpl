package async

import (
	"time"

	"github.com/comet11x/go-fpl/pkg/core"
)

// # DefferCall constructor
//
// This constructor creates an item of DeferredCall which provides a deferred function execution.
//
// A deferred function will be called after a determined timeout.
func Defer[T any, R any](fn func(T) R, arg T, timeout time.Duration) DeferredCall[R] {
	i := deferredCall[R]{timeout: timeout, try: core.Call(fn, arg), out: nil}
	go i.call()
	return &i
}
