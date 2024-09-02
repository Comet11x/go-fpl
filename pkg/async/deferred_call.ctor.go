package async

import (
	"time"

	"github.com/comet11x/go-fpl/pkg/core"
)

func Defer[T any, R any](fn func(T) R, arg T, timeout time.Duration) DeferredCall[R] {
	i := deferredCall[R]{timeout: timeout, try: core.Call(fn, arg), out: nil}
	go i.call()
	return &i
}
