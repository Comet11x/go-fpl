package async

import (
	"github.com/comet11x/go-fpl/pkg/core"
)

type DeferredCall[T any] interface {

	// Extends Awaiter
	// If the deferred call is done it will return core.Left(core.Some(core.Try[T]))
	// If the deferred call is canceled it will return core.Left(core.None[core.Try[T]]())
	Awaiter[core.Option[core.Try[T]]]

	// Cancels the deferred call
	Cancel()

	// Returns true if the call was canceled
	IsCanceled() bool
}
