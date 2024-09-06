package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type DeferredCall[T any] interface {

	// Extends Awaiter
	// If the deferred function is done it will return core.Left(core.Some(core.Try[T]))
	// If the deferred function is canceled it will return core.Left(core.None[core.Try[T]]())
	Awaiter[core.Option[core.Try[T]]]

	// Cancels the deferred function
	Cancel()

	// Returns true if the deferred function was canceled
	IsCanceled() bool
}
